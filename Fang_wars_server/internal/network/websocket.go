package network

import (
	"encoding/json"
	"fang_wars_server/internal/state"
	"fang_wars_server/internal/types"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

type ResponseMessage struct {
    BodyArray string `json:"bodyArray"`
    HeadDirection    string `json:"headDirection"`
}

var clients = make(map[string]*websocket.Conn)
var clientsMutex sync.Mutex

func ConnRequestHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Upgrade error:", err)
        return
    }
    defer conn.Close()

    clientsMutex.Lock()
    clients[conn.RemoteAddr().String()] = conn
    clientsMutex.Unlock()

    for {
        _, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("Read error:", err)
            break
        }
        log.Printf("Received: %s", message)

        var clientCommand types.ClientCommand

err = json.Unmarshal([]byte(message),&clientCommand);
        if err != nil {
            log.Println("converting to json:", err)
            log.Println("converting to json message:",message)
            break
        }
state.HandleChangeDirection(clientCommand.Direction)
    }

    clientsMutex.Lock()
    delete(clients, conn.RemoteAddr().String())
    clientsMutex.Unlock()
}

func Broadcast(message types.SnakeState ) {
    clientsMutex.Lock()
    defer clientsMutex.Unlock()

    jsonMessage, err := json.Marshal(message)
    if err != nil {
        log.Println("Failed to marshal message:", err)
        return
    }

    for id, conn := range clients {
        err := conn.WriteMessage(websocket.TextMessage, jsonMessage)
        if err != nil {
            log.Printf("Failed to send message to %s: %v", id, err)
            conn.Close()
            delete(clients, id)
        }
    }
}

func main() {
    http.HandleFunc("/ws", ConnRequestHandler)
    log.Println("Server started on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("Server error:", err)
    }
}
