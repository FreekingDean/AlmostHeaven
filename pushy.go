package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func PushySocket(w http.ResponseWriter, r *http.Request) {
	scheme := r.Header.Get("X-Forward-Scheme")
	log.Println(scheme)
	r.Header["Connection"] = []string{"upgrade"}
	r.Header["Upgrade"] = []string{"websocket"}
	client, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer client.Close()

	for {
		mt, message, err := client.ReadMessage()
		if err != nil {
			log.Println("CLIENT ERROR: ", err)
			return
		}
		log.Printf("CLIENT GET: %d -- %s\n", mt, message)
	}
}
