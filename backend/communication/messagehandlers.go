package communication

const makeMoveMessage = "makeMove"
const connectMessage = "connect"
const chat = "chat"

type clientMessageHandler func(clientMessage interface{}, clientIndex int) (response interface{}, err error)

func (s *Server) makeMoveHandler(clientId uint32, clientMessage interface{}) (response interface{}, err error) {
	responseO, err = s.game.GetState(s.clientId2Index[clientId])
	if err != nil {
		err = Errorf("Error getting game state: %w", err)
		return
	}

	response = interface{}(responseO)

	return
}
