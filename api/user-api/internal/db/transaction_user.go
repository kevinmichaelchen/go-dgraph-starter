package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/MyOrg/user-api/internal/db/models"
	"github.com/MyOrg/user-api/internal/obs"
	userV1 "github.com/MyOrg/user-api/pkg/pb/myorg/user/v1"
	"github.com/golang/protobuf/ptypes"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
	// Create new tracing span
	ctx, span := obs.NewSpan(ctx, "GetUser")
	defer span.End()

	// Perform query
	user, err := models.FindUser(ctx, tx.tx, id)

	// Handle error
	if err != nil {
		return nil, err
	}

	// Convert time.Time to Timestamp
	createdAt, err := ptypes.TimestampProto(user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("found user with invalid createdAt timestamp: %w", err)
	}

	// Return payload
	return &userV1.User{
		Id:        user.ID,
		CreatedAt: createdAt,
		Name:      user.Name,
	}, nil
}

func (tx *userTransactionImpl) CreateUser(ctx context.Context, item *userV1.User) error {
	ctx, span := obs.NewSpan(ctx, "CreateUser")
	defer span.End()

	createdAt, err := ptypes.Timestamp(item.CreatedAt)
	if err != nil {
		return fmt.Errorf("found user with invalid createdAt timestamp: %w", err)
	}

	user := models.User{
		ID:        item.Id,
		CreatedAt: createdAt,
		Name:      item.Name,
	}

	return user.Insert(ctx, tx.tx, boil.Infer())
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
