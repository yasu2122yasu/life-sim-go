package boundaries

import (
	"github.com/yasu2122yasu/life-sim-go/app/usecase/user/requests"
	"github.com/yasu2122yasu/life-sim-go/app/usecase/user/responses"
)

type UserInputBoundary interface {
	CreateUser(*requests.UserCreateRequest) (*responses.UserCreateResponse, error)
	GetUser(*requests.UserGetRequest) (*responses.UserGetResponse, error)
}
