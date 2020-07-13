package chinesepoker

import (
	"fmt"
	"github.com/chehsunliu/poker"
)

func NewChinesePokerGame() *ChinesePokerGame {
	var deck = poker.NewDeck()
	var game = ChinesePokerGame{
		deck:                    deck,
		hands:                   [PlayerNumber][HandNumber][]poker.Card{},
		top:                     deck.Draw(1)[0],
		iteration:               InitIterations + 1,
		cardsInCurrentIteration: 0,
		playerTurnIndex:         0,
		gameOver:                false,
	}

	for playerIndex := 0; playerIndex < PlayerNumber; playerIndex++ {
		for handIndex := 0; handIndex < HandSize; handIndex++ {
			game.hands[playerIndex][handIndex] = append(game.hands[playerIndex][handIndex], game.deck.Draw(InitIterations)...)
		}
	}

	return &game
}

func (g *ChinesePokerGame) checkGameOver() bool {
	return g.cardsInCurrentIteration == PlayerNumber*HandNumber && g.iteration == LastIteration
}

func (g *ChinesePokerGame) updateHands(move ChinesePokerMove) (legal bool, response interface{}) {

	if move.HandIndex >= len(g.hands[g.playerTurnIndex]) {
		response = fmt.Sprintf("Hand index %d exceeds hand number (must be 0-%d)", move.HandIndex, HandNumber-1)
		return
	}

	var hand = g.hands[g.playerTurnIndex][move.HandIndex]

	legal = len(hand) == g.iteration-1
	if !legal {
		response = fmt.Sprintf("Card already assigned to hand on index %d", move.HandIndex)
		return
	}

	g.hands[g.playerTurnIndex][move.HandIndex] = append(g.hands[g.playerTurnIndex][move.HandIndex], g.top)
	g.cardsInCurrentIteration++

	return
}

func (g *ChinesePokerGame) updateTurn() (gameOver bool, err error) {
	if g.deck.Empty() {
		err = fmt.Errorf("Card cannot be draw from deck as it is empty")
		return
	}

	g.top = g.deck.Draw(1)[0]

	g.playerTurnIndex = (g.playerTurnIndex + 1) % PlayerNumber

	if g.cardsInCurrentIteration == PlayerNumber*HandNumber {
		g.cardsInCurrentIteration = 0
		g.iteration++
	}

	return
}

func (g *ChinesePokerGame) MakeMove(playerIndex int, moveI interface{}) (legal bool, response interface{}, gameOver bool, err error) {
	if playerIndex != g.playerTurnIndex {
		response = fmt.Errorf("It is not player %d turn to play. Should be %d", playerIndex, g.playerTurnIndex)
		legal = false
		return
	}
	if g.gameOver {
		err = fmt.Errorf("Game is already over")
		return
	}
	move, ok := moveI.(ChinesePokerMove)
	if !ok {
		err = fmt.Errorf("Bad move format")
		return
	}

	legal, response = g.updateHands(move)
	if !legal {
		return
	}

	if gameOver = g.checkGameOver(); gameOver {
		g.gameOver = true
		return
	}

	gameOver, err = g.updateTurn()
	return
}

func (g *ChinesePokerGame) getResponseCards(requestingPlayerIndex int) (hands [PlayerNumber][HandNumber][]*poker.Card) {

	for playerIndex, playerHands := range g.hands {
		for handIndex, handCards := range playerHands {
			hands[playerIndex][handIndex] = make([]*poker.Card, len(handCards))
			for cardIndex, _ := range handCards {

				if cardIndex == LastIteration-1 && requestingPlayerIndex != playerIndex {
					continue
				}
				hands[playerIndex][handIndex][cardIndex] = &g.hands[playerIndex][handIndex][cardIndex]
			}
		}
	}
	return
}

func (g *ChinesePokerGame) isLastIteration() bool {
	return g.iteration == LastIteration
}

func (g *ChinesePokerGame) GetState(requestingPlayerIndex int) (state map[string]interface{}, err error) {
	if requestingPlayerIndex >= PlayerNumber {
		err = fmt.Errorf("Player index does not exists (must be 0-%d)", PlayerNumber-1)
		return
	}

	state = map[string]interface{}{
		"hands": g.getResponseCards(requestingPlayerIndex),
	}

	if requestingPlayerIndex == g.playerTurnIndex {
		state["top"] = g.top
	}

	return
}

func (g *ChinesePokerGame) GetResult() (result map[string]interface{}, err error) {
	if !g.gameOver {
		err = fmt.Errorf("Game is not over")
		return
	}

	handEvaluations := [PlayerNumber][HandNumber]map[string]interface{}{}
	handWinners := make([][]int, HandNumber)

	maxRanks := make([]int32, HandNumber)

	for playerIndex, playerHands := range g.hands {
		for handIndex, hand := range playerHands {
			rank := poker.Evaluate(hand)
			handEvaluations[playerIndex][handIndex] = map[string]interface{}{
				"rank":       rank,
				"rankString": poker.RankString(rank),
			}
			if rank == maxRanks[playerIndex] {
				handWinners[handIndex] = append(handWinners[handIndex], playerIndex)
			} else if rank > maxRanks[playerIndex] {
				handWinners[handIndex] = []int{playerIndex}
				maxRanks[handIndex] = rank
			}

		}
	}

	result = map[string]interface{}{
		"evaluations": handEvaluations,
		"winners":     handWinners,
	}
	return

}
