package repository

import (
	"app/database"
	"app/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserRepository(c *gin.Context) (*model.User, error) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, err
	}

	var user model.User
	result := database.Db.First(&user, userID)
	if result.Error != nil {
		// データベースエラーの処理
		return nil, result.Error
	}

	return &user, nil
}
