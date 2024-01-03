package main

import (
	"app/controller"
	"app/database"

	"gorm.io/gorm"
)

var Db *gorm.DB

func main() {
	database.ConnectDatabase()
	// Ginのルーターの初期化
	controller.Router()

}
