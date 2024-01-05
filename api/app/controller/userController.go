package controller

import (
	"app/database"
	"app/middleware"
	"app/model"
	"app/repository"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUser(c *gin.Context) {
	user, err := repository.GetUserRepository(c)

	fmt.Println(err)

	c.JSON(http.StatusOK, user)
}

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

	tokenString, err := middleware.GenerateToken(user.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect tokenString"})
		return
	}

	// トークンを返す
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// 認証したユーザーのみがアクセスできる
func GetOnlyVerifyUser(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	// JWT tokenの頭についている"Bearer "を削除
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := middleware.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	var users []model.User
	result := database.Db.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if claims, ok := token.Claims.(*middleware.Claims); ok && token.Valid {
		c.JSON(http.StatusOK, gin.H{"email": claims.Email})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	}
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

	c.JSON(http.StatusOK, users)
}
