package communication

import (
	"encoding/json"
	"fmt"
	"github.com/danieligazit/chinese-poker/backend/games"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

const PasswordURLKey = "password"
const PlayerIdURLKey = "clientId"
const BaseURI = "/%s"

type Server struct {
	Id                  string
	password            string
	clients             map[uint64]*Client
	clientId2Index      map[uint64]int
	upgrader            *websocket.Upgrader
	game                games.IGame
	messageType2Handler map[string]clientMessageHandler
}

func NewServer(id, password string, game games.IGame) *Server {
	var server = Server{
		Id:             id,
		password:       password,
		clients:        map[uint64]*Client{},
		clientId2Index: map[uint64]int{},
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		game: game,
	}
	server.messageType2Handler = map[string]clientMessageHandler{
		makeMoveMessage: server.makeMoveHandler,
		connectMessage:  server.connectHandler,
		// 		chatMEssage:     server.chatHandler,
	}

	uri := fmt.Sprintf(BaseURI, server.Id)
	http.HandleFunc(uri, server.handler)
	log.Infof("New server at endpoint %s", uri)
	return &server
}

func (s *Server) handler(w http.ResponseWriter, req *http.Request) {

	password := req.URL.Query().Get(PasswordURLKey)
	if password != s.password {
		http.Error(w, "Password is incorrect", http.StatusForbidden)
		return
	}

	clientIdStr := req.URL.Query().Get(PlayerIdURLKey)
	clientId, err := strconv.ParseUint(clientIdStr, 10, 64)
	if err != nil {
		err = fmt.Errorf("Error parsing clientId: %w", err)
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	conn, err := s.upgrader.Upgrade(w, req, nil)
	if err != nil {
		err = fmt.Errorf("Failed to upgrade to websockets: %w", err)
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := NewClient(clientId, conn, s)
	s.clients[clientId] = client
	if _, exists := s.clientId2Index[clientId]; !exists {
		s.clientId2Index[clientId] = len(s.clientId2Index)
	}

	log.Infof("Added new client with id %d. There are currently %d clients connected", clientId, len(s.clients))
	client.Listen()

}

func (s *Server) sendErrorToClient(clientId uint64, err error) {
	log.Errorf(err.Error())
	s.sendToClient(clientId, ErrorMessage{err.Error()})
}

func (s *Server) HandleClientMessage(clientId uint64, message []byte) {
	var clientMessage ClientMessage
	if err := json.Unmarshal(message, &clientMessage); err != nil {
		s.sendErrorToClient(clientId, fmt.Errorf("Error unmarshaling client message: %w", err))
		return
	}

	messageHandler, ok := s.messageType2Handler[clientMessage.ActionType]
	if !ok {
		s.sendErrorToClient(clientId, fmt.Errorf("Unspported action"))
		return
	}

	response, err := messageHandler(clientId, clientMessage)
	if err != nil {
		s.sendErrorToClient(clientId, fmt.Errorf("Internal server error: %w", err))
		return
	}

	s.sendToClient(clientId, response)

}

func (s *Server) sendToClient(clientId uint64, message interface{}) {
	client, ok := s.clients[clientId]
	if !ok {
		log.Errorf("Client %d not found", clientId)
		return

	}
	bytes, err := json.Marshal(message)
	if err != nil {
		err = fmt.Errorf("Error marshaling message to json: %w, message=%v", err, message)
		log.Errorf(err.Error())
		bytes = []byte(err.Error())
	}
	client.SendMessage(&bytes)
}

func (s *Server) sendToAllClients(message interface{}) {
	for clientId, _ := range s.clients {
		s.sendToClient(clientId, message)
	}
}
