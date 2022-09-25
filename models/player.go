package models

type Player struct {
	ID         int    `json:"id"`
	PlayerName string `json:"player_name"`
	Position   string `json:"position"`
	Age        uint   `json:"age"`
}
