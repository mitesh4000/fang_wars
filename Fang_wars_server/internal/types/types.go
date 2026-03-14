package types

// Define a struct for pixel coordinates
type PixelCoordinates struct {
	Column int `json:"column"`
	Row    int `json:"row"`
}

// Define a struct for the snake state
type SnakeState struct {
	HeadDirection string            `json:"headDirection"`
	BodyArray     []PixelCoordinates `json:"bodyArray"`
}

type ClientCommand struct {
	Type string `json:"type"`
	Direction string `json:"direction"`
}