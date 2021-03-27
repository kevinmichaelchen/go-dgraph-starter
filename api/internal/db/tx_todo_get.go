package db

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/MyOrg/todo-api/internal/models"
	"github.com/MyOrg/todo-api/internal/obs"
	todoV1 "github.com/MyOrg/todo-api/pkg/pb/myorg/todo/v1"
	"github.com/golang/protobuf/ptypes"
)

func (tx *todoTransactionImpl) GetTodoByID(ctx context.Context, id string) (*todoV1.Todo, error) {
	ctx, span := obs.NewSpan(ctx, "GetTodoByID")
	defer span.End()

	logger := obs.ToLogger(ctx)

	type response struct {
		Todo []models.Todo `json:"todo"`
	}

	// A query to get a single Todo
	query := `
		query getTodo($id: string) {
			todo(func: eq(dgraph.type, "Todo")) @filter(eq(id, $id)) {
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

	// Run query
	res, err := tx.tx.QueryWithVars(ctx, query, map[string]string{"$id": id})

	// Handle error
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON into response struct
	var r response
	if err := json.Unmarshal(res.Json, &r); err != nil {
		return nil, err
	}

	// Log latency
	logger.Info().Msgf("Retrieved Todo in %s", latency(res))

	if len(r.Todo) == 0 {
		return nil, ErrNotFound
	} else if len(r.Todo) > 1 {
		return nil, errors.New("more than one todo found with that id")
	}

	todo := r.Todo[0]

	createdAt, err := ptypes.TimestampProto(todo.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &todoV1.Todo{
		Id:        todo.ID,
		CreatedAt: createdAt,
		Title:     todo.Title,
		Done:      todo.Done,
		CreatorId: todo.Creator.ID,
	}, nil
}
