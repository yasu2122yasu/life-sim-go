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
	r.POST("/create-user", controller.CreateUser)
	r.GET("/users", controller.GetAllUser)
	r.POST("/login", controller.Login)
	r.GET("/verify", controller.GetOnlyVerifyUser)

	r.GET("/events", controller.GetAllEvent)

	r.Run(":8080")
}
