package communication

import (
	"github.com/danieligazit/chinese-poker/backend/games"
	"github.com/gorilla/websocket"
)

const channelBufSize = 100

const PlayerIdURLKey = "clientId"

type Client struct {
	Id     uint64
	ws     *websocket.Conn
	ch     chan *[]byte
	doneCh chan bool
	server *Server
}

type Server struct {
	id                  string
	clients             map[uint64]*Client
	clientId2Index      map[uint64]int
	upgrader            *websocket.Upgrader
	game                games.IGame
	started             bool
	messageType2Handler map[string]clientMessageHandler
}

type ClientMessage struct {
	ActionType string      `json:"actionType"`
	Action     interface{} `json:"action"`
}
