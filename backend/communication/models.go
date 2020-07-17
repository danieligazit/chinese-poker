package communication

import (
	"github.com/danieligazit/chinese-poker/backend/games"
	"github.com/danieligazit/chinese-poker/backend/utility"
	"github.com/gorilla/websocket"
)

const channelBufSize = 100

const UsernameKey = "username"

type Client struct {
	Id     uint32
	ws     *websocket.Conn
	ch     chan *[]byte
	doneCh chan bool
	lobby  *Lobby
}

type Lobby struct {
	id                  string
	clients             map[uint32]*Client
	upgrader            *websocket.Upgrader
	game                games.IGame
	started             bool
	messageType2Handler map[string]clientMessageHandler
	usernameRegistry    utility.UserNameRegistry
}

type ClientMessage struct {
	ActionType string      `json:"actionType"`
	Action     interface{} `json:"action"`
}

type ConnectionStatus struct {
	Usernames  []string `json:"usernames"`
	Active     bool     `json:"active"`
	MinPlayers int      `json:"minPlayers"`
	MaxPlayers int      `json:"maxPlayers"`
}
