package routers

import (
	"strings"
	"time"

	"github.com/a3510377/control-panel-api/routers/api"
	"github.com/a3510377/control-panel-api/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	NoRouteHandlers []gin.HandlerFunc
	AllowOrigins    []string
}

func Routers(container utils.Container, config RouterConfig) *gin.Engine {
	app := gin.Default()
	app.Use(cors.New(corsConfig(config)))

	app.NoRoute(append(config.NoRouteHandlers, func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.JSON(404, gin.H{
				"code":    404,
				"message": "Not Found",
			})
			c.Abort()
		}
	})...)

	apiRouter := app.Group("/api")
	api.RegisterRouter(container, apiRouter)
	api.RegisterRouter(container, apiRouter.Group("/v1"))

	return app
}

func corsConfig(router_config RouterConfig) cors.Config {
	config := cors.DefaultConfig()

	if gin.Mode() == gin.DebugMode {
		config.AllowAllOrigins = true
		config.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
		config.AllowHeaders = []string{
			"Authorization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host",
			"Access-Control-Request-Method", "Access-Control-Request-Headers",
		}
	} else {
		config.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
		config.AllowHeaders = []string{
			"Authorization", "Content-Type", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host",
		}
		config.AllowOrigins = router_config.AllowOrigins
	}

	config.MaxAge = 1 * time.Hour
	config.AllowCredentials = true
	config.ExposeHeaders = []string{"Content-Length"}

	return config
}
