package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/MyOrg/user-api/internal/obs"
	userV1 "github.com/MyOrg/user-api/pkg/pb/myorg/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserTransaction interface {
	GetUser(ctx context.Context, id string) (*userV1.User, error)
	CreateUser(ctx context.Context, item *userV1.User) error
}

type userTransactionImpl struct {
	tx          *sql.Tx
	redisClient RedisClient
}

func (tx *userTransactionImpl) GetUser(ctx context.Context, id string) (*userV1.User, error) {
	ctx, span := obs.NewSpan(ctx, "GetUser")
	defer span.End()

	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func (tx *userTransactionImpl) CreateUser(ctx context.Context, item *userV1.User) error {
	ctx, span := obs.NewSpan(ctx, "CreateUser")
	defer span.End()

	return status.Error(codes.Unimplemented, "Unimplemented")
}

func (tx *userTransactionImpl) cacheTodo(ctx context.Context, item *userV1.User) error {
	ctx, span := obs.NewSpan(ctx, "cacheTodo")
	defer span.End()

	longevity := time.Hour * 24
	return tx.redisClient.Set(ctx, redisKeyForUser(item.Id), item, longevity)
}

func redisKeyForUser(id string) string {
	return fmt.Sprintf("user-%s", id)
}
