package controllers

import (
	"league-simulator/services"

	"github.com/gofiber/fiber/v2"
)

func SimulateWeek(c *fiber.Ctx) error {
	message, err := services.SimulateCurrentWeek()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	return c.JSON(fiber.Map{"success": true, "message": message})
}

func SimulateAll(c *fiber.Ctx) error {
	message, err := services.SimulateAllRemainingWeeks()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "message": message})
}

func ResetLeague(c *fiber.Ctx) error {
	if err := services.ResetMatches(); err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "message": "Maçlar sıfırlandı, takımlar korundu."})
}
