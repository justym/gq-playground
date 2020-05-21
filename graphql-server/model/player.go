package model

type Player struct {
	ID             int      `json:"int"`
	Name           string   `json:"name"`
	HighScore      int      `json:"highScore"`
	IsOnline       bool     `json:"isOnline"`
	Location       string   `json:"location"`
	LevelsUnlocked []string `json:"levelsUnlocked"`
}

var players = []Player{
	Player{
		ID: 123, 
		Name: "Gabi", 
		HighScore: 1100,
		IsOnline: true,
		Location: "Italy",
	},
	Player{
		ID: 230, 
		Name: "Justym", 
		HighScore: 21, 
		IsOnline: false,
		Location: "Japan",
	},
	Player{
		ID: 1000,
		Name: "John",
		HighScore: 2020,
		IsOnline: true,
		Location: "US",
	},
	Player{
		ID: 2010,
		Name: "Xiu",
		HighScore: 20000,
		IsOnline: false,
		Location: "China",
	},
}

func NewPlayers() []Player {
	return players
}