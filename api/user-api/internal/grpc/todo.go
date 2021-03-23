package grpc

import (
	"context"

	userV1 "github.com/MyOrg/user-api/pkg/pb/myorg/user/v1"
)

func (s Server) GetUser(ctx context.Context, request *todoV1.GetUserRequest) (*todoV1.GetUserResponse, error) {
	return s.service.GetUser(ctx, request)
}
