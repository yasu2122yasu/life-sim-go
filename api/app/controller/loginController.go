package controller

import (
	"app/database"
	"app/middleware"
	"app/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var creds LoginCredentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// データベースからユーザーを検索
	var user model.User
	result := database.Db.Where("email = ?", creds.Email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect email"})
		return
	}

	// パスワードが一致するか確認
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect email or password"})
		return
	}

	tokenString, err := middleware.GenerateToken()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect tokenString"})
		return
	}

	// トークンを返す
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
