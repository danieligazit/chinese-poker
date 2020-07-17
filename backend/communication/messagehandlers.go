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

type clientMessageHandler func(clientIndex uint32, clientMessage ClientMessage) (err error)

func (s *Lobby) userConnectHandler(clientId uint32, clientMessage ClientMessage) (err error) {
	log.Infof("Added new client with id %d. There are currently %d clients connected", clientId, len(s.clients))
	s.sendConnectionStatus()
	s.sendGameState(clientId)

	if minPlayers, _ := s.game.GetPlayerNum(); len(s.clients) >= minPlayers && !s.started {
		s.startGame()
	}

	return
}

func (s *Lobby) makeMoveHandler(clientId uint32, clientMessage ClientMessage) (err error) {
	if !s.started {
		s.sendErrorToClient(clientId, fmt.Errorf("Game hasn't started yet"))
		return
	}

	legal, gameOver, response, err := s.game.MakeMove(clientId, clientMessage.Action)
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

	for clientId, _ := range s.clients {
		result, err := s.game.GetResult(clientId)
		if err != nil {
			log.Errorf("Error getting result clientId %d: %w", clientId, err)
		}
		s.sendToClient(clientId, ClientMessage{gameOverResponse, result})
	}

	return
}

func (s *Lobby) sendGameState(clientId uint32) {
	state, err := s.game.GetState(clientId)
	if err != nil {
		s.sendErrorToClient(clientId, fmt.Errorf("Error getting game state: %w", err))
		return
	}

	s.sendToClient(clientId, ClientMessage{stateResponse, state})
}
