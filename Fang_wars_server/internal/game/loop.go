package game

import (
	"fang_wars_server/internal/network"
	"fang_wars_server/internal/state"
	"fang_wars_server/internal/types"
	"fmt"
	"time"
)

// Initialize the item coordinate
var itemCoordinate = types.PixelCoordinates{Column: 1, Row: 1}

// Function to move the snake forward based on its head direction
func moveSnakeForward() {
	head := state.GameState.BodyArray[0]
	newHead := types.PixelCoordinates{Column: head.Column, Row: head.Row}

	// Move the head    based on the current direction
	switch state.GameState.HeadDirection {
	case "L":
		newHead.Row--

	case "R":
		newHead.Row++
	case "D":
		newHead.Column++
	case "U":
		newHead.Column--
	}

	// Prepend the new head and remove the tail
	state.GameState.BodyArray = append([]types.PixelCoordinates{newHead}, state.GameState.BodyArray[:len(state.GameState .BodyArray)-1]...)
}

func GameLoop() {
	count := 0
	for {
		// Move the snake forward
		moveSnakeForward()

		// Broadcast the snake's coordinates
		msg := types.SnakeState{
			BodyArray: state.GameState .BodyArray,
			HeadDirection:state.GameState.HeadDirection,
		}
		network.Broadcast(msg)

		fmt.Println(state.GameState )
 
		count++
		time.Sleep(time.Second/10)
	}
}
