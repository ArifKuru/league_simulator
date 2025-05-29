package controllers

import (
	"league-simulator/config"
	"league-simulator/models"
	"sort"

	"github.com/gofiber/fiber/v2"
)

func GetStandings(c *fiber.Ctx) error {
	db := config.DB

	var teams []models.Team
	var matches []models.Match

	if err := db.Find(&teams).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	if err := db.Where("is_played = ?", true).Find(&matches).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	standings := map[uint]*models.Standing{}

	for _, team := range teams {
		standings[team.ID] = &models.Standing{
			TeamID:   team.ID,
			TeamName: team.Name,
		}
	}

	// Procced matches to score
	for _, match := range matches {
		home := standings[match.HomeTeamID]
		away := standings[match.AwayTeamID]

		home.Played++
		away.Played++

		home.GoalsFor += match.HomeScore
		home.GoalsAgainst += match.AwayScore

		away.GoalsFor += match.AwayScore
		away.GoalsAgainst += match.HomeScore

		if match.HomeScore > match.AwayScore {
			home.Won++
			away.Lost++
			home.Points += 3
		} else if match.HomeScore < match.AwayScore {
			away.Won++
			home.Lost++
			away.Points += 3
		} else {
			home.Drawn++
			away.Drawn++
			home.Points++
			away.Points++
		}
	}

	var table []models.Standing
	for _, s := range standings {
		s.GoalDifference = s.GoalsFor - s.GoalsAgainst
		table = append(table, *s)
	}

	// Sort
	sort.Slice(table, func(i, j int) bool {
		if table[i].Points == table[j].Points {
			return table[i].GoalDifference > table[j].GoalDifference
		}
		return table[i].Points > table[j].Points
	})

	return c.JSON(fiber.Map{
		"success": true,
		"data":    table,
	})
}
