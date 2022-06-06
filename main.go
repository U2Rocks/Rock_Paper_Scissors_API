package main

// ***THINGS TO DO***
// NOTE: player objects in return json for games are duds/but routes still work
// NOTE: wins do not go up when a player wins

import (
	"RockPaper_Api/database"
	"RockPaper_Api/models"
	"RockPaper_Api/routes"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

// this function loads all the routes for the api
func initialize_routes(app *fiber.App) {
	// player routes
	app.Get("/rockpaper/player/all", routes.AllPlayers)                // get a slice of all players
	app.Get("/rockpaper/player/leaderboard", routes.PlayerLeaderboard) // get the name and number of wins for all players
	app.Post("/rockpaper/player/add", routes.AddPlayer)                // add a player to the database
	app.Post("/rockpaper/player/getplayer", routes.GetPlayer)          // get a player by their playername
	app.Delete("/rockpaper/player/delete", routes.DeletePlayer)        // remove a player using their name
	app.Put("/rockpaper/player/tagline", routes.UpdateTagline)         // update the tagline of a player

	// game routes
	app.Get("/rockpaper/games/all", routes.AllGames)               // get all games[TESTED]
	app.Get("/rockpaper/games/complete", routes.CompleteGames)     // get all complete games[TESTED]
	app.Get("/rockpaper/games/incomplete", routes.IncompleteGames) // get all incomplete games[TESTED]
	app.Post("/rockpaper/games/get", routes.GetGame)               // get a game by its name[TESTED]
	app.Post("/rockpaper/games/add", routes.NewGame)               // create a new game[TESTED]
	app.Post("/rockpaper/games/turn", routes.SubmitMove)           // submit a players move[TESTED]
	app.Put("/rockpaper/games/cancel", routes.CancelGame)          // cancel a game[TESTED]
	app.Delete("/rockpaper/games/delete", routes.DeleteGame)       // soft delete a game[TESTED]

}

// this function initializes the database connection
func initialize_database() {
	var err error
	database.DBC, err = gorm.Open(sqlite.Open("RockPaper.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	fmt.Println("Connection to database now open")
	database.DBC.AutoMigrate(&models.Game{}, &models.Player{})
	fmt.Println("Database Successfully Migrated")
}

func main() {
	// define new fiber object
	app := fiber.New()

	// initialize routes and database connection
	initialize_database()
	initialize_routes(app)

	// create static route for admin panel
	app.Static("/", "./public", fiber.Static{Index: "index.html"})

	// wrap listen in log.fatal
	log.Fatal(app.Listen(":3000"))
}
