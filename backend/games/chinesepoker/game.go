package chinesepoker

import (
	"fmt"
	"github.com/chehsunliu/poker"
	"github.com/danieligazit/chinese-poker/backend/games"
	"github.com/danieligazit/chinese-poker/backend/utility"
)

func NewGame(params interface{}) games.IGame {
	deck := poker.NewDeck()
	card := deck.Draw(1)[0]

	var game = Game{
		deck:                    deck,
		hands:                   [PlayerNumber][HandNumber][]poker.Card{},
		top:                     &card,
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

func (g *Game) GetGameName() string {
	return GameName
}

func (g *Game) checkGameOver() bool {
	return g.cardsInCurrentIteration == PlayerNumber*HandNumber && g.iteration == LastIteration
}

func (g *Game) updateHands(move Move) (legal bool, message string) {

	if move.HandIndex >= len(g.hands[g.playerTurnIndex]) {
		message = fmt.Sprintf("Hand index %d exceeds hand number (must be 0-%d)", move.HandIndex, HandNumber-1)
		return
	}

	var hand = g.hands[g.playerTurnIndex][move.HandIndex]

	legal = len(hand) == g.iteration-1
	if !legal {
		message = fmt.Sprintf("Card already assigned to hand on index %d", move.HandIndex)
		return
	}

	g.hands[g.playerTurnIndex][move.HandIndex] = append(g.hands[g.playerTurnIndex][move.HandIndex], *g.top)
	g.cardsInCurrentIteration++

	return
}

func (g *Game) GetPlayerNum() (int, int) {
	return PlayerNumber, PlayerNumber
}

func (g *Game) updateTurn() (err error) {
	if g.deck.Empty() {
		err = fmt.Errorf("Card cannot be draw from deck as it is empty")
		return
	}

	card := g.deck.Draw(1)[0]
	g.top = &card

	g.playerTurnIndex = (g.playerTurnIndex + 1) % PlayerNumber

	if g.cardsInCurrentIteration == PlayerNumber*HandNumber {
		g.cardsInCurrentIteration = 0
		g.iteration++
	}

	return
}

func (g *Game) MakeMove(playerIndex int, moveI interface{}) (legal, gameOver bool, response interface{}, err error) {
	moveResponse := MoveResponse{}
	if playerIndex != g.playerTurnIndex {
		moveResponse.Message = fmt.Sprintf("It is not player %d turn to play. Should be %d", playerIndex, g.playerTurnIndex)
		moveResponse.Legal = false
		response = moveResponse
		return
	}
	if g.gameOver {
		err = fmt.Errorf("Game is already over")
		return
	}
	var move Move
	err = utility.Interface2Object(moveI, &move)
	if err != nil {
		err = fmt.Errorf("Bad move format: %w", err)
		return
	}

	if moveResponse.Legal, moveResponse.Message = g.updateHands(move); !moveResponse.Legal {
		response = moveResponse
		return
	}

	legal = true
	response = moveResponse

	gameOver = g.checkGameOver()
	g.gameOver = gameOver

	return
}

func (g *Game) getResponseCards(requestingPlayerIndex int) (hands [PlayerNumber][HandNumber][]*poker.Card) {

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

func (g *Game) isLastIteration() bool {
	return g.iteration == LastIteration
}

func (g *Game) GetState(requestingPlayerIndex int) (state interface{}, err error) {
	if requestingPlayerIndex >= PlayerNumber {
		err = fmt.Errorf("Player index does not exists (must be 0-%d)", PlayerNumber-1)
		return
	}

	curState := State{
		Hands:         g.getResponseCards(requestingPlayerIndex),
		IsCurrentTurn: false,
		Iteration:     g.iteration,
		PlayerIndex:   requestingPlayerIndex,
	}

	if requestingPlayerIndex == g.playerTurnIndex {
		curState.Top = g.top
		curState.IsCurrentTurn = true
	}

	state = curState

	return
}

func (g *Game) GetResult() (result interface{}, err error) {
	if !g.gameOver {
		err = fmt.Errorf("Game is not over")
		return
	}

	handEvaluations := [PlayerNumber][HandNumber]Evaluation{}
	handWinners := make([][]int, HandNumber)

	maxRanks := make([]int32, HandNumber)

	for playerIndex, playerHands := range g.hands {
		for handIndex, hand := range playerHands {
			rank := poker.Evaluate(hand)

			handEvaluations[playerIndex][handIndex] = Evaluation{
				Rank:    rank,
				RankStr: poker.RankString(rank),
			}

			if rank == maxRanks[playerIndex] {
				handWinners[handIndex] = append(handWinners[handIndex], playerIndex)
			} else if rank > maxRanks[playerIndex] {
				handWinners[handIndex] = []int{playerIndex}
				maxRanks[handIndex] = rank
			}
		}
	}

	result = Result{
		Evaluations: handEvaluations,
		Winners:     handWinners,
	}
	return

}
