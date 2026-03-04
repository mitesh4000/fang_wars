package game

import (
	"fang_wars_server/internal/game"
	"fang_wars_server/internal/network"
	"fang_wars_server/internal/types"
	"fmt"
	"time"
)

// Initialize the item coordinate
var itemCoordinate = types.PixelCoordinates{Column: 1, Row: 1}

// Function to move the snake forward based on its head direction
func moveSnakeForward() {
	head := GameState.BodyArray[0]
	newHead := types.PixelCoordinates{Column: head.Column, Row: head.Row}

	// Move the head based on the current direction
	switch GameState.HeadDirection {
	case "U":
		newHead.Row--
	case "D":
		newHead.Row++
	case "L":
		newHead.Column--
	case "R":
		newHead.Column++
	}

	// Prepend the new head and remove the tail
	GameState .BodyArray = append([]types.PixelCoordinates{newHead}, GameState .BodyArray[:len(GameState .BodyArray)-1]...)
}

func GameLoop() {
	count := 0
	for {
		// Move the snake forward
		moveSnakeForward()

		// Broadcast the snake's coordinates
		msg := types.SnakeState{
			BodyArray: GameState .BodyArray,
			HeadDirection:"D",
		}
		network.Broadcast(msg)

		fmt.Println("game" + game.GameState)
 
		count++
		time.Sleep(time.Second*3)
	}
}
