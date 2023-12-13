package controller

import (
	"app/database"
	"app/model"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte("secret") // 本番環境ではもっと安全な方法でキーを管理する

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
	fmt.Println(122)

	fmt.Println(creds.Email)

	fmt.Println("バックエンドを通過しているのか確認")

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

	// JWTトークンの生成
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	// トークンを返す
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
