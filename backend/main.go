package main

import (
	"flag"
	// "fmt"
	"github.com/danieligazit/chinese-poker/backend/client"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, req *http.Request) {

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, req, nil)

	if err != nil {
		log.Println(err)
		return
	}

	client := client.NewClient(conn)
	// client.Listen()
	message := []byte(`hi`)
	client.SendMessage(&message)
	client.Listen()
	client.Listen()

}

func main() {
	flag.Parse()
	router := http.NewServeMux()
	router.HandleFunc("/", homeHandler)
	log.Printf("serving on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
