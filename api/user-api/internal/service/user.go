package service

import (
	"context"

	"github.com/MyOrg/user-api/internal/db"
	"github.com/MyOrg/user-api/internal/obs"
	userV1 "github.com/MyOrg/user-api/pkg/pb/myorg/user/v1"
)

func (s Service) GetUser(ctx context.Context, request *userV1.GetUserRequest) (*userV1.GetUserResponse, error) {
	ctx, span := obs.NewSpan(ctx, "GetUser")
	defer span.End()

	var userPB *userV1.User

	// Perform database query
	err := s.dbClient.RunInReadOnlyTransaction(ctx, func(ctx context.Context, tx db.Transaction) error {
		if user, err := tx.GetUser(ctx, request.Id); err != nil {
			return err
		} else {
			userPB = user
		}

		return nil
	})

	// Handle error
	if err != nil {
		return nil, err
	}

	return &userV1.GetUserResponse{
		User: userPB,
	}, nil
}
