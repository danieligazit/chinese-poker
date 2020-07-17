package games

type IGame interface {
	MakeMove(playerIndex uint32, moveI interface{}) (legal, gameOver bool, response interface{}, err error)
	GetState(playerIndex uint32) (state interface{}, err error)
	GetResult(playerIndex uint32) (result interface{}, err error)
	GetPlayerNum() (min int, max int)
	GetGameName() string
}
