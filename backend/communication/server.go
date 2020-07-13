package communication

import (
	"encoding/json"
	"fmt"
	"github.com/danieligazit/chinese-poker/backend/game"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const PasswordURLKey = "password"
const PlayerIdURLKey = "clientId"

type Server struct {
	Id                  string
	Password            string
	clients             map[uint32]*Client
	clientId2Index      map[uint32]int
	upgrader            *websocket.Upgrader
	game                game.Game
	messageType2Handler map[string]clientMessageHandler
}

func NewServer(id, password string, g game.Game) *Server {
	var server = Server{
		Id:             id,
		Password:       password,
		clients:        map[uint32]*Client{},
		clientId2Index: map[uint32]int{},
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		Game: g,
	}
	server.messageType2Handler = map[string]clientMessageHandler{
		cakeMoveMessage: server.makeMoveHandler,
		// 		connectMessage:  connectHandler,
		// 		chatMEssage:     chatHandler,
	}
	http.HandleFunc(server.Id, g.handler)

}

func (s *Server) handler(w http.ResponseWriter, req *http.Request) {

	password := r.URL.Query().Get(PasswordURLKey)
	if password != s.password {
		http.Errorf(w, "Password is incorrect", http.StatusForbidden)
	}

	clientIdStr := r.URL.Query().Get(PlayerIdURLKey)
	clientId, err := strconv.Atoi(idUser)
	if err != nil {
		http.Errorf(w, fmt.Errorf("Error parsing clientId: %w", err), http.StatusBadRequest)
		return
	}

	conn, err := s.upgrader.Upgrade(w, req, nil)
	if err != nil {
		err = fmt.Errorf("Failed to upgrade to websockets: %w", err)
		http.Errorf(w, err, http.StatusInternalServerError)
		return
	}

	client := client.NewClient(clientId, conn, server)
	s.clients[clientId] = &client
	if _, exists := s.clientId2Index[clientId]; !exists {
		s.clientId2Index[clientId] = len(s.clientId2Index)
	}

	log.Infof("Added new client with id %d. There are currently %d clients connected", clientId, len(s.clients))
	client.Listen()

}

type ClientMessage struct {
	actionType string
	action     interface{}
}

func (s *Server) sendErrorToClient(clientId uint32, err err) {
	log.Errorf(err)
	s.sendToClient(clientId, err)
}

func (s *Server) HandleClientMessage(clientId uint32, clientMessageI interface{}) {
	var clientMessage ClientMessage
	if err = json.Unmarshal(clientMessageI, &clientMessage); err != nil {
		s.sendErrorToClient(clientId, fmt.Errorf("Error unmarshaling client message: %w", err))
	}

	messageHandler, ok := s.messageType2Handler[clientMessage.actionType]
	if !ok {
		s.sendErrorToClient(clientId, fmt.Errorf("Unspported action %s", clientMessage.actionType))
	}

	response, err := messageHandler(clientId, clientMessage.action)
	if err != nil {
		s.sendErrorToClient(clientId, fmt.Errorf("Unspported action %s", clientMessage.actionType))
	}

	s.sendToClient(clientId, response)

}

func (s *Server) sendToClient(clientID uint32, message interface{}) {
	bytes, err := json.Marshal(message)
	if err != nil {
		err = fmt.Errorf("Error marshaling message to json: %w, message=%v", err, message)
		log.Errorf(err)
		client.SendMessage([]bytes(err.Error()))
	}
	client, ok := s.clients[clientID]
	if ok {
		client.SendMessage(bytes)
	} else {
		log.Errorf("Client %d not found", clientID)
		return
	}
}

func (s *Server) sendToAllClients(message interface{}) {
	bytes := marshalMessage(message)
	for _, c := range s.clients {
		c.SendMessage(bytes)
	}
}
