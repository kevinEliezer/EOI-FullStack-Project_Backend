package models

import "time"

// CharacterSection : Character section structure
type CharacterSection struct {
	ID        uint `gorm:"primary_key" json:"id"`
	UpdatedAt time.Time
	CreatedAt time.Time
}