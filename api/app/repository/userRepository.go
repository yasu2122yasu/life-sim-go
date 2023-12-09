package repository

import (
	"app/database"
	"app/model"
)

func GetUserRepository(id string) (model.User, error) {
	var user model.User
	result := database.Db.First(&user, id)
	return user, result.Error
}
