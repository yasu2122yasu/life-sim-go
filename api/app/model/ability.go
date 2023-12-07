package model

type Ability struct {
	ID             uint      `gorm:"primaryKey"`
	Knowledge      uint8     `gorm:"column:knowledge"`
	Strength       uint8     `gorm:"column:strength"`
	Social         uint8     `gorm:"column:social"`
	Stress         uint8     `gorm:"column:stress"`
	Decision       uint8     `gorm:"column:decision"`
	Luck           uint8     `gorm:"column:luck"`
	CharaCharacter Character `gorm:"foreignKey:CharacterID"`
}
