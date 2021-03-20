package service

import (
	"os"
	"testing"

	"github.com/MyOrg/go-dgraph-starter/internal/configuration"
	"github.com/MyOrg/go-dgraph-starter/internal/db"
	"github.com/MyOrg/go-dgraph-starter/internal/search"
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
	searchClient := search.NewClient(config.SearchConfig)

	svc = NewService(config, dbClient, searchClient)

	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}
