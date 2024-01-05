package database

import (
	"app/model"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"gorm.io/gorm"
)

// パスワードをハッシュ化する
func hashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

func UserSeed(db *gorm.DB) error {
	var count int64
	db.Model(&model.User{}).Count(&count)
	if count > 0 {
		return nil
	}
	users := model.User{Name: "abe", Email: "mailmail@gmail.com", Password: "password"}

	// seeding時には、パスワードのハッシュ化はおこなわない
	// passwordHash := hashPassword(users.Password)
	// users.Password = passwordHash

	if err := db.Create(&users).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	return nil
}

func EventSeed(db *gorm.DB) error {
	var count int64
	db.Model(&model.Event{}).Count(&count)
	if count > 0 {
		return nil
	}
	events := model.Event{Description: "event1"}

	// seeding時には、パスワードのハッシュ化はおこなわない
	// passwordHash := hashPassword(users.Password)
	// users.Password = passwordHash

	if err := db.Create(&events).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	return nil
}
