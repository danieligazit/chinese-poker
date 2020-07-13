package main

import (
	// "flag"
	// "fmt"
	"github.com/danieligazit/chinese-poker/backend/game/chinesepoker"
	// "github.com/gorilla/websocket"
	"encoding/json"
	"fmt"
)

func Interface2Object(origin interface{}, target interface{}) (err error) {
	byteData, err := json.Marshal(origin)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteData, &target)
}

func main() {
	g := chinesepoker.NewChinesePokerGame()

	fmt.Println(g.GetState(0))
	fmt.Println(g.GetState(1))

	player := 0
	for true {
		var handIndex int
		_, err := fmt.Scanf("%d", &handIndex)

		legal, response, gameOver, err := g.MakeMove(player, chinesepoker.ChinesePokerMove{
			HandIndex: handIndex,
		})
		fmt.Printf("legal=%t response=%v gameOver=%t err=%s\n", legal, response, gameOver, err)
		s1, _ := g.GetState(0)
		s2, _ := g.GetState(1)

		fmt.Println(s1)
		fmt.Println(s2)
		if gameOver {
			break
		}

		player = (player + 1) % 2
	}
	fmt.Println(g.GetResult())
	// flag.Parse()
	// router := http.NewServeMux()
	// router.HandleFunc("/", homeHandler)
	// log.Printf("serving on port 8081")
	// log.Fatal(http.ListenAndServe(":8081", router))
}
