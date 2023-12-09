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

	passwordHash := hashPassword(users.Password)
	users.Password = passwordHash

	if err := db.Create(&users).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	return nil
}
