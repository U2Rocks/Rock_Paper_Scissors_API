package helperfunctions

import (
	"RockPaper_Api/models"
)

func FillGame(g *models.Game) {
	g.Status = "Starting Duel"
	g.Winner_Found = false
	g.PlOneInput = "N/A"
	g.PlTwoInput = "N/A"
}

// function that checks if a game has tied
func CheckForTie(g *models.Game) bool {
	if g.PlOneInput == g.PlTwoInput {
		return true
	}
	return false
}

// function that checks if a player has won the game
func CheckForWin(g *models.Game) string {
	// if a player reaches a win condition return one or two...or else return "false"
	// paper beats rock
	if g.PlOneInput == "paper" && g.PlTwoInput == "rock" {
		return "one"
	}
	if g.PlOneInput == "rock" && g.PlTwoInput == "paper" {
		return "two"
	}
	// rock beats scissors
	if g.PlOneInput == "rock" && g.PlTwoInput == "scissors" {
		return "one"
	}
	if g.PlOneInput == "scissors" && g.PlTwoInput == "rock" {
		return "two"
	}
	// scissors beats paper
	if g.PlOneInput == "scissors" && g.PlTwoInput == "paper" {
		return "one"
	}
	if g.PlOneInput == "paper" && g.PlTwoInput == "scissors" {
		return "two"
	}
	return "false"
}

// this function finishes a game
func FinishGame(g *models.Game, winnernum string, winnername string) {
	if winnernum == "one" {
		g.Status = winnername + " has won the game"
	}
	if winnernum == "two" {
		g.Status = winnername + " has won the game"
	}
	g.Winner_Found = true
}
