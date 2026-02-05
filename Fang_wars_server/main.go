package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

type Player struct {
	ID int32 `json:"ID"`
	color string `json:"Color"`
}

type Server struct {
	conns map[*websocket.Conn]bool
	player *Player
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}




func sendMoveMentCordinatesByClientRequest(clientRequest []byte,ws *websocket.Conn){
ws.Write(clientRequest)
}

type Direction struct {
	Type string `json:"type"`
	Direction string  `json:"direction"`
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buff := make([]byte, 1024)
	for {
		n, err := ws.Read(buff)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error in reading", err)
			continue
		}
		msg := buff[:n]
		fmt.Println(n)

		var direction Direction
		jsonData := []byte(msg)
		err = json.Unmarshal(jsonData,&direction)

		if err != nil {
		fmt.Println(err)

// sendMoveMentCordinatesByClientRequest([]byte(string(err)),ws)
		}
sendMoveMentCordinatesByClientRequest([]byte(direction.Direction),ws)
	}
}

func (s *Server) handleWs(ws *websocket.Conn) {
	fmt.Println("new connection requested", ws.RemoteAddr())
	ws.Write([]byte(("connection established !!")))
	s.conns[ws] = true
	s.readLoop(ws)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWs))
	http.ListenAndServe(":3000", nil)
	fmt.Println("server started on port 3000")

}
