package router

import (
	"app/controller"
	"app/middleware"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.Use(middleware.Cors())

	r.GET("/users/:id", controller.GetUser)
	r.POST("/users", controller.CreateUser)
	r.GET("/userinfo", controller.GetAllUser)
	r.POST("/login", controller.Login)

	r.GET("/events", controller.GetAllEvent)

	r.GET("characters/:id", controller.GetCharacter)
	r.GET("character", controller.GetCharacterWithTurn)

	r.Run(":8080")
}
