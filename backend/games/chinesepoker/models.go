package chinesepoker

import (
	"github.com/chehsunliu/poker"
)

const InitIterations = 4
const LastIteration = 5
const PlayerNumber = 2
const HandNumber = 5
const HandSize = 5

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
	Message string `josn:"message"`
}

type State struct {
	Top           *poker.Card                             `json:"top"`
	Hands         [PlayerNumber][HandNumber][]*poker.Card `json:"hands"`
	IsCurrentTurn bool                                    `json:"isCurrentTurn"`
	IsGameOver    bool                                    `json:"isGameOver"`
}
