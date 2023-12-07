package model

type Event struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `gorm:"column:description"`
}
