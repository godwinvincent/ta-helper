package handlers

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Notifier struct {
	connections      map[string][]*WebSocketConnection
	lastConnectionId int
	mx               sync.RWMutex
}

type WebSocketConnection struct {
	conn   *websocket.Conn
	id     int
	userID string
}

type msg struct {
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func NewNotifier() *Notifier {
	notify := &Notifier{
		connections:      make(map[string][]*WebSocketConnection),
		lastConnectionId: 0,
	}
	return notify
}

func (n *Notifier) add(id string, conn *websocket.Conn) *WebSocketConnection {
	n.mx.Lock()
	defer n.mx.Unlock()
	if n.connections[id] == nil {
		n.connections[id] = make([]*WebSocketConnection, 0)
	}
	n.lastConnectionId++
	connection := &WebSocketConnection{conn, n.lastConnectionId, id}
	n.connections[id] = append(n.connections[id], connection)
	return connection
}

func (n *Notifier) remove(userID string, connectionID int) {
	n.mx.Lock()
	defer n.mx.Unlock()
	if n.connections[userID] != nil {
		for i, conn := range n.connections[userID] {
			if conn.id == connectionID {
				//remove that conn from the array
				n.connections[userID] = append(n.connections[userID][:i], n.connections[userID][i+1:]...)
				break
			}
		}
	}
}

func (n *Notifier) Dispatch(userIDs []string, msg []byte) {
	if len(userIDs) == 0 || userIDs == nil {
		for _, connections := range n.connections {
			for _, connection := range connections {
				if err := connection.conn.WriteJSON(msg); err != nil {
					log.Printf(err.Error())
					connection.conn.Close()
					n.remove(connection.userID, connection.id)
				}
			}
		}
	} else {
		for _, userID := range userIDs {
			conns := n.connections[userID]
			if len(conns) > 0 {
				for _, connection := range conns {
					if err := connection.conn.WriteJSON(msg); err != nil {
						connection.conn.Close()
						n.remove(connection.userID, connection.id)
					}
				}
			}
		}
	}
}

func (ctx *Context) WebSocketConnectionHandler(w http.ResponseWriter, r *http.Request, currSession *SessionState) {
	// TODO: change accepted origin URL when deploying to
	if r.Header.Get("Origin") != "http://localhost:3000" {
		http.Error(w, "Websocket Connection Refused, bad origin", 403)
		return
	}
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Failed to open websocket connection", 401)
	}
	// connection := ctx.NotificationStore.add(currSession.User.ID, conn)
	connection := ctx.NotificationStore.add(currSession.User.UserName, conn)
	go ctx.handleSocket(connection)
}

func (ctx *Context) handleSocket(connection *WebSocketConnection) {
	for {
		m := msg{}
		conn := connection.conn
		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
			conn.Close()
			break
		}
	}
	ctx.NotificationStore.remove(connection.userID, connection.id)
}
