package communication

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Server struct{
    Id string
    Passwowd string
	clients []Client
	upgrader *websocket.Upgrader
}

func NewServer(id, password string) *ClientPool {
    
    return &ClientPool{
        Id: id,
        Passwowd: password,
        clients: []ClientPool{},
        upgrader : websocket.Upgrader{
    		ReadBufferSize:  1024,
    		WriteBufferSize: 1024,
    		CheckOrigin: func(r *http.Request) bool {
    			return true
    		},
    	},
    }
}

func (c* ClientPool) Handler(w http.ResponseWriter, req *http.Request) {

	conn, err := c.upgrader.Upgrade(w, req, nil)

	if err != nil {
		log.Println(err)
		return
	}

	client := client.NewClient(conn)

}