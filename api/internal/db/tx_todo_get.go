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

func (tx *todoTransactionImpl) GetTodos(ctx context.Context, in *todoV1.GetTodosRequest) (*todoV1.GetTodosResponse, error) {
	ctx, span := obs.NewSpan(ctx, "GetTodos")
	defer span.End()

	logger := obs.ToLogger(ctx)

	// The requested page-size (or the default page size if the client input was invalid)
	// TODO do not hard-code
	pageSize := 20

	// A struct for unmarshalling JSON responses into
	type response struct {
		Todo []models.Todo `json:"todo"`
	}

	// A query to get all Todos
	query := `
		query getTodo($id: string) {
			todo(func: eq(dgraph.type, "Todo")) {
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

	// Run query
	res, err := tx.tx.Query(ctx, query)

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
	logger.Info().Msgf("Retrieved Todos in %s", latency(res))

	var edges []*todoV1.TodoEdge
	for _, todo := range r.Todo {
		createdAt, err := ptypes.TimestampProto(todo.CreatedAt)
		if err != nil {
			return nil, err
		}

		todoPB := &todoV1.Todo{
			Id:        todo.ID,
			CreatedAt: createdAt,
			Title:     todo.Title,
			Done:      todo.Done,
			AuthorId:  todo.Creator.ID,
		}
		edges = append(edges, &todoV1.TodoEdge{
			Cursor: todoPB.Id,
			Node:   todoPB,
		})
	}

	numEdges := len(edges)
	var endCursor string
	if numEdges > 0 {
		endCursor = edges[numEdges-1].Cursor
	}

	return &todoV1.GetTodosResponse{
		Edges: edges,
		PageInfo: &todoV1.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: numEdges < pageSize,
		},
		TotalCount: int32(numEdges),
	}, nil
}

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
				done
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
		AuthorId:  todo.Creator.ID,
	}, nil
}
