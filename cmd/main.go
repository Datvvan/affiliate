package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/datvvan/affiliate/api"
	"github.com/datvvan/affiliate/config"
	"github.com/datvvan/affiliate/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	defaultConfigPath = "./app.env"
)

var (
	engine *gin.Engine
	ctx    context.Context
	cancel context.CancelFunc
)

func init() {
	ctx, cancel = context.WithCancel(context.Background())

	// init config
	config.Init(defaultConfigPath)

	// connect db
	_, err := db.New()
	if err != nil {
		log.Fatalln("Connect db error: ", err.Error())
	}

	engine = gin.Default()
	engine.Use(cors.Default())
}

func setupGracefulShutdown(ctx context.Context, port string, engine *gin.Engine) {
	signalForExit := make(chan os.Signal, 1)
	signal.Notify(signalForExit, os.Interrupt)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: engine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.WithFields(log.Fields{"bind": port}).Info("Running application")
	stop := <-signalForExit
	log.Info("Stop signal received ", stop)
	if err := srv.Shutdown(ctx); err != nil {
		log.Info("engine.Shutdown err ", err)
	}
}

func main() {
	api.RegisterAPI(engine)

	setupGracefulShutdown(ctx, config.Default.Port, engine)
	cancel()
}
