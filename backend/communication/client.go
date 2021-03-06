package communication

import (
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

// NewClient initializes a new Client struct with given websocket.
func NewClient(clientId uint32, ws *websocket.Conn, lobby *Lobby) *Client {
	ch := make(chan *[]byte, channelBufSize)
	doneCh := make(chan bool)

	return &Client{clientId, ws, ch, doneCh, lobby}
}

// Listen Write and Read request via chanel
func (c *Client) Listen() {
	go c.listenWrite()
	c.listenRead()
}

// Listen write request via chanel
func (c *Client) listenWrite() {
	defer func() {
		err := c.ws.Close()
		if err != nil {
			log.Errorf(fmt.Errorf("Cannot close websocket: %w", err).Error())
		}
	}()

	log.Info("Listening write to client")
	for {
		select {

		case bytes := <-c.ch:
			err := c.ws.WriteMessage(websocket.BinaryMessage, *bytes)

			if err != nil {
				log.Errorf("Error writing message to websocket: %w", err)
			}

		case <-c.doneCh:
			c.doneCh <- true
			return
		}
	}
}

func (c *Client) listenRead() {
	defer func() {
		err := c.ws.Close()
		if err != nil {
			log.Println("Error:", err.Error())
		}
	}()

	log.Println("Listening read from client")
	for {
		select {

		case <-c.doneCh:
			c.doneCh <- true
			return

		default:
			c.readFromWebSocket()
		}
	}
}

// SendMessage sends game state to the client.
func (c *Client) SendMessage(bytes *[]byte) {
	select {
	case c.ch <- bytes:
	}
}

// Done sends done message to the Client which closes the conection.
func (c *Client) Done() {
	c.doneCh <- true
}

func (c *Client) readFromWebSocket() {
	messageType, data, err := c.ws.ReadMessage()
	if err != nil {
		log.Errorf(err.Error())
		c.doneCh <- true
		c.lobby.HandleClientDissconnect(c.Id)
	} else if messageType != websocket.BinaryMessage {
		log.Errorf("Non binary message recived, ignoring")
	} else {
		c.lobby.HandleClientMessage(c.Id, data)
	}
}
