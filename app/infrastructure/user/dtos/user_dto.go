package dtos

import (
	"github.com/yasu2122yasu/life-sim-go/app/domain/user/entities"
	"github.com/yasu2122yasu/life-sim-go/app/domain/user/ids"
	"github.com/yasu2122yasu/life-sim-go/app/domain/user/values"
	"github.com/yasu2122yasu/life-sim-go/app/infrastructure"
)

// UserDTO NOTE: 利便性のため，DTOはパブリックフィールドとします．
type UserDTO struct {
	UserId            int `gorm:"primaryKey"`
	UserLastName      string
	UserFirstName     string
	UserLastKanaName  string
	UserFirstKanaName string
	UserGenderType    int
}

type UsersDTO []*UserDTO

var _ infrastructure.DTO = &UserDTO{}

// TableName テーブル名を定義します．
func (ud UserDTO) TableName() string {
	return "users"
}

// ToUser DTOをユーザエンティティに変換します．
func (ud UserDTO) ToUser() *entities.User {

	user := entities.NewUser(
		ids.UserId(ud.UserId),
		values.NewUserName(ud.UserLastName, ud.UserFirstName, ud.UserLastKanaName, ud.UserFirstKanaName),
		values.UserGenderType(ud.UserGenderType),
	)

	return user
}

// ToUsers UsersDTOをUsersに変換します．
func (usd UsersDTO) ToUsers() entities.Users {
	users := entities.Users{}

	for i, ud := range usd {
		users[i] = ud.ToUser()
	}

	return users
}
