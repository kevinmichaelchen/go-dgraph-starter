package app

import (
	"sync"

	"github.com/MyOrg/go-dgraph-starter/internal/db"
	"github.com/MyOrg/go-dgraph-starter/internal/grpc"

	"github.com/MyOrg/go-dgraph-starter/internal/configuration"
	"github.com/MyOrg/go-dgraph-starter/internal/service"
)

type App struct {
	config configuration.Config
}

func NewApp(c configuration.Config) App {
	return App{
		config: c,
	}
}

func (a App) Run() {
	config := a.config

	//redisClient := db.NewRedisClient(config.RedisConfig)

	// Connect to the database
	var dbClient db.Client
	//dbClient := db.NewClient(gormDB, redisClient)

	svc := service.NewService(config, dbClient)

	var wg sync.WaitGroup

	wg.Add(1)
	grpcServer := grpc.NewServer(config, svc)
	go grpcServer.Run()

	wg.Wait()
}
