package games

import (
	"fmt"
	"github.com/danieligazit/chinesepoker/backend/games/chinesepoker"
)

const Poker = "poker"

type gameConstructor func(interface{}) *IGame

var game2Constructor = map[string]game2Constructor{
	Poker: chinesepoker.NewGame,
}

func NewGame(gameStr string, params interface{}) (game *IGame, err error) {
	constructor, exists := game2Constructor[gameStr]
	if !exists {
		err = fmt.Errorf("Game %s does not exists", gameStr)
	}

	game = constructor(params)
	return
}
