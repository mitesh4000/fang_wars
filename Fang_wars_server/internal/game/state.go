package game

import "fang_wars_server/internal/types"



var GameState = types.SnakeState{
	HeadDirection: "L",
	BodyArray: []types.PixelCoordinates{
		{Column: 7, Row: 7},
		{Column: 7, Row: 8},
		{Column: 7, Row: 9},
	},
}