package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"sync"
	"time"

	"github.com/Patrick-Old/Office-Hour-Helper/sessions"
	"github.com/go-redis/redis"
)

type Director func(r *http.Request)

func CustomDirector(target *url.URL) Director {
	return func(r *http.Request) {
		r.Header.Add("X-Forwarded-Host", r.Host)
		r.Host = target.Host
		r.URL.Host = target.Host
		r.URL.Scheme = target.Scheme
	}
}

type ServiceRegistry struct {
	Registry map[string]*ServiceInfo
	Mx       sync.Mutex
	Redis    *redis.Client
}

type ServiceInfo struct {
	Addresses   map[string]time.Time
	PathPattern string
	count       int
	mx          sync.Mutex
	Priviledged bool
}

type ServiceEvent struct {
	ServiceName   string    `json:"name"`
	PathPattern   string    `json:"pathPattern"`
	Address       string    `json:"address"`
	LastHeartbeat time.Time `json:"lastHeartbeat"`
	Priviledged   bool      `json:"priviledged"`
}

func (sr *ServiceRegistry) Update() {
	res, err := sr.Redis.LRange("ServiceEvents", 0, -1).Result()
	if err != nil {
		log.Println(err)
		sr.Redis.Del("ServiceEvents").Result()
	}
	for _, event := range res {
		se := &ServiceEvent{}
		if err := json.Unmarshal([]byte(event), se); err != nil {
			log.Println(err)
		}
		if sr.Registry[se.ServiceName] == nil {
			sr.Registry[se.ServiceName] = &ServiceInfo{
				Addresses:   make(map[string]time.Time),
				PathPattern: se.PathPattern,
				Priviledged: se.Priviledged,
			}
		} else {
			sr.Registry[se.ServiceName].Addresses[se.Address] = se.LastHeartbeat
		}
	}
	sr.Redis.Del("ServiceEvents").Result()
	for name, service := range sr.Registry {
		for addr, timestamp := range service.Addresses {
			if timestamp.Before(time.Now().Add(-time.Second * time.Duration(20))) {
				delete(service.Addresses, addr)
				if len(service.Addresses) == 0 {
					delete(sr.Registry, name)
				}
			}
		}
	}
}

func (ctx *Context) ServiceDiscovery(sr *ServiceRegistry) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		for _, service := range sr.Registry {
			match, _ := regexp.MatchString(service.PathPattern, r.URL.Path)
			if match && len(service.Addresses) > 0 {
				if service.Priviledged {
					currSession := SessionState{}
					if _, err := sessions.GetState(r, ctx.SigningKey, ctx.SessionStore, &currSession); err != nil {
						http.Error(w, "please sign-in", http.StatusUnauthorized)
						r.Header.Del("X-User")
						return
					}
					jsonUsr, _ := json.Marshal(currSession.User)
					r.Header.Set("X-User", string(jsonUsr))
				}
				service.mx.Lock()
				options := make([]string, len(service.Addresses))
				count := 0
				for key := range service.Addresses {
					options[count] = key
					count++
				}
				choice := options[service.count%len(options)]
				service.count++
				proxy := &httputil.ReverseProxy{Director: CustomDirector(&url.URL{Scheme: "http", Host: choice})}
				proxy.ServeHTTP(w, r)
				service.mx.Unlock()
				return
			}
		}
		w.WriteHeader(404)
		w.Write([]byte("Page not found"))
	})
}
