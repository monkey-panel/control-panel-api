package api

import (
	"errors"

	"github.com/a3510377/control-panel-api/common"
	"github.com/a3510377/control-panel-api/common/codes"
	"github.com/a3510377/control-panel-api/common/database"

	"github.com/gin-gonic/gin"
)

func init() {
	router = append(router, registerAuthRouter, registerUsersRouter)
}

func registerAuthRouter(container common.Container, app *gin.RouterGroup) {
	db, authRouter := container.DB, app.Group("/auth")

	authRouter.POST("/login", func(c *gin.Context) {
	})

	authRouter.POST("/register", func(c *gin.Context) {
		var newUSer database.NewUser

		if err := c.Bind(&newUSer); err != nil {
			c.JSON(400, codes.Response[error](
				codes.InvalidFormBody,
				nil,
				common.TranslateError(err),
			))
			return
		}

		user, err := db.CreateUser(newUSer)
		if errors.Is(err, codes.UsernameAlreadyExists) {
			c.JSON(400, codes.Response[error](codes.UsernameAlreadyExists, nil, nil))
			return
		}

		c.JSON(200, codes.Response(codes.OK, user, nil))
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
