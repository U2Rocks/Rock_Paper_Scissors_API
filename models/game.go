package models

import (
	"gorm.io/gorm"
)

// this struct represents a game of rock_paper_scissors
type Game struct {
	gorm.Model
	GameName     string `json:"gamename"`
	PlOne        Player `gorm:"foreignKey:PlOneID"`
	PlOneID      int
	PlTwo        Player `gorm:"foreignKey:PlTwoID"`
	PlTwoID      int
	Winner_Found bool   `json:"winner_found"`
	Status       string `json:"status"`
	PlOneInput   string `json:"pl_one_input"`
	PlTwoInput   string `json:"pl_two_input"`
}

// this struct provides information to start a game
type NewGame struct {
	GameName  string `json:"gamename"`
	PlOneName string `json:"plonename"`
	PlTwoName string `json:"pltwoname"`
}

// this struct provides info on what each player has submitted
type GameInfo struct {
	PlOneName  string `json:"plonename"`
	PlTwoName  string `json:"pltwoname"`
	PlOneInput string `json:"pl_one_input"`
	PlTwoInput string `json:"pl_two_input"`
}

// this struct is used to submit a turn for players
type GameInput struct {
	GameName    string `json:"gamename"`
	PlayerName  string `json:"playername"`
	PlayerInput string `json:"playerinput"`
}

// this struct is used to search for a game by ID
type GameById struct {
	GameID int `json:"gameid"`
}

// this struct is used to search for a game by its gamename
type GameByName struct {
	GameName string `json:"gamename"`
}

// this struct is used to search for a game by its two players
type GameByPlayers struct {
	PlOne string `json:"plone"`
	PlTwo string `json:"pltwo"`
}

// this struct is used to search for a game by a single player(grabs first...)
type GameByPlayer struct {
	PlName string `json:"plname"`
}

// this function sets an input value for player one
func (G *Game) SetPlayerOneInput(newInput string) {
	G.PlOneInput = newInput
}

// this function sets an input value for player two
func (G *Game) SetPlayerTwoInput(newInput string) {
	G.PlOneInput = newInput
}

func (G *Game) SetWinnerFound(newVal bool) {
	G.Winner_Found = newVal
}

func (G *Game) SetGameStatus(newStatus string) {
	G.Status = newStatus
}
