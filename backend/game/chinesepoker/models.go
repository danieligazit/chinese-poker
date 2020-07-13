package chinesepoker

import (
	"github.com/chehsunliu/poker"
)

const InitIterations = 4
const LastIteration = 5
const PlayerNumber = 2
const HandNumber = 5
const HandSize = 5

type ChinesePokerGame struct {
	deck                    *poker.Deck
	hands                   [PlayerNumber][HandNumber][]poker.Card
	iteration               int
	cardsInCurrentIteration int
	top                     poker.Card
	playerTurnIndex         int
	gameOver                bool
}

type ChinesePokerMove struct {
	HandIndex int
}
