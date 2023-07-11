package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/a3510377/control-panel-api/common"
	"github.com/a3510377/control-panel-api/common/database"
	"github.com/a3510377/control-panel-api/common/utils"
	"github.com/a3510377/control-panel-api/routers"

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
	// generate database
	db, err := database.NewDB("data/db.db")
	if err != nil {
		panic(err)
	}

	container := common.Container{DB: db}
	mode := gin.ReleaseMode
	if len(os.Getenv("DEV")) > 0 {
		mode = gin.DebugMode
	}

	gin.SetMode(mode)
	gin.ForceConsoleColor()

	srv := &http.Server{
		Addr: config.Address,
		Handler: routers.Routers(container, routers.RouterConfig{
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
