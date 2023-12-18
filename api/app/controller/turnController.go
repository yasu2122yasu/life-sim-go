package controller

import (
	"app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCharacterWithTurn(c *gin.Context) {
	var testCase int = 1
	responseData := repository.GetCharacterTurn(testCase)

	// エラー処理は後から追加する。
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	c.JSON(http.StatusOK, responseData)
}

// func CreateTurn(ctx *gin.Context) {
// 	var turn repository.Turn
// 	if err := ctx.ShouldBindJSON(&turn); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	err := c.turnRepository.CreateTurn(&turn)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, turn)
// }

// func UpdateTurn(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	var turn repository.Turn
// 	if err := ctx.ShouldBindJSON(&turn); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	err := c.turnRepository.UpdateTurn(id, &turn)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, turn)
// }

// func DeleteTurn(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	err := c.turnRepository.DeleteTurn(id)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{"message": "Turn deleted successfully"})
// }
