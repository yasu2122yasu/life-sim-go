package repository

import (
	"app/database"
	"app/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func GetCharacterTurns(id int) ([]model.Turn, error) {
	var character model.Character
	result := database.Db.First(&character, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Handle not found error
			fmt.Println("Character not found")
		} else {
			// Handle other errors
			fmt.Println("Database error:", result.Error)
		}
		return nil, result.Error
	}

	var turns []model.Turn
	response := database.Db.Where("character_id = ?", character.ID).Find(&turns)

	if response.Error != nil {
		return nil, response.Error
	}

	return turns, nil
}
