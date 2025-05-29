package controllers

import (
	"league-simulator/config"
	"league-simulator/models"

	"github.com/gofiber/fiber/v2"
)

// List teams
func GetTeams(c *fiber.Ctx) error {
	var teams []models.Team
	result := config.DB.Find(&teams)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    teams,
	})
}
