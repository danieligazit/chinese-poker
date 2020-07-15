package communication

import (
	"encoding/json"
	"fmt"
	"github.com/danieligazit/chinese-poker/backend/games"
	"github.com/danieligazit/chinese-poker/backend/utility"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func NewServer(id string, game games.IGame) *Server {
	var server = Server{
		id:             id,
		clients:        map[uint64]*Client{},
		clientId2Index: map[uint64]int{},
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		game:    game,
		started: false,
	}
	server.messageType2Handler = map[string]clientMessageHandler{
		makeMoveMessage: server.makeMoveHandler,
		connectMessage:  server.userConnectHandler,
		// 		chatMessage:     server.chatHandler,
	}

	uri := fmt.Sprintf("/%s/%s", server.game.GetGameName(), id)
	http.HandleFunc(uri, server.handler)
	log.Infof("New server at endpoint %s", uri)
	return &server
}

func (s *Server) handler(w http.ResponseWriter, req *http.Request) {
	utility.SetupResponseCORS(&w, req)

	if _, maxClientNumber := s.game.GetPlayerNum(); len(s.clients) >= maxClientNumber {
		http.Error(w, "Server is full", http.StatusLocked)
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

	client.Listen()
}

func (s *Server) startGame() {
	s.started = true
	s.sendToAllClients(ClientMessage{ActionType: startGameResponse})
}

func (s *Server) sendErrorToClient(clientId uint64, err error) {
	log.Errorf(err.Error())
	s.sendToClient(clientId, ClientMessage{errorResponse, err.Error()})
}

func (s *Server) sendErrorToAllClients(err error) {
	log.Errorf(err.Error())
	s.sendToAllClients(ClientMessage{errorResponse, err.Error()})
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

	if err := messageHandler(clientId, clientMessage); err != nil {
		s.sendErrorToClient(clientId, fmt.Errorf("Internal server error: %w", err))
	}

	return

}

func (s *Server) HandleClientDissconnect(clientId uint64) {
	delete(s.clients, clientId)
	s.sendConnectionStatus()
}

func (s *Server) sendToClient(clientId uint64, message ClientMessage) {
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

func (s *Server) sendToAllClients(message ClientMessage) {
	for clientId, _ := range s.clients {
		s.sendToClient(clientId, message)
	}
}

func (s *Server) sendConnectionStatus() {
	clientIds := []uint64{}
	for clientId, _ := range s.clients {
		clientIds = append(clientIds, clientId)
	}
	minPlayers, maxPlayers := s.game.GetPlayerNum()
	s.sendToAllClients(ClientMessage{ActionType: clientConnectionStatusResponse, Action: ConnectionStatus{clientIds, s.started, minPlayers, maxPlayers}})
}
