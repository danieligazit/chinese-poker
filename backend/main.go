package main

import (
	// "flag"
	// "fmt"
	// "github.com/danieligazit/chinese-poker/backend/client"
	// "github.com/gorilla/websocket"
	"encoding/json"
	"fmt"
	"github.com/chehsunliu/poker"
)

const InitIterations = 1
const PlayerNumber = 2
const HandNumber = 5
const HandSize = 5

func Interface2Object(origin interface{}, target interface{}) (err error) {
	byteData, err := json.Marshal(origin)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteData, &target)
}

type IGame interface {
	MakeMove(playerIndex int) (legal bool, response interface{}, isGameOver bool)
	GetState() interface{}
	GetResult() interface{}
}

type ChinesePokerGame struct {
	deck            *poker.Deck
	hands           [PlayerNumber][HandNumber][]poker.Card
	iteration       int
	top             poker.Card
	playerTurnIndex int
}

func NewChinesePokerGame() *ChinesePokerGame {
	var deck = poker.NewDeck()
	var game = ChinesePokerGame{
		deck:            deck,
		hands:           [PlayerNumber][HandNumber][]poker.Card{},
		top:             deck.Draw(1)[0],
		iteration:       1,
		playerTurnIndex: 0,
	}

	for i := 0; i < InitIterations; i++ {
		for playerIndex := 0; playerIndex < PlayerNumber; playerIndex++ {
			for handIndex := 0; handIndex < HandSize; handIndex++ {
				game.hands[playerIndex][handIndex] = append(game.hands[playerIndex][handIndex], game.deck.Draw(1)...)
			}
		}
	}

	return &game

}

type ChinesePokerMove struct {
	HandIndex int
}

type ChinesePokerState struct {
	hands [PlayerNumber][HandNumber][]string
	top   poker.Card
}

func (g *ChinesePokerGame) updateHands(move ChinesePokerMove) (legal bool, response interface{}) {

	if move.HandIndex >= len(g.hands[g.playerTurnIndex]) {
		response = fmt.Sprintf("Hand index %d exceeds hand number (must be 0-%d)", move.HandIndex, HandNumber-1)
		return
	}

	var hand = g.hands[g.playerTurnIndex][move.HandIndex]

	legal = len(hand) == g.iteration
	if !legal {
		response = fmt.Sprintf("Card already assigned to hand on index %d", move.HandIndex)
		return
	}

	g.hands[g.playerTurnIndex][move.HandIndex] = append(g.hands[g.playerTurnIndex][move.HandIndex], g.top)
	return
}

func (g *ChinesePokerGame) updateTurn() (gameOver bool) {
	g.top = g.deck.Draw(1)[0]
	g.playerTurnIndex++

	if g.playerTurnIndex == PlayerNumber {
		g.playerTurnIndex = 0
	}

	return g.deck.Empty()
}

func (g *ChinesePokerGame) MakeMove(playerIndex int, moveI interface{}) (legal bool, response interface{}, gameOver bool, err error) {
	move, ok := moveI.(ChinesePokerMove)
	if !ok {
		var responseStr = "Bad move format"
		response = responseStr
		err = fmt.Errorf(responseStr)
		return
	}

	legal, response = g.updateHands(move)
	if !legal {
		return
	}

	gameOver = g.updateTurn()
	return
}

func (g *ChinesePokerGame) GetState(playerIndex int) (state map[string]interface{}, err error) {
	if playerIndex >= PlayerNumber {
		err = fmt.Errorf("Player index does not exists (must be 0-%d)", PlayerNumber-1)
		return
	}

	state = map[string]interface{}{
		"hands": g.hands,
	}

	if playerIndex == g.playerTurnIndex {
		state["top"] = g.top
	}

	return
}

func (g *ChinesePokerGame) GetResult() (result interface{}, err error) {
	return
}

func main() {
	g := NewChinesePokerGame()

	fmt.Println(g.GetState(0))
	fmt.Println(g.GetState(1))

	var (
		gameOver bool
	)

	for !gameOver {
		var handIndex int
		_, err := fmt.Scanf("%d", &handIndex)

		legal, response, gameOver, err := g.MakeMove(0, ChinesePokerMove{
			HandIndex: handIndex,
		})
		fmt.Printf("legal=%t response=%v gameOver=%t err=%s\n", legal, response, gameOver, err)
		s1, _ := g.GetState(0)
		s2, _ := g.GetState(1)

		fmt.Println(s1)
		fmt.Println(s2)
	}
	// flag.Parse()
	// router := http.NewServeMux()
	// router.HandleFunc("/", homeHandler)
	// log.Printf("serving on port 8081")
	// log.Fatal(http.ListenAndServe(":8081", router))
}
