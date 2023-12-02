package controller

import (
	"app/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetUser(c *gin.Context) {
	userID := c.Param("id")

	var user model.User
	result := db.First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetAllUser(c *gin.Context) {
	var users []model.User
	result := db.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, users)
}
