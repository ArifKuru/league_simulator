package db

import (
	"log"

	"league-simulator/config"
	"league-simulator/models"
)

func SeedTeams() {
	db := config.DB

	// Check if exist
	var count int64
	db.Model(&models.Team{}).Count(&count)
	if count > 0 {
		log.Println("Already Exists, Skipping seed.")
		return
	}

	teams := []models.Team{
		{Name: "Liverpool", Strength: 90, Morale: 75},
		{Name: "Arsenal", Strength: 85, Morale: 75},
		{Name: "Manchester City", Strength: 80, Morale: 75},
		{Name: "Chelsea", Strength: 75, Morale: 75},
	}

	for _, team := range teams {
		if err := db.Create(&team).Error; err != nil {
			log.Println("Couldnt added:", team.Name, err)
		}
	}

	log.Println("Team succesfully added.")
}
