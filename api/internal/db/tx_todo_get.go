package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/MyOrg/go-dgraph-starter/internal/models"
	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/golang/protobuf/ptypes"
)

func (tx *todoTransactionImpl) GetTodos(ctx context.Context, in *todoV1.GetTodosRequest) (*todoV1.GetTodosResponse, error) {
	ctx, span := obs.NewSpan(ctx, "GetTodos")
	defer span.End()

	logger := obs.ToLogger(ctx)

	// A struct for unmarshalling JSON responses into
	type response struct {
		Todos []models.Todo `json:"todos"`
	}

	// The requested page-size
	base64EncodedCursor, pageSize, isForwardsPagination := getPaginationInfo(in.PaginationRequest)

	var cursorDirection string
	var orderKey string
	if isForwardsPagination {
		cursorDirection = "gt"
		orderKey = "orderasc"
	} else {
		cursorDirection = "lt"
		orderKey = "orderdesc"
	}

	// TODO handle OrderBy

	var cursorField, cursor string
	if c, err := parseCursor(ctx, base64EncodedCursor); err != nil {
		return nil, err
	} else {
		cursorField = c.field
		cursor = c.value
	}

	// A query to get all Todos
	query := fmt.Sprintf(`
		query getTodos($cursor: string, $pageSize: int) {
			todos(func: eq(dgraph.type, "Todo"), %s: created_at, first: $pageSize) @filter(%s(%s, $cursor)) {
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
	`, orderKey, cursorDirection, cursorField)

	// Run query
	res, err := tx.tx.QueryWithVars(ctx, query, map[string]string{
		"$cursor":   cursor,
		"$pageSize": strconv.Itoa(pageSize),
	})

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

	if len(r.Todos) == 0 {
		return &todoV1.GetTodosResponse{
			Edges:      []*todoV1.TodoEdge{},
			PageInfo:   emptyPageInfo(),
			TotalCount: 0,
		}, nil
	}

	var edges []*todoV1.TodoEdge
	for _, todo := range r.Todos {
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
			// TODO created_at may not always be the cursor field, don't hard-code
			Cursor: newCursor(cursorField, todo.CreatedAt.Format(time.RFC3339)).encode(),
			Node:   todoPB,
		})
	}

	numEdges := len(edges)
	var endCursor string
	if numEdges > 0 && edges != nil {
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
