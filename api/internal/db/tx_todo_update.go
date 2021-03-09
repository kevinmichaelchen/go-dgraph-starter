package db

import (
	"context"

	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/dgraph-io/dgo/v200/protos/api"
)

func pathMap(in []string) map[string]bool {
	out := make(map[string]bool)
	for _, e := range in {
		out[e] = true
	}
	return out
}

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
				is_done
				creator {
					id
					name
					created_at
				}
			}
		}
	`

	// Check our FieldMask to see which fields we want to update
	var nquads []*api.NQuad
	fields := pathMap(request.FieldMask.Paths)

	if _, ok := fields["title"]; ok {
		nquads = append(nquads, nquadStr("uid(todo)", "title", request.Title))
	}

	if _, ok := fields["done"]; ok {
		nquads = append(nquads, nquadBool("uid(todo)", "is_done", request.Done))
	}

	if len(nquads) > 0 {
		mu := &api.Mutation{
			// Only mutate if we find one entity
			Cond: `@if(eq(len(todo), 1))`,
			Set:  nquads,
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
	}

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
