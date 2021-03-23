package app

import (
	"sync"

	"github.com/rs/zerolog/log"

	"github.com/MyOrg/user-api/internal/db"
	"github.com/MyOrg/user-api/internal/grpc"

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

	dbConn, err := config.SQLConfig.Connect()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to SQL database")
	}

	dbClient := db.NewClient(dbConn, redisClient)

	svc := service.NewService(config, dbClient)

	var wg sync.WaitGroup

	wg.Add(1)
	grpcServer := grpc.NewServer(config, svc)
	go grpcServer.Run()

	wg.Wait()
}
