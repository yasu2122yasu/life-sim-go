package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

type Character struct {
	gorm.Model
	Name        string
	Description string
	UserID      int
	User        User
}

type Ability struct {
	gorm.Model
	Knowledge   uint8
	Strength    uint8
	Social      uint8
	Stress      uint8
	Decision    uint8
	Luck        uint8
	CharacterID int
	Character   Character
}

type Event struct {
	gorm.Model
	Description  string
	EventDetails []EventDetail `gorm:"foreignKey:EventId"`
}

type EventDetail struct {
	gorm.Model
	EventId uint
	Event   Event
	Rate    uint8
}

type Turn struct {
	gorm.Model
	Week        string
	CharacterID uint `gorm:"column:CharacterId"`
}
