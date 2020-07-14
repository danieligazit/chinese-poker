package communication

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

const connectMessage = "connect"
const makeMoveMessage = "makeMove"
const chatMessage = "chat"

const startGameResponse = "startGame"
const clientConnectionStatusResponse = "clientConnectionStatus"
const makeMoveResponse = "makeMoveResponse"
const stateResponse = "setState"
const gameOverResponse = "gameOver"
const errorResponse = "error"

type clientMessageHandler func(clientIndex uint64, clientMessage ClientMessage) (err error)

func (s *Server) userConnectHandlers(clientId uint64, clientMessage ClientMessage) (err error) {
	if _, exists := s.clientId2Index[clientId]; !exists {
		s.clientId2Index[clientId] = len(s.clientId2Index)
	}

	log.Infof("Added new client with id %d. There are currently %d clients connected", clientId, len(s.clients))
	s.sendConnectionStatus()
	s.sendGameState(clientId)

	if minPlayers, _ := s.game.GetPlayerNum(); len(s.clientId2Index) >= minPlayers && !s.started {
		s.startGame()
	}

	return
}

func (s *Server) makeMoveHandler(clientId uint64, clientMessage ClientMessage) (err error) {
	if !s.started {
		s.sendErrorToClient(clientId, fmt.Errorf("Game hasn't started yet"))
		return
	}

	legal, gameOver, response, err := s.game.MakeMove(s.clientId2Index[clientId], clientMessage.Action)
	if err != nil {
		s.sendErrorToClient(clientId, fmt.Errorf("Error getting game state: %w", err))
		return
	}

	s.sendToClient(clientId, ClientMessage{ActionType: makeMoveResponse, Action: response})
	if !legal {
		return
	}

	for clientId, _ := range s.clients {
		s.sendGameState(clientId)
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

func (s *Server) sendGameState(clientId uint64) {
	playerIndex, ok := s.clientId2Index[clientId]
	if !ok {
		return
	}
	state, err := s.game.GetState(playerIndex)
	if err != nil {
		s.sendErrorToClient(clientId, fmt.Errorf("Error getting game state: %w", err))
		return
	}

	s.sendToClient(clientId, ClientMessage{stateResponse, state})
}
