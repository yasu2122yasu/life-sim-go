package controller

import (
	"app/database"
	"app/middleware"
	"app/model"
	"app/repository"
	"app/request"
	"app/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(c *gin.Context) {
	user, err := repository.GetUserRepository(c)

	fmt.Println(err)

	c.JSON(http.StatusOK, user)
}

// 認証したユーザー自身の詳細を取得する。認証ユーザーのみアクセス可能
func GetUserDetail(c *gin.Context) {
	var clreq request.CheckLoginRequest
	var clres response.CheckLoginResponse

	if err := c.ShouldBindJSON(&clreq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwtToken, err := middleware.VerifyToken(clreq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JWT tokenが異なる"})
	}

	fmt.Println(jwtToken)

	clres.Status = true
	c.JSON(http.StatusOK, clres.Status)
}

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// データベースにユーザーを作成
	result := database.Db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetAllUser(c *gin.Context) {
	var users []model.User
	result := database.Db.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	fmt.Println(5739729)

	c.JSON(http.StatusOK, users)
}
