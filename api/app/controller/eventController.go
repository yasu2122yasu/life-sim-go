package controller

import (
	"app/database"
	"app/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllEvent(c *gin.Context) {
	var events []model.Event
	result := database.Db.Find(&events)
	fmt.Println(result)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, events)
}
