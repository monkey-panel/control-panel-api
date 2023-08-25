package api

import (
	"errors"

	"github.com/monkey-panel/control-panel-api/common"
	"github.com/monkey-panel/control-panel-api/common/codes"
	"github.com/monkey-panel/control-panel-api/common/database"
	"github.com/monkey-panel/control-panel-api/common/utils"
	"github.com/monkey-panel/control-panel-api/global"

	"github.com/gin-gonic/gin"
)

func init() {
	router = append(router, registerAuthRouter, registerUsersRouter)
}

func registerAuthRouter(app *gin.RouterGroup) {
	db, authRouter := global.DB, app.Group("/auth")

	authRouter.POST("/login", func(c *gin.Context) {
		currentUser := GetUserFromContext(c)
		var loginUser database.LoginUser

		if err := c.Bind(&loginUser); err != nil {
			c.JSON(codes.Response[error](
				codes.InvalidFormBody,
				nil,
				common.TranslateError(currentUser.Lang, err),
			))
			return
		}

		user := db.GetUserFromName(loginUser.Username)
		if user == nil {
			c.JSON(codes.Response[error](codes.UnknownUser, nil, nil))
			return
		}
		if utils.BcryptCheck(loginUser.Password, user.Password) {
			user_info := user.ToUserInfo()
			user_info.AttachToken()
			c.JSON(codes.Response(codes.OK, user_info, nil))
		} else {
			c.JSON(codes.Response[error](codes.InvalidPassword, nil, nil))
		}
	})

	authRouter.POST("/register", func(c *gin.Context) {
		currentUser := GetUserFromContext(c)
		var newUSer database.NewUser

		if err := c.Bind(&newUSer); err != nil {
			c.JSON(codes.Response[error](
				codes.InvalidFormBody,
				nil,
				common.TranslateError(currentUser.Lang, err),
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

func registerUsersRouter(app *gin.RouterGroup) {
	usersRouter := app.Group("/users", AuthorizationMiddleware)
	usersRouterMe := usersRouter.Group("/@me")
	usersRouterOther := usersRouter.Group("/:id")

	usersRouterMe.GET("/", func(c *gin.Context) {
		c.JSON(codes.Response(codes.OK, GetUserFromContext(c), nil))
	})
	usersRouterMe.PATCH("/", func(c *gin.Context) {
		currentUser := GetUserFromContext(c)
		userEdit := database.EditUser{}
		if err := c.ShouldBindJSON(&userEdit); err != nil {
			c.JSON(codes.Response[error](
				codes.InvalidFormBody,
				nil,
				common.TranslateError(currentUser.Lang, err),
			))
			return
		}

		db := database.GetDBFromContext(c)
		user := database.DBUser{ID: GetUserFromContext(c).ID}
		if d := db.Model(&user).Omit("permissions").Updates(userEdit); d.Error != nil {
			c.JSON(codes.Response[error](
				codes.UnknownUser,
				nil,
				common.TranslateError(currentUser.Lang, d.Error),
			))
			return
		}
		user = *db.GetUserFromID(user.ID)

		c.JSON(codes.Response(codes.OK, user.ToUserInfo(), nil))
	})
	usersRouterMe.GET("/instances", func(c *gin.Context) {
		c.JSON(codes.Response(codes.OK, GetUserFromContext(c), nil))
	})

	usersRouterOther.GET("/:id")
	usersRouterOther.PATCH("/:id")
}
