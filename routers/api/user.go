package api

import (
	"errors"

	"github.com/a3510377/control-panel-api/common"
	"github.com/a3510377/control-panel-api/common/codes"
	"github.com/a3510377/control-panel-api/common/database"
	"gorm.io/gorm/clause"

	"github.com/gin-gonic/gin"
)

func init() {
	router = append(router, registerAuthRouter, registerUsersRouter)
}

func registerAuthRouter(container common.Container, app *gin.RouterGroup) {
	db, authRouter := container.DB, app.Group("/auth")

	authRouter.POST("/login", func(c *gin.Context) {
		var loginUser database.LoginUser

		if err := c.Bind(&loginUser); err != nil {
			c.JSON(codes.Response[error](
				codes.InvalidFormBody,
				nil,
				common.TranslateError("zh_tw", err),
			))
			return
		}

		user := db.GetUserFromName(loginUser.Username)
		if user == nil {
			c.JSON(codes.Response[error](codes.UnknownUser, nil, nil))
		}

		user.AttachToken()
		c.JSON(codes.Response(codes.OK, user, nil))
	})

	authRouter.POST("/register", func(c *gin.Context) {
		var newUSer database.NewUser

		if err := c.Bind(&newUSer); err != nil {
			c.JSON(codes.Response[error](
				codes.InvalidFormBody,
				nil,
				common.TranslateError("zh_tw", err),
			))
			return
		}

		user, err := db.CreateUser(newUSer)
		if errors.Is(err, codes.UsernameAlreadyExists) {
			c.JSON(codes.Response[error](codes.UsernameAlreadyExists, nil, nil))
			return
		}

		user.AttachToken()
		c.JSON(codes.Response(codes.OK, user, nil))
	})
}

func registerUsersRouter(container common.Container, app *gin.RouterGroup) {
	usersRouter := app.Group("/users", AuthorizationMiddleware)
	usersRouterMe := usersRouter.Group("/@me")
	usersRouterOther := usersRouter.Group("/:id")

	usersRouterMe.GET("/", func(c *gin.Context) {
		c.JSON(codes.Response(codes.OK, GetUserFromContext(c), nil))
	})
	usersRouterMe.PATCH("/", func(c *gin.Context) {
		user := map[string]any{}
		if err := c.ShouldBindJSON(&user); err != nil || len(user) == 0 {
			c.JSON(codes.Response[error](
				codes.InvalidFormBody,
				nil,
				common.TranslateError("zh_tw", err),
			))
			return
		}

		db := database.GetDBFromContext(c)
		currentUser := database.DBUser{ID: GetUserFromContext(c).ID}
		db.Model(&currentUser).Clauses(clause.Returning{}).Omit("permissions").Updates(user)

		c.JSON(codes.Response(codes.OK, currentUser, nil))
	})
	usersRouterMe.GET("/instances", func(c *gin.Context) {
		c.JSON(codes.Response(codes.OK, GetUserFromContext(c), nil))
	})
	usersRouterMe.GET("/instances/:id/members")

	usersRouterOther.GET("/:id")
	usersRouterOther.PATCH("/:id")
	usersRouterOther.GET("/:id/instances")
	usersRouterOther.GET("/:id/instances/:id/members")
}
