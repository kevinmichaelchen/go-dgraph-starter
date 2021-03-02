package graphql

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/MyOrg/go-dgraph-starter/internal/configuration"
	"github.com/MyOrg/go-dgraph-starter/internal/service"
	"github.com/graphql-go/handler"
)

type Server struct {
	config  configuration.Config
	service service.Service
}

func (s Server) String() string {
	return "graphql.Server"
}

func NewServer(config configuration.Config, service service.Service) Server {
	return Server{
		config:  config,
		service: service,
	}
}

func (s Server) Run() {
	h := handler.New(&handler.Config{
		Schema:   buildSchema(),
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)

	address := fmt.Sprintf(":%d", s.config.HTTPPort)

	log.Info().Msgf("Starting %s on %s...", s, address)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal().Err(err).Msgf("Failed to serve %s on address: %s", s, address)
	}
}
