package grpc

import (
	"fmt"
	"net"

	"github.com/rs/zerolog/log"

	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"

	health "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/MyOrg/go-dgraph-starter/internal/configuration"
	"github.com/MyOrg/go-dgraph-starter/internal/service"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	config  configuration.Config
	service service.Service
}

func (s Server) String() string {
	return "grpc.Server"
}

func NewServer(config configuration.Config, service service.Service) Server {
	return Server{
		config:  config,
		service: service,
	}
}

func (s Server) Run() {
	address := fmt.Sprintf(":%d", s.config.GrpcPort)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to listen on address: %s", address)
	}

	log.Info().Msgf("Starting %s on %s...", s, address)
	grpcServer := newServer()

	todoV1.RegisterTodoServiceServer(grpcServer, s)
	health.RegisterHealthServer(grpcServer, s)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Msgf("Failed to serve %s on address: %s", s, address)
	}
}
