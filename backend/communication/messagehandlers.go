package communication

import (
	"fmt"
)

const makeMoveMessage = "makeMove"
const connectMessage = "connect"
const chat = "chat"

const stateResponse = "setState"
const errorResponse = "error"

type clientMessageHandler func(clientIndex uint64, clientMessage ClientMessage) (response ClientMessage, err error)

func (s *Server) connectHandler(clientId uint64, clientMessage ClientMessage) (response ClientMessage, err error) {
	state, err := s.game.GetState(s.clientId2Index[clientId])
	if err != nil {
		err = fmt.Errorf("Error getting game state: %w", err)
		return
	}

	response = ClientMessage{stateResponse, state}

	return
}

func (s *Server) makeMoveHandler(clientId uint64, clientMessage ClientMessage) (response ClientMessage, err error) {
	return
}
