package service

import (
	"os"
	"testing"

	"github.com/MyOrg/todo-api/internal/configuration"
	"github.com/MyOrg/todo-api/internal/db"
	"github.com/MyOrg/todo-api/internal/search"
	"github.com/dgraph-io/dgo/v200"
	"github.com/rs/zerolog/log"
)

var dgraphClient *dgo.Dgraph
var svc Service

func TestMain(m *testing.M) {
	config := configuration.LoadConfig()

	log.Info().Msg("Connecting to Redis...")
	redisClient := db.NewRedisClient(config.RedisConfig)

	// Connect to the database
	log.Info().Msg("Connecting to Dgraph...")
	dgraphClient = config.DgraphConfig.Connect()
	dbClient := db.NewClient(dgraphClient, redisClient, config)

	// Connect to search index
	searchClient := search.NewClient(config.SearchConfig)

	// Dial gRPC connection to Users service
	usersConn, err := config.UsersConfig.Dial()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to Users service")
	}

	svc = NewService(config, dbClient, searchClient, usersConn)

	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}
