package api

import (
	"github.com/monkey-panel/control-panel-api/common"

	"github.com/gin-gonic/gin"
)

var router = []func(common.Container, *gin.RouterGroup){}

func RegisterRouter(container common.Container, app *gin.RouterGroup) {
	app.Use(JSONMiddleware)

	for _, f := range router {
		f(container, app)
	}
}
