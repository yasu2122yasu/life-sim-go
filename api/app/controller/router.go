package controller

import (
	"app/database"
	"app/middleware"

	"github.com/gin-gonic/gin"
)

func Router() {
	database.ConnectDatabase()
	r := gin.Default()
	r.Use(middleware.Cors())

	r.GET("/users/:id", GetUser)
	r.POST("/users", CreateUser)
	r.GET("/users", GetAllUser)
	r.POST("/login", Login)

	r.GET("characters/:id", GetCharacter)
	// repository内で指定されているid=1の情報を取得する。
	r.GET("character", GetCharacterWithTurn)

	r.Run(":8080")
}