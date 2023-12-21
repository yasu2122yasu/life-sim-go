package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string
	Email      string `gorm:"unique"`
	Password   string
	Characters []Character
}

type Character struct {
	gorm.Model
	Name        string
	Description string
	UserID      uint
	User        User
	Abilities   []Ability `gorm:"foreignKey:CharacterID"`
	Turns       []Turn    `gorm:"foreignKey:CharacterID"`
}

type Ability struct {
	gorm.Model
	Knowledge   uint8
	Strength    uint8
	Social      uint8
	Stress      uint8
	Decision    uint8
	Luck        uint8
	CharacterID uint `gorm:"column:character_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Character   Character
}

type Event struct {
	gorm.Model
	Description  string
	EventDetails []EventDetail `gorm:"foreignKey:EventID"`
}

type EventDetail struct {
	gorm.Model
	EventID uint `gorm:"column:event_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Event   Event
	Rate    uint8
}

type Turn struct {
	gorm.Model
	Week        string
	CharacterID uint `gorm:"column:character_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Character   Character
}
