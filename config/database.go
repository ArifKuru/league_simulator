package config

import (
	"league-simulator/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("league.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = database.AutoMigrate(
		&models.Team{},
		&models.Match{},
		&models.SeasonState{},
	)
	if err != nil {
		log.Fatal("Migration error:", err)
	}

	DB = database
}
