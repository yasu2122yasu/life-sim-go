package main

import (
	"app/database"
	"app/router"

	"gorm.io/gorm"
)

var Db *gorm.DB

func main() {
	database.ConnectDatabase()
	// Ginのルーターの初期化
	router.Router()
}
