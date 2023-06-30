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
	db, err := database.NewDB("data/db.db")
	if err != nil {
		panic(err)
	}

	container := common.Container{DB: db}
	mode := gin.ReleaseMode
	if len(os.Getenv("DEV")) > 0 {
		mode = gin.DebugMode
	}

	if !utils.HasFile("data/server.pem") || !utils.HasFile("data/server.key") {
		utils.SummonCert()
	}
	gin.SetMode(mode)
	gin.ForceConsoleColor()

	app := routers.Routers(container, routers.RouterConfig{
		AllowOrigins: []string{"http://localhost:8080"}, // TODO: Add the allow origins config
	})

	srv := &http.Server{
		Addr:    "127.0.0.1:8000", // TODO: Add the address config
		Handler: app,
	}

	go func() {
		log.Println("Server starting...")
		err := srv.ListenAndServeTLS("data/server.pem", "data/server.key")
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
