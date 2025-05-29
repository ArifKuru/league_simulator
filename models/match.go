package models

import "gorm.io/gorm"

type Match struct {
	gorm.Model
	HomeTeamID uint
	AwayTeamID uint
	HomeScore  int
	AwayScore  int
	Week       int
	HomeTeam   Team `gorm:"foreignKey:HomeTeamID"`
	AwayTeam   Team `gorm:"foreignKey:AwayTeamID"`
}
