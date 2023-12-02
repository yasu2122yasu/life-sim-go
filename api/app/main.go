package main

import (
	"app/controller"
	"app/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// GORMの初期化
	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("Failed to migrate database")
	}

	// Ginのルーターの初期化
	r := gin.Default()
	r.Use(middleware.Cors())

	r.GET("/users/:id", controller.GetUser)
	r.POST("/users", controller.CreateUser)
	r.GET("/users", controller.GetAllUser)

	r.Run(":8080")
}
