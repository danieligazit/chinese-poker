package communication

import (
	"fmt"
	"github.com/danieligazit/chinese-poker/backend/games"
	log "github.com/sirupsen/logrus"
)

type clientMessageHandler func(clientIndex uint64, clientMessage ClientMessage) (err error)

type MessageHandler struct {
	lobby               *Lobby
	game                *games.IGame
	clientId2Index      map[uint64]int
	messageType2Handler map[string]clientMessageHandler
	active              bool
	marshaler           marshalHandler
}

func NewMessageHandler(game *games.IGame, lobby *Lobby, proto protocol) *MessageHandler {
	mh := MessageHandler{
		lobby:          lobby.lobby,
		game:           game,
		clientId2Index: map[uint64]int{},
		active:         false,
		marshaler:      protocol2marshalHandler[proto],
	}

	mh.messageType2Handler = map[uint64]clientMessageHandler{
		makeMoveMessage: mh.makeMoveHandler,
		connectMessage:  mh.userConnectHandler,
		// 		chatMessage:     server.chatHandler,
	}
}

func (mh *MessageHandler) unmarshalClientMessage(message []byte) (clientmessage ClientMessage, err error) {
	var clientMessage ClientMessage
	if err := mh.marshaler.unmarshal(message, &clientMessage); err != nil {
		err = fmt.Errorf("Error unmarshaling client message: %w", err)
		return
	}

	return
}

func (mh *MessageHandler) handleClientMessage(clientId uint64, message []byte) {
	clientMessage, err := mh.unmarshalClientMessage(message)
	if err != nil {
		log.Errorf("Error handling client message: %w", err)
	}

	messageHandler, ok := mh.messageType2Handler[clientMessage.ActionType]
	if !ok {
		return
	}

	if err := messageHandler(clientId, clientMessage); err != nil {
		log.Errorf("Internal server error: %w", err)
	}

	return
}

func (mh *MessageHandler) handleUserConnect(clientId uint64, clientMessage ClientMessage) (err error) {
	if _, exists := s.clientId2Index[clientId]; !exists {
		mh.lobby.clientId2Index[clientId] = len(lobby.clientId2Index)
	}

	log.Infof("Added new client with id %d. There are currently %d clients connected", clientId, len(s.clients))
	mh.sendConnectionStatus()
	mh.sendGameState(clientId)

	if minPlayers, _ := lobby.game.GetPlayerNum(); len(s.clientId2Index) >= minPlayers && !lobby.started {
		s.startGame()
	}

	return
}

func (lobby *Lobby) makeMoveHandler(clientId uint64, clientMessage ClientMessage) (err error) {
	if !s.started {
		s.sendErrorToClient(clientId, fmt.Errorf("Game hasn't started yet"))
		return
	}

	legal, gameOver, response, err := lobby.game.MakeMove(s.clientId2Index[clientId], clientMessage.Action)
	if err != nil {
		lobby.sendErrorToClient(clientId, fmt.Errorf("Error getting game state: %w", err))
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
		lobby.sendErrorToAllClients(fmt.Errorf("Error getting game results", err))
	}

	lobby.sendToAllClients(ClientMessage{gameOverResponse, result})
	return
}

func (mh *MessageHandler) sendGameState(clientId uint64) {
	playerIndex, ok := ms.clientId2Index[clientId]
	if !ok {
		return
	}

	state, err := mh.game.GetState(playerIndex)
	if err != nil {
		log.Errorf("Error getting game state: %w", err)
		return
	}

	s.sendToClient(clientId, ClientMessage{stateResponse, state})
}

func (mh *MessageHandler) sendConnectionStatus() {
	clientIds := []string{}
	for clientId, _ := range s.clients {
		clientIds = append(clientIds, clientId)
	}
	minPlayers, maxPlayers := s.game.GetPlayerNum()
	mh.sendToAllClients(ClientMessage{ActionType: clientConnectionStatusResponse, Action: ConnectionStatus{clientIds, s.active, minPlayers, maxPlayers}})
}

func (s *Lobby) startGame() {
	s.active = true
	s.sendToAllClients(ClientMessage{ActionType: startGameResponse})
}
