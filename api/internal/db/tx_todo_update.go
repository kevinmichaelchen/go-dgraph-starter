package db

import (
	"context"

	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/dgraph-io/dgo/v200/protos/api"
)

func (tx *todoTransactionImpl) UpdateTodo(ctx context.Context, request *todoV1.UpdateTodoRequest) (*todoV1.UpdateTodoResponse, error) {
	ctx, span := obs.NewSpan(ctx, "UpdateTodo")
	defer span.End()

	logger := obs.ToLogger(ctx)

	id := request.Id

	query := `
		query getTodo($id: string) {
			todo as var(func: eq(id, $id)) {
				id
				created_at
				title
				done
				creator {
					id
					name
					created_at
				}
			}
		}
	`

	mu := &api.Mutation{
		// Only mutate if we find one entity
		Cond: `@if(eq(len(todo), 1))`,
		Set: []*api.NQuad{
			nquadStr("uid(todo)", "title", request.Title),
			nquadBool("uid(todo)", "is_done", request.Done),
		},
	}
	req := &api.Request{
		Query:     query,
		Vars:      map[string]string{"$id": id},
		Mutations: []*api.Mutation{mu},
	}

	// Perform query
	res, err := tx.tx.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	// Log latency
	logger.Info().Msgf("Updated Todo in %s", latency(res))

	logger.Info().Msgf("Update response: %s", res.Json)

	var todoPB *todoV1.Todo
	if res, err := tx.GetTodoByID(ctx, id); err != nil {
		return nil, err
	} else {
		todoPB = res
	}

	return &todoV1.UpdateTodoResponse{
		Todo: todoPB,
	}, nil
}
