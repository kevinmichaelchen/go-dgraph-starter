package graphql

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"github.com/MyOrg/todo-api/internal/configuration"
	"github.com/MyOrg/todo-api/internal/service"
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
	r := mux.NewRouter()

	graphqlHandler := handler.New(&handler.Config{
		Schema:   s.buildSchema(),
		Pretty:   true,
		GraphiQL: true,
	})

	wrapWithCors := handlers.CORS(
		handlers.AllowedOrigins([]string{
			"http://localhost:3000",
		}),
		handlers.AllowedMethods([]string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			// http.MethodOptions,
		}),
		handlers.AllowedHeaders([]string{
			"Accept",
			"Accept-Encoding",
			"Accept-Language",
			"Access-Control-Request-Headers",
			"Access-Control-Request-Method",
			"Authorization",
			"Connection",
			"Content-Language",
			"Content-Type",
			"Host",
			"Origin",
			"Referer",
			"User-Agent",
		}),
		handlers.AllowCredentials(),
	)

	r.Handle("/graphql", graphqlHandler)

	address := fmt.Sprintf(":%d", s.config.HTTPPort)

	log.Info().Msgf("Starting %s on %s...", s, address)

	if err := http.ListenAndServe(address, wrapWithCors(r)); err != nil {
		log.Fatal().Err(err).Msgf("Failed to serve %s on address: %s", s, address)
	}
}
