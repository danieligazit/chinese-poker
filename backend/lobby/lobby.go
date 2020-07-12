package lobby

import (
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"

	"superstellar/backend/events"
	"superstellar/backend/monitor"
	"superstellar/backend/utils"

	"github.com/gorilla/websocket"
)

// Server struct holds server variables.
type Server struct {
	clients          map[uint32]*Client
	upgrader         *websocket.Upgrader
}

// NewServer initializes a new server.
func NewServer() *Server {
	return &Server{
		clients:          make(map[uint32]*Client),
	}
}


func (s *Server) SendToAllClients(message proto.Message) {
	bytes := marshalMessage(message)
	for _, c := range s.clients {
		c.SendMessage(bytes)
	}
}

func (s *Server) SendToClient(clientID uint32, message proto.Message) {
	bytes := marshalMessage(message)

	client, ok := s.clients[clientID]
	if ok {
		client.SendMessage(bytes)
	} else {
		log.Printf("Client %d not found\n", clientID)
		return
	}
}

func (s *Server) ClientIDs() []uint32 {
	clientIDs := make([]uint32, 0, len(s.clients))
	for k := range s.clients {
		clientIDs = append(clientIDs, k)
	}

	return clientIDs
}

func (s *Server) GetClient(clientId uint32) (*Client, bool) {
	client, ok := s.clients[clientId]
	return client, ok
}