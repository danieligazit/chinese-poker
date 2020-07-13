package main

import (
	"flag"
	"github.com/danieligazit/chinese-poker/backend/communication"
	"github.com/danieligazit/chinese-poker/backend/games/chinesepoker"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	game := chinesepoker.NewChinesePokerGame()
	communication.NewServer("test", "", game)
	// fmt.Println(g.GetState(0))
	// fmt.Println(g.GetState(1))

	// player := 0
	// for true {
	// 	var handIndex int
	// 	_, err := fmt.Scanf("%d", &handIndex)

	// 	legal, response, gameOver, err := g.MakeMove(player, chinesepoker.ChinesePokerMove{
	// 		HandIndex: handIndex,
	// 	})
	// 	fmt.Printf("legal=%t response=%v gameOver=%t err=%s\n", legal, response, gameOver, err)
	// 	s1, _ := g.GetState(0)
	// 	s2, _ := g.GetState(1)

	// 	fmt.Println(s1)
	// 	fmt.Println(s2)
	// 	if gameOver {
	// 		break
	// 	}

	// 	player = (player + 1) % 2
	// }
	// fmt.Println(g.GetResult())
	flag.Parse()

	log.Info("serving on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
