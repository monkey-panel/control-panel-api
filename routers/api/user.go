package api

import (
	"fmt"

	"github.com/a3510377/control-panel-api/common"
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
			c.JSON(400, gin.H{
				"code":    400,
				"message": "Bad Request",
			})
			return
		}

		data, err := db.CreateUser(newUSer)
		fmt.Println(data, err)
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
