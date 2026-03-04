package main

import (
	"fang_wars_server/internal/game"
	"fang_wars_server/internal/network"
	"log"
	"net/http"
)

func main() {
http.HandleFunc("/ws",network.ConnRequestHandler )
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Fang_wars_client/src/index.html")
	})
	go game.GameLoop()
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
