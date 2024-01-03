package controller

import (
	"app/middleware"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.Use(middleware.Cors())

	r.GET("/users/:id", GetUser)
	r.POST("/users", CreateUser)
	r.GET("/userinfo", GetAllUser)
	r.POST("/login", Login)

	r.GET("/events", GetAllEvent)

	r.GET("characters/:id", GetCharacter)
	r.GET("character", GetCharacterWithTurn)

	r.Run(":8080")
}
