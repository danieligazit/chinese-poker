package communication

import (
	"fmt"
)

const makeMoveMessage = "makeMove"
const connectMessage = "connect"
const chat = "chat"

const stateResponse = "setState"
const gameOverResponse = "gameOver"
const errorResponse = "error"

type clientMessageHandler func(clientIndex uint64, clientMessage ClientMessage) (err error)

func (s *Server) connectHandler(clientId uint64, clientMessage ClientMessage) (err error) {
	state, err := s.game.GetState(s.clientId2Index[clientId])
	if err != nil {
		s.sendErrorToClient(clientId, fmt.Errorf("Error getting game state: %w", err))
		return
	}

	s.sendToClient(clientId, ClientMessage{stateResponse, state})
	return
}


func (s *Server) makeMoveHandler(clientId uint64, clientMessage ClientMessage) (err error) {
	legal, gameOver, response, err := s.game.MakeMove(s.clientId2Index[clientId], clientMessage.Action)
	if err != nil {
		s.sendErrorToClient(clientId, fmt.Errorf("Error getting game state: %w", err))
		return
	}
	
	s.sendToClient(clientId, response)
	if !legal{
		return
	}
	
	for clientId, playerIndex := range s.clientId2Index {
		state, err := s.game.GetState(playerIndex)
		if err != nil {
			s.sendErrorToClient(clientId, fmt.Errorf("Error getting game state: %w", err))
			continue
		}
		
		s.sendToClient(clientId, ClientMessage{stateResponse, state})
	}
	
	if !gameOver {
		return
	}
	
	result, err := s.game.GetResult()
	if err != nil {
		s.sendErrorToAllClients(fmt.Errorf("Error getting game results", err))
	}
	
	s.sendToAllClients(ClientMessage{gameOverResponse, result})
	return
}
