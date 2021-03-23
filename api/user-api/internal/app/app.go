package app

import (
	"context"
	"sync"

	"github.com/rs/zerolog/log"

	"github.com/MyOrg/user-api/internal/db"
	"github.com/MyOrg/user-api/internal/graphql"
	"github.com/MyOrg/user-api/internal/grpc"
	"github.com/MyOrg/user-api/internal/search"

	"github.com/MyOrg/user-api/internal/configuration"
	"github.com/MyOrg/user-api/internal/service"
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

	log.Info().Msg("Connecting to Redis...")
	redisClient := db.NewRedisClient(config.RedisConfig)

	// Connect to the database
	log.Info().Msg("Connecting to Dgraph...")
	dgraphClient := config.DgraphConfig.Connect()
	dbClient := db.NewClient(dgraphClient, redisClient, config)

	// Drop all data and schema
	log.Info().Msg("Dropping all Dgraph data...")
	if err := db.NukeDataAndSchema(context.Background(), dgraphClient); err != nil {
		log.Fatal().Err(err).Msg("failed to nuke database")
	}

	// Build schema
	log.Info().Msg("Building Dgraph schema...")
	if err := db.BuildSchema(context.Background(), dgraphClient); err != nil {
		log.Fatal().Err(err).Msg("failed to build database schema")
	}

	svc := service.NewService(config, dbClient, searchClient)

	var wg sync.WaitGroup

	wg.Add(1)
	grpcServer := grpc.NewServer(config, svc)
	go grpcServer.Run()

	wg.Add(1)
	graphqlServer := graphql.NewServer(config, svc)
	go graphqlServer.Run()

	wg.Wait()
}
