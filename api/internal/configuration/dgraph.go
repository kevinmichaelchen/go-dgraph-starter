package configuration

import (
	"fmt"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type DgraphConfig struct {
	Host string
	Port int
}

func LoadDgraphConfig() DgraphConfig {
	c := DgraphConfig{
		Host: "localhost",
		Port: 9080,
	}
	return c
}

func (c DgraphConfig) Connect() *dgo.Dgraph {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)

	log.Info().Msgf("Connecting to Dgraph at address: %s", addr)

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to Dgraph")
	}

	return dgo.NewDgraphClient(api.NewDgraphClient(conn))
}
