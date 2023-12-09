package database

// import (
// 	"app/model"

// 	"gorm.io/gorm"
// )

// func AbilitiySeed(db *gorm.DB) {
// 	abilities := []model.Ability{
// 		{Knowledge: 100, Strength: 25, Social: 70, Stress: 30, Decision: 40, Luck: 60},
// 	}
// 	for _, ability := range abilities {
// 		db.Create(&ability)
// 	}

// 	// Characterの作成
// 	character := model.Character{
// 		Name:        "キャラクター名",
// 		Description: "キャラクターの説明",
// 		UserID:      user.ID,
// 		Abilities:   abilities,
// 	}
// 	db.Create(&character)
// }
