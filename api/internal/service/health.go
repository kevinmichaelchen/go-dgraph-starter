package service

import (
	"context"

	health "google.golang.org/grpc/health/grpc_health_v1"
)

func (s Service) Check(ctx context.Context, request *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	// Ping database(s)
	err := s.GetDatabase().Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}
