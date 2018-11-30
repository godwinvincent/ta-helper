package handlers

import "net/http"

type CorsHeader struct {
	handler http.Handler
}

//NewCorsHeader constructs a new CorsHeader middleware handler
func NewCorsHeader(handlerToWrap http.Handler) *CorsHeader {
	return &CorsHeader{handlerToWrap}
}

func (ch *CorsHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//add the headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	w.Header().Set("Access-Control-Max-Age", "600")
	//call the wrapped handler
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
		ch.handler.ServeHTTP(w, r)
	}
}
