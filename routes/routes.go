package routes

import (
	"league-simulator/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/teams", controllers.GetTeams)
	app.Post("/simulate", controllers.SimulateWeek) // yeni versiyon
	app.Post("/simulate/all", controllers.SimulateAll)
	app.Post("/reset", controllers.ResetLeague)
	app.Get("/matches", controllers.GetMatches)
	app.Get("/standings", controllers.GetStandings)
	app.Get("/predict", controllers.PredictChampion)
	app.Post("/matches/:id/edit", controllers.EditMatch)
	app.Get("/matches/last", controllers.GetLastPlayedWeekMatches)
}
