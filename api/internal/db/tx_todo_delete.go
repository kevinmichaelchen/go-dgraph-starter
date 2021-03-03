package db

import (
	"context"
	"encoding/json"

	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/dgraph-io/dgo/v200/protos/api"
)

func (tx *todoTransactionImpl) DeleteTodo(ctx context.Context, id string) (*todoV1.DeleteTodoResponse, error) {
	ctx, span := obs.NewSpan(ctx, "CreateTodo")
	defer span.End()

	logger := obs.ToLogger(ctx)

	type todoResponse struct {
		UID string `json:"uid"`
	}
	type getTodoResponse struct {
		Todo []todoResponse `json:"todo"`
	}

	var uid string

	if res, err := tx.tx.QueryWithVars(ctx, `
		query getTodo($id: string) {
			todo(func: eq(id, $id)) {
				uid
			}
		}
		`, map[string]string{"$id": id}); err != nil {
		return nil, err
	} else {
		logger.Info().Msgf("Received JSON %s in %s", string(res.Json), latency(res))
		var resp getTodoResponse
		if err := json.Unmarshal(res.Json, &resp); err != nil {
			return nil, err
		} else {
			if len(resp.Todo) == 0 {
				logger.Info().Msg("No Todo found to delete")
				// IDEMPOTENCE
				return &todoV1.DeleteTodoResponse{}, nil
			}
			uid = resp.Todo[0].UID
			logger.Info().Msgf("Deleting Todo %s", uid)
		}
	}

	if res, err := tx.tx.Mutate(ctx, &api.Mutation{
		Del: []*api.NQuad{
			nquadAll(uid),
		},
	}); err != nil {
		return nil, err
	} else {
		logger.Info().Msgf("Deleted Todo in %s", latency(res))
	}

	// TODO Delete from cache

	return &todoV1.DeleteTodoResponse{}, nil
}
