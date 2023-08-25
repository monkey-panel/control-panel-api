package api

import "github.com/gin-gonic/gin"


var router = []func(*gin.RouterGroup){}

func RegisterRouter(app *gin.RouterGroup) {
	app.Use(JSONMiddleware)

	for _, f := range router {
		f(app)
	}
}
