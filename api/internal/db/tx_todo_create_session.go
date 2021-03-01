package db

import (
	"context"
	"time"

	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/dgraph-io/dgo/v200/protos/api"
)

func (tx *todoTransactionImpl) CreateTodo(ctx context.Context, item *todoV1.Todo) error {
	ctx, span := obs.NewSpan(ctx, "CreateTodo")
	defer span.End()

	now := time.Now()
	nowStr := now.Format(time.RFC3339)

	logger := obs.ToLogger(ctx)

	// Insert into database
	res, err := tx.tx.Mutate(ctx, &api.Mutation{
		Set: []*api.NQuad{
			nquadStr("_:todo", "dgraph.type", "Todo"),
			nquadStr("_:todo", "title", "Todo 1"),
			nquadStr("_:todo", "created_at", nowStr),
			nquadBool("_:todo", "is_done", false),
			nquadRel("_:todo", "creator", "_:user1"),

			nquadStr("_:user1", "dgraph.type", "User"),
			nquadStr("_:user1", "name", "Alice"),
			nquadStr("_:user1", "created_at", nowStr),
		},
	})

	if err != nil {
		return err
	}

	// TODO Insert into cache

	logger.Info().Msgf("created uids: %v", res.Uids)

	return nil
}
