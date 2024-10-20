package main

import (
	"log"

	"github.com/desaias/argocd-demo/config"
	"github.com/desaias/argocd-demo/internal/app/hello"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error("Error syncing logger", zap.Error(err))
		}
	}()

	var env config.Config
	err := config.LoadConfig(&env)
	if err != nil {
		log.Panicf("Error loading config: %v\n", err)
	}
	app := hello.NewHelloService(env, logger)
	err = app.Run()
	if err != nil {
		log.Panicf("Error running server: %v\n", err)
	}
}
