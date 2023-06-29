package api

import (
	"errors"
	"strings"

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
		var loginUser database.LoginUser

		if err := c.Bind(&loginUser); err != nil {
			c.JSON(400, codes.Response[error](
				codes.InvalidFormBody,
				nil,
				common.TranslateError("zh_tw", err),
			))
			return
		}

		user := db.GetUserFromName(loginUser.Username)
		if user == nil {
			c.JSON(400, codes.Response[error](codes.UnknownUser, nil, nil))
		}

		user.AttachToken()
		c.JSON(200, codes.Response(codes.OK, user, nil))
	})

	authRouter.POST("/register", func(c *gin.Context) {
		var newUSer database.NewUser

		if err := c.Bind(&newUSer); err != nil {
			c.JSON(400, codes.Response[error](
				codes.InvalidFormBody,
				nil,
				common.TranslateError("zh_tw", err),
			))
			return
		}

		user, err := db.CreateUser(newUSer)
		if errors.Is(err, codes.UsernameAlreadyExists) {
			c.JSON(400, codes.Response[error](codes.UsernameAlreadyExists, nil, nil))
			return
		}

		user.AttachToken()
		c.JSON(200, codes.Response(codes.OK, user, nil))
	})
}

func GetCurrentUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	tokens := strings.Split(token, " ")
	tokens_len := len(tokens)

	if tokens_len > 2 {
		token = ""
	} else if tokens_len == 2 && tokens[0] == "Bearer" {
		token = tokens[1]
	}

	if token == "" {
		c.JSON(401, codes.Response[error](codes.UnknownToken, nil, nil))
		c.Abort()
		return
	}

	db := database.GetDBFromContext(c)
	if user := db.GetUserFromToken(token); user != nil {
		c.Set("user", user)
		c.Next()
	} else {
		c.JSON(401, codes.Response[error](codes.UnknownToken, nil, nil))
		c.Abort()
	}
}

func registerUsersRouter(container common.Container, app *gin.RouterGroup) {
	usersRouter := app.Group("/users", GetCurrentUser)
	usersRouterMe := usersRouter.Group("/@me")
	usersRouterOther := usersRouter.Group("/:id")

	usersRouterMe.GET("/", func(c *gin.Context) {
		c.JSON(200, codes.Response(codes.OK, c.MustGet("user").(*database.UserInfo), nil))
	})
	usersRouterMe.PATCH("/")
	usersRouterMe.GET("/instances")
	usersRouterMe.GET("/instances/:id/members")

	usersRouterOther.GET("/:id")
	usersRouterOther.PATCH("/:id")
	usersRouterOther.GET("/:id/instances")
	usersRouterOther.GET("/:id/instances/:id/members")
}
