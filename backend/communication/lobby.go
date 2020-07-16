package communication

import (
	"encoding/json"
	"fmt"
	"github.com/danieligazit/chinese-poker/backend/games"
	"github.com/danieligazit/chinese-poker/backend/utility"
	"github.com/gorilla/websocket"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

func NewLobby(id string, game games.IGame, proto protocol) *Lobby {
	var lobby = Lobby{
		id:      id,
		clients: map[uint64]*Client{},

		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		game: game,
	}

	uri := fmt.Sprintf("/%s/%s", server.game.GetGameName(), id)
	http.HandleFunc(uri, server.handler)
	log.Infof("New server at endpoint %s", uri)
	return &lobby
}

func (s *Lobby) handler(w http.ResponseWriter, req *http.Request) {
	utility.SetupResponseCORS(&w, req)

	if _, maxClientNumber := s.game.GetPlayerNum(); len(s.clients) >= maxClientNumber {
		http.Error(w, "Server is full", http.StatusLocked)
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

func (s *Lobby) HandleClientDissconnect(clientId uint64) {
	delete(s.clients, clientId)
	s.sendConnectionStatus()
}

func (l *Lobby) SendToClient(clientId uint64, bytes []byte) {
	client, ok := l.clients[clientId]
	if !ok {
		log.Errorf("Client %d not found", clientId)
		return

	}

	client.SendMessage(&bytes)
}

func (s *Lobby) SendToAllClients(message []byte) {
	for clientId, _ := range s.clients {
		s.sendToClient(clientId, message)
	}
}
