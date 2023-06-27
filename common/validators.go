package common

import (
	"github.com/a3510377/control-panel-api/common/codes"
	"github.com/a3510377/control-panel-api/common/database"

	"github.com/gin-gonic/gin"
)

func CheckAuthorization(c *gin.Context) {
	database.GetDBFromContext(c)

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, codes.Response[*uint8](codes.Unauthorized, nil))
		c.Abort()
		return
	}
}
