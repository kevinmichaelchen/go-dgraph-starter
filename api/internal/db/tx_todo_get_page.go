package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/MyOrg/todo-api/internal/models"
	"github.com/MyOrg/todo-api/internal/obs"
	todoV1 "github.com/MyOrg/todo-api/pkg/pb/myorg/todo/v1"
	"github.com/golang/protobuf/ptypes"
)

func (tx *todoTransactionImpl) GetTodos(ctx context.Context, in *todoV1.GetTodosRequest) (*todoV1.GetTodosResponse, error) {
	ctx, span := obs.NewSpan(ctx, "GetTodos")
	defer span.End()

	logger := obs.ToLogger(ctx)

	// Struct for unmarshalling JSON responses into
	type countContainer struct {
		Count int `json:"count"`
	}
	type response struct {
		Todos      []models.Todo    `json:"todos"`
		TotalCount []countContainer `json:"totalCount"`
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
				is_done
				creator {
					id
					name
					created_at
				}
			}
			totalCount(func: eq(dgraph.type, "Todo")) {
				count(uid)
			}
		}
	`,
		// e.g., "orderasc"
		orderKey,
		// e.g., "gt"
		cursorDirection,
		// e.g., "created_at"
		cursorField,
	)

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

	var totalCount int
	if len(r.TotalCount) == 0 {
		return nil, errors.New("empty count array; could not get total count")
	} else {
		totalCount = r.TotalCount[0].Count
	}

	// Log latency
	logger.Info().Msgf("Retrieved Todos in %s", latency(res))

	// Check for an empty response
	if len(r.Todos) == 0 {
		return &todoV1.GetTodosResponse{
			Edges:      []*todoV1.TodoEdge{},
			PageInfo:   emptyPageInfo(),
			TotalCount: int32(totalCount),
		}, nil
	}

	// We're conforming to the Relay "Cursor Connections Specification", so we use "edges" and "nodes".
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
			CreatorId: todo.Creator.ID,
		}

		cursor, err := newCursor(cursorField, todoPB)
		if err != nil {
			return nil, err
		}

		edges = append(edges, &todoV1.TodoEdge{
			Cursor: cursor.encode(),
			Node:   todoPB,
		})
	}

	// Get page info
	numEdges := len(edges)
	var startCursor, endCursor string
	if numEdges > 0 && edges != nil {
		startCursor = edges[0].Cursor
		endCursor = edges[numEdges-1].Cursor
	}

	// Check if there are more pages to fetch
	hasNextPage, err := tx.hasNextPage(ctx, orderKey, cursorDirection, cursorField, edges[numEdges-1].Node)
	if err != nil {
		return nil, err
	}

	return &todoV1.GetTodosResponse{
		Edges: edges,
		PageInfo: &todoV1.PageInfo{
			StartCursor: startCursor,
			EndCursor:   endCursor,
			HasNextPage: hasNextPage,
		},
		TotalCount: int32(totalCount),
	}, nil
}

func (tx *todoTransactionImpl) hasNextPage(ctx context.Context, orderKey, cursorDirection, cursorField string, todoPB *todoV1.Todo) (bool, error) {
	type response struct {
		Todos []models.Todo `json:"todos"`
	}

	// A query to see if there are more pages to retrieve
	query := fmt.Sprintf(`
		query getTodos($cursor: string, $pageSize: int) {
			todos(func: eq(dgraph.type, "Todo"), %s: created_at, first: 1) @filter(%s(%s, $cursor)) {
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
	`,
		// e.g., "orderasc"
		orderKey,
		// e.g., "gt"
		cursorDirection,
		// e.g., "created_at"
		cursorField,
	)

	// Create cursor
	cursor, err := newCursor(cursorField, todoPB)
	if err != nil {
		return false, err
	}

	// Run query
	res, err := tx.tx.QueryWithVars(ctx, query, map[string]string{
		"$cursor": cursor.value,
	})

	// Handle error
	if err != nil {
		return false, fmt.Errorf("failed to query if there are more pages: %w", err)
	}

	// Unmarshal JSON into response struct
	var r response
	if err := json.Unmarshal(res.Json, &r); err != nil {
		return false, fmt.Errorf("failed to unmarshal response to check if there is more to page: %w", err)
	}

	return len(r.Todos) > 0, nil
}
