package db

import (
	"context"
	"fmt"
	"github.com/dgraph-io/dgo/v200"
	"time"

	"github.com/MyOrg/go-dgraph-starter/internal/models"
	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
)

type TodoTransaction interface {
	GetTodoByID(ctx context.Context, id string) (*models.Todo, error)
	CreateTodo(ctx context.Context, item *todoV1.Todo) error
}

type todoTransactionImpl struct {
	tx          *dgo.Txn
	redisClient RedisClient
}

func redisKeyForTodo(id string) string {
	return fmt.Sprintf("todo-%s", id)
}

func (tx *todoTransactionImpl) cacheTodo(ctx context.Context, item *models.Todo) error {
	ctx, span := obs.NewSpan(ctx, "cacheTodo")
	defer span.End()

	longevity := time.Hour * 24
	return tx.redisClient.Set(ctx, redisKeyForTodo(item.ID), item, longevity)
}
