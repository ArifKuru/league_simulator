package models

import (
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Strength int
	Morale   int
}

// Power calc .Optionally some coefficients can be added for deacrease or increase strength or morale effect

func (t *Team) EffectivePower() float64 {
	return float64(t.Strength) * float64(t.Morale) / 100
}
