package common

type IGame interface {
	MakeMove(playerIndex int, moveI interface{}) (legal, gameOver bool, response interface{}, err error)
	GetState(playerIndex int) (state interface{}, err error)
	GetResult() (result interface{}, err error)
	GetPlayerNum() (min int, max int)
	GetGameName() string
}
