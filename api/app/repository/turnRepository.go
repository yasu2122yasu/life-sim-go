package repository

import (
	"app/database"
	"app/model"
	"fmt"

	"gorm.io/gorm"
)

func GetCharacterTurn(id int) *gorm.DB {
	var character model.Character
	result := database.Db.First(&character, 1)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("err")
		} else {
			fmt.Println("err")
		}
		return nil
	}

	var turns []model.Turn
	response := database.Db.Where("character_id = ?", character.ID).Find(&turns)

	return response
}
