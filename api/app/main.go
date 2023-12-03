package main

import (
	"app/controller"
	"app/database"
	"app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Db *gorm.DB

func main() {
	database.ConnectDatabase()
	// Ginのルーターの初期化
	r := gin.Default()
	r.Use(middleware.Cors())

	r.GET("/users/:id", controller.GetUser)
	r.POST("/users", controller.CreateUser)
	r.GET("/users", controller.GetAllUser)
	r.POST("/login", controller.Login)

	r.Run(":8080")
}
