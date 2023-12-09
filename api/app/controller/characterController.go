package controller

import (
	"app/database"
	"app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCharacter(c *gin.Context) {
	characterID := c.Param("id")

	var character model.Character
	result := database.Db.First(&character, characterID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}

	c.JSON(http.StatusOK, character)
}
