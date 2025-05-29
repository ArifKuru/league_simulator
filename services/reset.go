package services

import (
	"league-simulator/config"
	"league-simulator/models"
)

func ResetMatches() error {
	db := config.DB

	//Delete all matches
	if err := db.Exec("DELETE FROM matches").Error; err != nil {
		return err
	}

	//Reset Morale of all teams
	if err := db.Model(&models.Team{}).Where("1 = 1").Update("morale", 75).Error; err != nil {
		return err
	}

	//Reset Season
	if err := db.Exec("DELETE FROM season_states").Error; err != nil {
		return err
	}

	return nil
}
