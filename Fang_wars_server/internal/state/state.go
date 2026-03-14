package state

import "fang_wars_server/internal/types"



var GameState = types.SnakeState{
	HeadDirection: "L",
	BodyArray: []types.PixelCoordinates{
		{Column: 7, Row: 7},
		{Column: 7, Row: 8},
		{Column: 7, Row: 9},
	},
}

func HandleChangeDirection(direction string) {
	if GameState.HeadDirection == "U" && direction == "D" {
		return
	} else if GameState.HeadDirection == "D" && direction == "U" {
		return
	} else if GameState.HeadDirection == "L" && direction == "R" {
		return
	} else if GameState.HeadDirection == "R" && direction == "L" {
		return
	}

	GameState.HeadDirection = direction
}