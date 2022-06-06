package models

import (
	"gorm.io/gorm"
)

// this struct is used to represent a rock_paper_scissor player
type Player struct {
	gorm.Model
	Name        string `json:"name"`
	Tagline     string `json:"tagline"`
	Num_Of_Wins int    `json:"num_of_wins"`
}

// this struct is used to search for a player by their ID
type PlayerByID struct {
	PlayerID int `json:"playerbyid"`
}

// this struct is used to search for a player by their name
type PlayerByName struct {
	PlayerName string `json:"playername"`
}

// this struct returns information on a player
type PlayerInfo struct {
	Name        string `json:"name"`
	Tagline     string `json:"tagline"`
	Num_Of_Wins int    `json:"num_of_wins"`
}

// this struct creates a new player
type NewPlayerInfo struct {
	Name    string `json:"name"`
	Tagline string `json:"tagline"`
}

type PlayerLeaderboardInfo struct {
	Name        string `json:"name"`
	Num_Of_Wins int    `json:"num_of_wins"`
}

type UpdateTaglineInfo struct {
	PlayerName string `json:"playername"`
	NewTagline string `json:"newtagline"`
}

func (P *Player) SetNewTagLine(newLine string) {
	P.Tagline = newLine
}

func (P *Player) SetWinNumber(newWinNum int) {
	P.Num_Of_Wins = newWinNum
}
