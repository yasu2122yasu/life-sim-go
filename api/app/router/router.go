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
	r.GET("/users", controller.GetAllUser)
	r.POST("/login", controller.Login)
	r.POST("/verify", controller.GetUserDetail)

	r.GET("/events", controller.GetAllEvent)

	r.Run(":8080")
}
