package common

import (
	"github.com/monkey-panel/control-panel-api/common/codes"
	"github.com/monkey-panel/control-panel-api/common/database"

	"github.com/gin-gonic/gin"
)

func CheckAuthorization(c *gin.Context) {
	database.GetDBFromContext(c)

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(codes.Response[*uint8](codes.Unauthorized, nil, nil))
		c.Abort()
		return
	}
}
