package api

import (
	"github.com/a3510377/control-panel-api/utils"
	"github.com/gin-gonic/gin"
)

var router = []func(utils.Container, *gin.RouterGroup){}

func RegisterRouter(container utils.Container, app *gin.RouterGroup) {
	for _, f := range router {
		f(container, app)
	}
}
