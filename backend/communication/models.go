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
	lobby  *Lobby
}

type Lobby struct {
	id               string
	clients          map[uint64]*Client
	upgrader         *websocket.Upgrader
	game             games.IGame
	userNameRegistry *utils.UserNamesRegistry
	messageHandler   *MessageHandler
}
