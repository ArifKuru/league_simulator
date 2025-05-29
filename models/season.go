package models

import "gorm.io/gorm"

type SeasonState struct {
	gorm.Model
	CurrentWeek int
}
