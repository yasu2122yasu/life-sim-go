package repositories

import (
	"github.com/yasu2122yasu/life-sim-go/app/domain/user/entities"
	"github.com/yasu2122yasu/life-sim-go/app/domain/user/ids"
)

type UserRepository interface {
	Create(*entities.User) error
	FindById(ids.UserId) (*entities.User, error)
	FindAll() (entities.Users, error)
	Update(*entities.User) error
	Delete(ids.UserId) error
}
