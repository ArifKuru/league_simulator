package controllers

import (
	"league-simulator/config"
	"league-simulator/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MatchResponse struct {
	ID           uint   `json:"id"`
	Week         int    `json:"week"`
	HomeTeamName string `json:"home_team"`
	AwayTeamName string `json:"away_team"`
	HomeScore    int    `json:"home_score"`
	AwayScore    int    `json:"away_score"`
}

func GetMatches(c *fiber.Ctx) error {
	db := config.DB
	var matches []models.Match

	weekParam := c.Query("week")
	query := db.Preload("HomeTeam").Preload("AwayTeam")

	if weekParam != "" {
		week, err := strconv.Atoi(weekParam)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"error":   "Geçerli bir hafta numarası giriniz.",
			})
		}
		query = query.Where("week = ?", week)
	}

	if err := query.Find(&matches).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	var response []MatchResponse
	for _, m := range matches {
		response = append(response, MatchResponse{
			ID:           m.ID,
			Week:         m.Week,
			HomeTeamName: m.HomeTeam.Name,
			AwayTeamName: m.AwayTeam.Name,
			HomeScore:    m.HomeScore,
			AwayScore:    m.AwayScore,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    response,
	})
}

type EditMatchRequest struct {
	HomeScore int `json:"home_score"`
	AwayScore int `json:"away_score"`
}

type EditMatchResponse struct {
	ID        uint `json:"id"`
	Week      int  `json:"week"`
	HomeScore int  `json:"home_score"`
	AwayScore int  `json:"away_score"`
}

func EditMatch(c *fiber.Ctx) error {
	id := c.Params("id")

	var match models.Match
	if err := config.DB.First(&match, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"error":   "Maç bulunamadı.",
		})
	}

	var req EditMatchRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "Geçerli skorlar girilmedi.",
		})
	}

	match.HomeScore = req.HomeScore
	match.AwayScore = req.AwayScore

	if err := config.DB.Save(&match).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Skor güncellenemedi.",
		})
	}

	response := EditMatchResponse{
		ID:        match.ID,
		Week:      match.Week,
		HomeScore: match.HomeScore,
		AwayScore: match.AwayScore,
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Maç başarıyla güncellendi.",
		"data":    response,
	})
}

func GetLastPlayedWeekMatches(c *fiber.Ctx) error {
	db := config.DB
	var latestWeek int
	err := db.Model(&models.Match{}).
		Select("MAX(week)").
		Scan(&latestWeek).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success":     false,
			"latest_week": 0,
			"data":        []any{},
			"error":       "Hafta bilgisi alınamadı.",
		})
	}

	if latestWeek == 0 {
		return c.JSON(fiber.Map{
			"success":     true,
			"latest_week": 0,
			"data":        []any{},
		})
	}

	// O haftanın maçlarını çek
	var matches []models.Match
	err = db.Preload("HomeTeam").Preload("AwayTeam").
		Where("week = ?", latestWeek).
		Find(&matches).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success":     false,
			"latest_week": 0,
			"data":        []any{},
			"error":       err.Error(),
		})
	}

	var response []MatchResponse
	for _, m := range matches {
		response = append(response, MatchResponse{
			ID:           m.ID,
			Week:         m.Week,
			HomeTeamName: m.HomeTeam.Name,
			AwayTeamName: m.AwayTeam.Name,
			HomeScore:    m.HomeScore,
			AwayScore:    m.AwayScore,
		})
	}

	return c.JSON(fiber.Map{
		"success":     true,
		"latest_week": latestWeek,
		"data":        response,
	})
}
