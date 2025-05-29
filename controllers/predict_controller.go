package controllers

import (
	"league-simulator/config"
	"league-simulator/services"

	"github.com/gofiber/fiber/v2"
)

func PredictChampion(c *fiber.Ctx) error {
	db := config.DB
	var seasonState struct {
		CurrentWeek int
	}
	db.Table("season_states").Select("current_week").First(&seasonState)

	if seasonState.CurrentWeek < 5 {
		return c.JSON(fiber.Map{
			"success":         false,
			"prediction_week": 0,
			"error":           "At least 4 weeks must be played before making predictions.",
		})
	}

	var predictor services.Predictor

	switch c.Query("method") {
	case "montecarlo":
		predictor = services.MonteCarloPredictor{}
	default:
		predictor = services.DefaultPredictor{}
	}

	result, err := predictor.Predict(1000)
	predictionWeek := seasonState.CurrentWeek - 1

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success":         false,
			"prediction_week": predictionWeek,
			"error":           err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success":         true,
		"prediction_week": predictionWeek,
		"prediction":      result,
	})
}
