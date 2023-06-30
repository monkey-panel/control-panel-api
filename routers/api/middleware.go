package api

import (
	"strings"

	"github.com/a3510377/control-panel-api/common/codes"
	"github.com/a3510377/control-panel-api/common/database"
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(c *gin.Context) {
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

func GetUserFromContext(c *gin.Context) *database.UserInfo {
	return c.MustGet("user").(*database.UserInfo)
}

func JSONMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
}
