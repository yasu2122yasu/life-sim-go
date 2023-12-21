package controller

import (
	"app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// characterに紐づくturnを取得する
func GetCharacterWithTurn(c *gin.Context) {
	var characterId int = 1
	responseData, err := repository.GetCharacterTurns(characterId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, responseData)
}
