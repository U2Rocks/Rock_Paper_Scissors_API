package routes

import (
	"RockPaper_Api/database"
	"RockPaper_Api/helperfunctions"
	"RockPaper_Api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// this function creates a new game[gamename: string/plonename: string/pltwoname: string]
func NewGame(c *fiber.Ctx) error {
	// make new variables
	var PlayerOne models.Player
	var PlayerTwo models.Player
	var MakeGame models.Game
	NewGameInfo := new(models.NewGame)
	db := database.DBC

	// parse json body into gameinfo model
	if err := c.BodyParser(NewGameInfo); err != nil {
		return helperfunctions.ReturnJSONError(c, 400, "Could not parse new game body")
	}
	// query for first player ID
	db.Raw("SELECT * FROM players WHERE name = '" + NewGameInfo.PlOneName + "' AND deleted_at IS NULL").Scan(&PlayerOne)
	if PlayerOne.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 400, "Malformed Query(new game/first player query)")
	}
	// query for second player ID
	db.Raw("SELECT * FROM players WHERE name = '" + NewGameInfo.PlTwoName + "' AND deleted_at IS NULL").Scan(&PlayerTwo)
	if PlayerTwo.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 400, "Malformed Query(new game/second player query)")
	}
	// form completed Game object(fill in required fields)
	MakeGame.PlOneID = int(PlayerOne.ID)
	MakeGame.PlTwoID = int(PlayerTwo.ID)
	MakeGame.GameName = NewGameInfo.GameName
	helperfunctions.FillGame(&MakeGame)
	// check if game model was populated
	if MakeGame.GameName == "" {
		return helperfunctions.ReturnJSONError(c, 400, "New Game not populated correctly")
	}
	// add game to database
	db.Create(&MakeGame)
	// return message to user
	finalMessage := "The game: " + MakeGame.GameName + " has been created"
	return helperfunctions.ReturnJSONResponse(c, 200, finalMessage)
}

// this function soft deletes a game from the database[gamename: string]
func DeleteGame(c *fiber.Ctx) error {
	// declare variables
	var getGame models.Game
	getGameInfo := new(models.GameByName)
	db := database.DBC
	// parse body
	if err := c.BodyParser(getGameInfo); err != nil {
		return helperfunctions.ReturnJSONError(c, 400, "Could not parse delete game body")
	}
	// grab game info
	db.Raw("SELECT * FROM games WHERE game_name = '" + getGameInfo.GameName + "'  AND deleted_at IS NULL").Scan(&getGame)
	if getGame.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 200, "Malformed Query(deletegame/getting game information)")
	}
	// delete game
	db.Delete(&getGame)
	// return message to user
	finalMessage := "The game: " + getGameInfo.GameName + " has been successfully deleted"
	return helperfunctions.ReturnJSONResponse(c, 200, finalMessage)
}

// this function lists all games in the database
func AllGames(c *fiber.Ctx) error {
	var GameList []models.Game
	db := database.DBC

	db.Find(&GameList)
	return c.Status(200).JSON(GameList)
}

// this function lists all completed games in the database
func CompleteGames(c *fiber.Ctx) error {
	var completeList []models.Game
	db := database.DBC

	db.Raw("SELECT * FROM games WHERE winner_found = true AND deleted_at IS NULL").Scan(&completeList)
	return c.Status(200).JSON(completeList)
}

// this function lists all ongoing games in the database
func IncompleteGames(c *fiber.Ctx) error {
	var completeList []models.Game
	db := database.DBC

	db.Raw("SELECT * FROM games WHERE winner_found = false AND deleted_at IS NULL").Scan(&completeList)
	return c.Status(200).JSON(completeList)
}

// this function gets a games information[gamename: string]
func GetGame(c *fiber.Ctx) error {
	// declare variables
	var getGame models.Game
	getGameInfo := new(models.GameByName)
	db := database.DBC
	// parse body
	if err := c.BodyParser(getGameInfo); err != nil {
		return helperfunctions.ReturnJSONError(c, 400, "Could not parse get game body")
	}
	// get game info from database
	db.Raw("SELECT * FROM games WHERE game_name = '" + getGameInfo.GameName + "'  AND deleted_at IS NULL").Scan(&getGame)
	if getGame.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 200, "Malformed Query(getgame/getting game information)")
	}
	// return game info to the user
	return c.Status(200).JSON(getGame)
}

// this function submits a turn for a player[gamename: string/playername: string/playerinput: string]
func SubmitMove(c *fiber.Ctx) error {
	// declare variables
	var playerNum int
	var playerOne models.Player
	var playerTwo models.Player
	var getGame models.Game
	submitInfo := new(models.GameInput)
	db := database.DBC

	// parse body(player name/input/gamename)
	if err := c.BodyParser(submitInfo); err != nil {
		return helperfunctions.ReturnJSONError(c, 400, "Could not parse submit move body")
	}
	// get game info
	db.Raw("SELECT * FROM games WHERE game_name = '" + submitInfo.GameName + "'  AND deleted_at IS NULL").Scan(&getGame)
	if getGame.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 400, "Malformed Query(submitmove/getting game info)")
	}

	// check if game is already won
	if getGame.Winner_Found == true {
		return helperfunctions.ReturnJSONError(c, 400, getGame.Status)
	}

	// get info on both players in the game
	db.Raw("SELECT * FROM players WHERE id = " + strconv.FormatInt(int64(getGame.PlOneID), 10) + " AND deleted_at IS NULL").Scan(&playerOne)
	if playerOne.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 400, "Malformed Query(submitmove/getting player one info)")
	}
	db.Raw("SELECT * FROM players WHERE id = " + strconv.FormatInt(int64(getGame.PlTwoID), 10) + " AND deleted_at IS NULL").Scan(&playerTwo)
	if playerTwo.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 400, "Malformed Query(submitmove/getting player two info)")
	}

	// check if input player name is in the game
	if submitInfo.PlayerName != playerOne.Name && submitInfo.PlayerName != playerTwo.Name {
		return helperfunctions.ReturnJSONError(c, 400, "submitmove/no player with that name exists in this game")
	}

	// -- figure out what player is currently playing
	if submitInfo.PlayerName == playerOne.Name {
		playerNum = 0
	} else if submitInfo.PlayerName == playerTwo.Name {
		playerNum = 1
	} else {
		return helperfunctions.ReturnJSONError(c, 400, "submitmove/could not select a player properly")
	}
	// -- update input value to variable
	if playerNum == 0 {
		if getGame.PlOneInput == "rock" || getGame.PlOneInput == "paper" || getGame.PlOneInput == "scissors" {
			return helperfunctions.ReturnJSONError(c, 400, "You have already submitted a move player one")
		}
		getGame.PlOneInput = submitInfo.PlayerInput
		db.Save(&getGame)
	} else if playerNum == 1 {
		if getGame.PlTwoInput == "rock" || getGame.PlTwoInput == "paper" || getGame.PlTwoInput == "scissors" {
			return helperfunctions.ReturnJSONError(c, 400, "You have already submitted a move player two")
		}
		getGame.PlTwoInput = submitInfo.PlayerInput
		db.Save(&getGame)
	} else {
		return helperfunctions.ReturnJSONError(c, 400, "submitmove/could not update player input")
	}
	// run functions to check for wins or ties

	// -- check for tie first and reset inputs if that happens
	isTie := helperfunctions.CheckForTie(&getGame)
	// reset inputs and status if a tie has been found
	if isTie == true {
		getGame.PlOneInput = "N/A"
		getGame.PlTwoInput = "N/A"
		getGame.Status = "Players have tied a round"
		db.Save(&getGame)
	}
	// -- check for wins second and complete game values if that happens
	isWin := helperfunctions.CheckForWin(&getGame)
	if isWin != "false" {
		// finish the game depending on who won
		if isWin == "one" {
			helperfunctions.FinishGame(&getGame, "one", playerOne.Name)
			playerOne.Num_Of_Wins += 1
		} else if isWin == "two" {
			helperfunctions.FinishGame(&getGame, "two", playerTwo.Name)
			playerTwo.Num_Of_Wins += 1
		}
		// update the database
		db.Save(&getGame)
		// return a custom message to the user via JSON
		if isWin == "one" {
			winningMessage := playerOne.Name + " has won the game"
			return helperfunctions.ReturnJSONResponse(c, 200, winningMessage)
		}
		if isWin == "two" {
			winningMessage := playerTwo.Name + " has won the game"
			return helperfunctions.ReturnJSONResponse(c, 200, winningMessage)
		}
		return helperfunctions.ReturnJSONResponse(c, 200, "Generic Winning Message")
	}
	// if no win or tie return successful message to the user
	finalMessage := submitInfo.PlayerName + " has successfully submitted a move"
	return helperfunctions.ReturnJSONResponse(c, 200, finalMessage)
}

// this function cancels an ongoing game by marking it as cancelled[gamename: string]
func CancelGame(c *fiber.Ctx) error {
	// declare variables
	var getGame models.Game
	getGameInfo := new(models.GameByName)
	db := database.DBC
	// parse body
	if err := c.BodyParser(getGameInfo); err != nil {
		return helperfunctions.ReturnJSONError(c, 400, "Could not parse cancel game body")
	}
	// get game info via name
	db.Raw("SELECT * FROM games WHERE game_name = '" + getGameInfo.GameName + "'  AND deleted_at IS NULL").Scan(&getGame)
	if getGame.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 200, "Malformed Query(cancelgame/getting game information)")
	}
	// update input varibles/winner_found/status
	getGame.PlOneInput = "CANCELLED"
	getGame.PlTwoInput = "CANCELLED"
	getGame.Winner_Found = true
	getGame.Status = "Game Cancelled"
	db.Save(&getGame)
	// return message to user
	finalMessage := "The game: " + getGameInfo.GameName + " has been successfully cancelled"
	return helperfunctions.ReturnJSONResponse(c, 200, finalMessage)
}
