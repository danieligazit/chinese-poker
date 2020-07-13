package game

type IGame interface {
	MakeMove(playerIndex int) (legal bool, response interface{}, isGameOver bool)
	GetState() interface{}
	GetResult() interface{}
}