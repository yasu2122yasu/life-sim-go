package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yasu2122yasu/life-sim-go/app/infrastructure/database"
	"github.com/yasu2122yasu/life-sim-go/app/infrastructure/middlewares"
)

type Router struct {
	router *gin.Engine
	db     *database.DB
}

// NewRouter コンストラクタ
func NewRouter(router *gin.Engine, db *database.DB) *Router {
	return &Router{
		router: router,
		db:     db,
	}
}

// Run 全てのルーティングを実行します．
func (r *Router) Run() error {
	r.router.Use(
		middlewares.HandleError(),
	)

	HealthCheckRouter(r.router)

	UserRouter(r.router, r.db)

	return r.router.Run()
}
