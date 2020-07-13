package games

type IGame interface {
	MakeMove(playerIndex int, moveI interface{}) (response interface{}, err error)
	GetState(playerIndex int) (state interface{}, err error)
	GetResult() (result map[string]interface{}, err error)
	GetPlayerNum() (int, int)
}
