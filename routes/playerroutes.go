package routes

import (
	"RockPaper_Api/database"
	"RockPaper_Api/helperfunctions"
	"RockPaper_Api/models"

	"github.com/gofiber/fiber/v2"
)

// this function creates a new player in the database[name: string/tagline: string]
func AddPlayer(c *fiber.Ctx) error {
	newPlayer := new(models.Player)
	db := database.DBC

	if err := c.BodyParser(newPlayer); err != nil {
		return helperfunctions.ReturnJSONError(c, 400, "Could not parse new user body")
	}

	// set number of wins for new user to zero
	newPlayer.Num_Of_Wins = 0

	// check if bodyparser has assigned proper values
	if newPlayer.Name == "" {
		return helperfunctions.ReturnJSONError(c, 400, "Incoming entry missing name")
	}
	if newPlayer.Tagline == "" {
		return helperfunctions.ReturnJSONError(c, 400, "Incoming entry missing tagline")
	}
	if newPlayer.Num_Of_Wins != 0 {
		return helperfunctions.ReturnJSONError(c, 400, "Incoming entry missing number of wins")
	}

	// create a new entry in the database
	db.Create(&newPlayer)
	// create message to return to user
	finalMessage := "Successfully created the player: " + newPlayer.Name

	return helperfunctions.ReturnJSONResponse(c, 200, finalMessage)
}

// this function soft deletes a player from the database[playername: string]
func DeletePlayer(c *fiber.Ctx) error {
	// create models and db variables
	playerDeleteInfo := new(models.PlayerByName)
	var getPlayerInfo models.Player
	db := database.DBC
	// body parse info into model
	if err := c.BodyParser(playerDeleteInfo); err != nil {
		return helperfunctions.ReturnJSONError(c, 400, "Could not parse delete player body")
	}
	// query database for player info and scan into a variable
	db.Raw("SELECT * FROM players WHERE name = '" + playerDeleteInfo.PlayerName + "' AND deleted_at IS NULL").Scan(&getPlayerInfo)
	// check if valid model returned from query
	if getPlayerInfo.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 400, "Malformed Query(delete player)")
	}
	// delete the address of a model
	db.Delete(&getPlayerInfo)
	// return a helpful message to the user
	finalMessage := "The user " + playerDeleteInfo.PlayerName + " has been successfully deleted"
	return helperfunctions.ReturnJSONResponse(c, 200, finalMessage)
}

// this function lists all players in the database
func AllPlayers(c *fiber.Ctx) error {
	var PlayerList []models.Player
	db := database.DBC

	db.Find(&PlayerList)
	return c.Status(200).JSON(PlayerList)
}

// this function gets a players information[playername: string]
func GetPlayer(c *fiber.Ctx) error {
	playerSearch := new(models.PlayerByName)
	var foundPlayer models.Player
	db := database.DBC

	if err := c.BodyParser(playerSearch); err != nil {
		return helperfunctions.ReturnJSONError(c, 400, "Could not parse get player body")
	}

	db.Raw("SELECT * FROM players WHERE name = '" + playerSearch.PlayerName + "' AND deleted_at IS NULL").Scan(&foundPlayer)

	// error checking for bad name queries
	if foundPlayer.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 400, "Player Name Invalid/Malformed Query")
	}

	return c.Status(200).JSON(foundPlayer)
}

// this function lists a leaderboard of all players and their wins in descending order of wins
func PlayerLeaderboard(c *fiber.Ctx) error {
	var LeaderboardList []models.PlayerLeaderboardInfo
	db := database.DBC

	db.Raw("SELECT name, num_of_wins FROM players WHERE deleted_at IS NULL ORDER BY num_of_wins DESC").Scan(&LeaderboardList)
	return c.Status(200).JSON(LeaderboardList)
}

// this function updates a users tag line[playername: string/newtagline: string]
func UpdateTagline(c *fiber.Ctx) error {
	// initialize models and db variable
	playerNameInfo := new(models.UpdateTaglineInfo)
	var getPlayerInfo models.Player
	db := database.DBC
	// scan body for playername and updated tagline
	if err := c.BodyParser(playerNameInfo); err != nil {
		return helperfunctions.ReturnJSONError(c, 400, "Could not parse update tagline body")
	}

	// get player information from the database
	db.Raw("SELECT * FROM players WHERE name = '" + playerNameInfo.PlayerName + "' AND deleted_at IS NULL").Scan(&getPlayerInfo)
	// check if valid model returned from query
	if getPlayerInfo.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 400, "Malformed Query(update tag line)")
	}
	// update player info via variable
	getPlayerInfo.Tagline = playerNameInfo.NewTagline
	// save changes to database
	db.Save(&getPlayerInfo)
	// return a message to the user
	finalMessage := "You have successfully updated the records for user: " + playerNameInfo.PlayerName
	return helperfunctions.ReturnJSONResponse(c, 200, finalMessage)
}
