package model

type AbilityCharacter struct {
	AbilityID   uint      `gorm:"primaryKey;column:ability_id"`
	CharacterID uint      `gorm:"primaryKey;column:character_id"`
	Ability     Ability   `gorm:"foreignKey:AbilityID"`
	Character   Character `gorm:"foreignKey:CharacterID"`
}
