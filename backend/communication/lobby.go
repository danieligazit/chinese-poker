package communication

import (
	"encoding/json"
	"fmt"
	"github.com/danieligazit/chinese-poker/backend/games"
	"github.com/danieligazit/chinese-poker/backend/utility"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func NewLobby(id string, game games.IGame, usernameRegistry utility.UserNameRegistry) *Lobby {
	var lobby = Lobby{
		id:             id,
		clients:        make(map[uint32]*Client),
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		usernameRegistry: usernameRegistry,
		game:    game,
		started: false,
	}
	lobby.messageType2Handler = map[string]clientMessageHandler{
		makeMoveMessage: lobby.makeMoveHandler,
		connectMessage:  lobby.userConnectHandler,
		// 		chatMessage:     lobby.chatHandler,
	}

	uri := fmt.Sprintf("/%s/%s", lobby.game.GetGameName(), id)
	http.HandleFunc(uri, lobby.handler)
	log.Infof("New lobby at endpoint %s", uri)
	return &lobby
}

func (s *Lobby) handler(w http.ResponseWriter, req *http.Request) {
	utility.SetupResponseCORS(&w, req)

	if _, maxClientNumber := s.game.GetPlayerNum(); len(s.clients) >= maxClientNumber {
		http.Error(w, "Lobby is full", http.StatusLocked)
		return
	}

	user := req.URL.Query().Get(UsernameKey)
	if user == "null" {
		log.Errorf("Got empty username")
		http.Error(w, "Username cannot be empty", http.StatusBadRequest)
		return
	}
	
	clientId := s.usernameRegistry.AddUserName(user)
	if _, maxPlayers := s.game.GetPlayerNum(); int(clientId) >= maxPlayers{
		http.Error(w, "Lobby is full", http.StatusLocked)
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

func (s *Lobby) startGame() {
	s.started = true
	s.sendToAllClients(ClientMessage{ActionType: startGameResponse})
}

func (s *Lobby) sendErrorToClient(clientId uint32, err error) {
	log.Errorf(err.Error())
	s.sendToClient(clientId, ClientMessage{errorResponse, err.Error()})
}

func (s *Lobby) sendErrorToAllClients(err error) {
	log.Errorf(err.Error())
	s.sendToAllClients(ClientMessage{errorResponse, err.Error()})
}

func (s *Lobby) HandleClientMessage(clientId uint32, message []byte) {
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

func (s *Lobby) HandleClientDissconnect(clientId uint32) {
	delete(s.clients, clientId)
	s.sendConnectionStatus()
}

func (s *Lobby) sendToClient(clientId uint32, message ClientMessage) {
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

func (s *Lobby) sendToAllClients(message ClientMessage) {
	for clientId, _ := range s.clients {
		s.sendToClient(clientId, message)
	}
}

func (s *Lobby) sendConnectionStatus() {
	usernames := []string{}
	for clientId, _ := range s.clients {
		username, _ := s.usernameRegistry.GetUserName(clientId)
		usernames = append(usernames, username)
	}
	minPlayers, maxPlayers := s.game.GetPlayerNum()
	s.sendToAllClients(ClientMessage{ActionType: clientConnectionStatusResponse, Action: ConnectionStatus{usernames, s.started, minPlayers, maxPlayers}})
}
