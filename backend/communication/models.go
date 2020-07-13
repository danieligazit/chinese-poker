package communication

import (
	"github.com/gorilla/websocket"
	"github.com/danieligazit/chinese-poker/backend/games"
)

const channelBufSize = 100

const PasswordURLKey = "password"
const PlayerIdURLKey = "clientId"
const BaseURI = "/%s"


type Client struct {
	Id     uint64
	ws     *websocket.Conn
	ch     chan *[]byte
	doneCh chan bool
	server *Server
}

type Server struct {
	Id                  string
	password            string
	clients             map[uint64]*Client
	clientId2Index      map[uint64]int
	upgrader            *websocket.Upgrader
	game                games.IGame
	messageType2Handler map[string]clientMessageHandler
}

type ClientMessage struct {
	ActionType string      `json:"actionType"`
	Action     interface{} `json:"action"`
}


