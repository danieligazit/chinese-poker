package chinesepoker

import (
	"github.com/chehsunliu/poker"
)

const InitIterations = 1
const LastIteration = 5
const PlayerNumber = 2
const HandNumber = 5
const HandSize = 5

const GameName = "chinese-poker"

type Game struct {
	deck                    *poker.Deck
	hands                   [PlayerNumber][HandNumber][]poker.Card
	iteration               int
	cardsInCurrentIteration int
	top                     *poker.Card
	playerTurnIndex         int
	gameOver                bool
}

type Move struct {
	HandIndex int `json:"handIndex"`
}

type MoveResponse struct {
	Legal   bool   `json:"legal"`
	Message string `json:"message"`
}

type State struct {
	Top           *poker.Card                             `json:"top"`
	Hands         [PlayerNumber][HandNumber][]*poker.Card `json:"hands"`
	IsCurrentTurn bool                                    `json:"isCurrentTurn"`
	Iteration     int                                     `json:"iteration"`
	PlayerIndex   int                                     `json:"playerIndex"`
}

type Evaluation struct {
	Rank    int32  `json:"rank"`
	RankStr string `json:"rankStr"`
}
type Result struct {
	Evaluations [PlayerNumber][HandNumber]Evaluation `json:"evaluations"`
	Winners     [][]int                              `winners`
}
