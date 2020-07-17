package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/danieligazit/chinese-poker/backend/communication"
	"github.com/danieligazit/chinese-poker/backend/games"
	"github.com/danieligazit/chinese-poker/backend/games/chinesepoker"
	"github.com/danieligazit/chinese-poker/backend/utility"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type NewLobbyResponse struct {
	LobbyId string `json:"lobbyId"`
	URL     string `json:"url"`
}

func newLobby(w http.ResponseWriter, req *http.Request) {
	utility.SetupResponseCORS(&w, req)

	gameStr := req.URL.Path[len("/new/"):]
	game, err := NewGame(gameStr, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	serverId := xid.New().String()
	communication.NewLobby(serverId, game, utility.NewUserNameRegistryMap())

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(NewLobbyResponse{serverId, fmt.Sprintf("/%s/%s", gameStr, serverId)})

}

type gameConstructor func(interface{}) games.IGame

var game2Constructor = map[string]gameConstructor{
	chinesepoker.GameName: chinesepoker.NewGame,
}

func NewGame(gameStr string, params interface{}) (game games.IGame, err error) {
	constructor, exists := game2Constructor[gameStr]
	if !exists {
		err = fmt.Errorf("Game %s does not exists", gameStr)
		return
	}

	game = constructor(params)
	return
}

func main() {
	http.HandleFunc("/new/", newLobby)
	flag.Parse()

	log.Info("serving on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
