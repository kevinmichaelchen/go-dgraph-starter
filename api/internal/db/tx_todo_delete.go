package db

import (
	"context"

	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/dgraph-io/dgo/v200/protos/api"
)

func (tx *todoTransactionImpl) DeleteTodo(ctx context.Context, id string) (*todoV1.DeleteTodoResponse, error) {
	ctx, span := obs.NewSpan(ctx, "CreateTodo")
	defer span.End()

	logger := obs.ToLogger(ctx)

	res, err := tx.tx.Do(ctx, &api.Request{
		Query: `
		query getTodo($id: string) {
			todo as var(func: eq(id, $id))
		}
		`,
		Mutations: []*api.Mutation{
			{
				// Only delete if we find one Todo with that ID.
				// Otherwise, we'll assume there are no Todos with that ID,
				// and we don't perform the Deletion, thus ensuring Idempotence.
				Cond: `@if(eq(len(todo), 1))`,
				Del: []*api.NQuad{
					// The pattern S * * deletes all the known edges out of a node,
					// any reverse edges corresponding to the removed edges,
					// and any indexing for the removed data.
					nquadStr("uid(todo)", "*", "*"),
				},
			},
		},
		Vars: map[string]string{
			"$id": id,
		},
	})

	if err != nil {
		return nil, err
	}

	// TODO Delete from cache

	logger.Info().Msgf("Deleted Todo in %s", latency(res))

	return &todoV1.DeleteTodoResponse{}, nil
}
