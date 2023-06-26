package api

import (
	"github.com/a3510377/control-panel-api/common"
	"github.com/gin-gonic/gin"
)

func init() {
	router = append(router, registerAuthRouter, registerUsersRouter)
}

func registerAuthRouter(container common.Container, app *gin.RouterGroup) {
	authRouter := app.Group("/auth")

	authRouter.POST("/login", func(c *gin.Context) {
	})

	authRouter.POST("/register", func(c *gin.Context) {
	})
}

func registerUsersRouter(container common.Container, app *gin.RouterGroup) {
	usersRouter := app.Group("/users")

	usersRouter.GET("/@me")
	usersRouter.PATCH("/@me")
	usersRouter.GET("/@me/instances")
	usersRouter.GET("/@me/instances/:id/members")
	usersRouter.GET("/@me/instances/:id/members/:id")

	usersRouter.GET("/:id")
	usersRouter.PATCH("/:id")
	usersRouter.GET("/:id/instances")
	usersRouter.GET("/:id/instances/:id/members")
	usersRouter.GET("/:id/instances/:id/members/:id")
}

// username:
// password:
