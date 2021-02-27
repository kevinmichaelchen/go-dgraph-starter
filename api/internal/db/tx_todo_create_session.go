package db

import (
	"context"
	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
)

func (tx *todoTransactionImpl) CreateTodo(ctx context.Context, item *todoV1.Todo) error {
	ctx, span := obs.NewSpan(ctx, "CreateTodo")
	defer span.End()

	// Insert into database
	// Insert into cache
	return nil
}
