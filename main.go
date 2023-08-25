package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/monkey-panel/control-panel-api/common/utils"
	"github.com/monkey-panel/control-panel-api/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	os.MkdirAll("data", os.ModePerm)

	// generate cert
	config := utils.Config()
	if config.EnableTLS && (!utils.HasFile("data/server.pem") || !utils.HasFile("data/server.key")) {
		ca, privateKey := utils.GenerateCACertificate()
		ssl := utils.GenerateCertificate(ca, privateKey, []string{"console-panel-api"})
		utils.AutoWriteFile("data/server.pem", ssl.ServerKey, os.ModePerm)
		utils.AutoWriteFile("data/server.key", ssl.ServerPem, os.ModePerm)
	}

	mode := gin.ReleaseMode
	if len(os.Getenv("DEV")) > 0 {
		mode = gin.DebugMode
	}

	gin.SetMode(mode)
	gin.ForceConsoleColor()

	srv := &http.Server{
		Addr: config.Address,
		Handler: routers.Routers(routers.RouterConfig{
			AllowOrigins: config.AllowOrigins,
		}),
	}

	go func() {
		log.Println("Server starting...")
		var err error

		if config.EnableTLS {
			err = srv.ListenAndServeTLS("data/server.pem", "data/server.key")
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
