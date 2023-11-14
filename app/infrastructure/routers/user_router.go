package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yasu2122yasu/life-sim-go/app/infrastructure/database"
	"github.com/yasu2122yasu/life-sim-go/app/infrastructure/user/repositories"
	"github.com/yasu2122yasu/life-sim-go/app/interfaces/user/controllers"
	"github.com/yasu2122yasu/life-sim-go/app/usecase/user/interactors"
)

// UserRouter ユーザに関してルーティングを実行します．
func UserRouter(router *gin.Engine, db *database.DB) {
	userRouter := router.Group("/users")
	{
		c := controllers.NewUserController(interactors.NewUserInteractor(repositories.NewUserRepository(db)))
		userRouter.GET("/:id", c.GetUser)
		userRouter.POST("/", c.CreateUser)
	}
}
