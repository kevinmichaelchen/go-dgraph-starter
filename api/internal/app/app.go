package app

import (
	"context"
	"github.com/rs/zerolog/log"
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

	redisClient := db.NewRedisClient(config.RedisConfig)

	// Connect to the database
	dgraphClient := config.DgraphConfig.Connect()
	dbClient := db.NewClient(dgraphClient, redisClient, config)

	// Drop all data and schema
	if err := db.Nuke(context.Background(), dgraphClient); err != nil {
		log.Fatal().Err(err).Msg("failed to nuke database")
	}

	// Build schema
	if err := db.BuildSchema(context.Background(), dgraphClient); err != nil {
		log.Fatal().Err(err).Msg("failed to build database schema")
	}

	svc := service.NewService(config, dbClient)

	var wg sync.WaitGroup

	wg.Add(1)
	grpcServer := grpc.NewServer(config, svc)
	go grpcServer.Run()

	wg.Wait()
}
