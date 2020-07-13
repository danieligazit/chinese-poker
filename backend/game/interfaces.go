package game

type Game interface {
	MakeMove(playerIndex int, moveI interface{}) (legal bool, response interface{}, gameOver bool, err error)
	GetState(playerIndex int) (state map[string]interface{}, err error)
	GetResult(playerIndex int) (result map[string]interface{}, err error)
}
