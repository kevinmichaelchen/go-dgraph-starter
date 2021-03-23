package grpc

import (
	"context"

	"github.com/rs/zerolog/log"

	health "google.golang.org/grpc/health/grpc_health_v1"
)

func (s Server) Check(ctx context.Context, request *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return s.service.Check(ctx, request)
}

// rpc Watch(HealthCheckRequest) returns (stream HealthCheckResponse);
func (s Server) Watch(request *health.HealthCheckRequest, stream health.Health_WatchServer) error {
	ctx := stream.Context()

	// Ping database(s)
	err := s.service.GetDatabase().Ping(ctx)
	if err != nil {
		// Send unhealthy
		if err := stream.Send(&health.HealthCheckResponse{
			Status: health.HealthCheckResponse_NOT_SERVING,
		}); err != nil {
			log.Error().Err(err).Msg("failed to ping database")
		}
	}

	// Send healthy
	if err := stream.Send(&health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}); err != nil {
		log.Error().Err(err).Msg("failed to send response into gRPC stream")
	}

	return nil
}
