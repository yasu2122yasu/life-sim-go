package model

type Character struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"column:user_id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	User        User   `gorm:"foreignKey:UserID"`
}
