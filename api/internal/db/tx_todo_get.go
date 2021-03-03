package db

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/MyOrg/go-dgraph-starter/internal/models"
	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/golang/protobuf/ptypes"
)

func (tx *todoTransactionImpl) GetTodoByID(ctx context.Context, id string) (*todoV1.Todo, error) {
	ctx, span := obs.NewSpan(ctx, "GetTodoByID")
	defer span.End()

	logger := obs.ToLogger(ctx)

	type response struct {
		Todo []models.Todo `json:"todo"`
	}

	query := `
		query getTodo($id: string) {
			todo(func: eq(dgraph.type, "Todo")) @filter(eq(id, $id)) {
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

	res, err := tx.tx.QueryWithVars(ctx, query, map[string]string{"$id": id})
	if err != nil {
		return nil, err
	}

	var r response
	if err := json.Unmarshal(res.Json, &r); err != nil {
		return nil, err
	}

	logger.Info().Msgf("Retrieved Todo in %s", latency(res))

	if len(r.Todo) == 0 {
		return nil, errors.New("no todo found")
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
		AuthorId:  todo.Creator.ID,
	}, nil
}
