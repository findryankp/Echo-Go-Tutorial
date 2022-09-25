package models

type Team struct {
	ID       int    `json:"id"`
	TeamName string `json:"team_name"`
}

type TransactionPlayer struct {
	ID       int `json:"id"`
	IdTeam   int `json:"id_team"`
	IdPlayer int `json:"id_player"`
	Player
}
