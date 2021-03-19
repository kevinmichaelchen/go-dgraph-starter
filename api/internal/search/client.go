package search

import (
	"context"
	"errors"
	"fmt"

	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/meilisearch/meilisearch-go"
)

const (
	indexForTodos = "todos"
)

type TodoID string

type Client interface {
	AddOrUpdate(ctx context.Context, todo *todoV1.Todo) error
	Query(ctx context.Context, query string) ([]TodoID, error)
}

type impl struct {
	client meilisearch.ClientInterface
}

func NewClient(config Config) Client {
	client := config.NewClient()
	return &impl{client: client}
}

func (i *impl) AddOrUpdate(ctx context.Context, todo *todoV1.Todo) error {
	logger := obs.ToLogger(ctx)

	documents := todoToDocument(ctx, todo)

	res, err := i.client.Documents(indexForTodos).AddOrUpdate(documents)
	if err != nil {
		return fmt.Errorf("failed to AddOrUpdate documents to search index: %w", err)
	}

	logger.Info().Msgf("Updated Meilisearch index: %d", res.UpdateID)

	return nil
}

func (i *impl) Query(ctx context.Context, query string) ([]TodoID, error) {
	// Search the index
	res, err := i.client.Search(indexForTodos).Search(meilisearch.SearchRequest{
		Query: query,
		Limit: 10,
	})

	// Handle error
	if err != nil {
		return nil, fmt.Errorf("failed to search index: %w", err)
	}

	todos, err := hitsToProtobufs(ctx, res.Hits)
	if err != nil {
		return nil, fmt.Errorf("failed to convert search results to protobufs: %w", err)
	}

	var ids []TodoID
	for _, todo := range todos {
		ids = append(ids, TodoID(todo.Id))
	}

	return ids, nil
}

func todoToDocument(ctx context.Context, todo *todoV1.Todo) []map[string]interface{} {
	return []map[string]interface{}{
		{"todo_id": todo.Id, "title": todo.Title},
	}
}

func hitsToProtobufs(ctx context.Context, hits []interface{}) ([]*todoV1.Todo, error) {
	var out []*todoV1.Todo

	for _, hit := range hits {
		pb := &todoV1.Todo{}
		if hitMap, ok := hit.(map[string]interface{}); !ok {
			return nil, errors.New("could not convert hit to map")
		} else {

			// Get the ID
			if id, ok := hitMap["todo_id"]; !ok {
				return nil, errors.New("todo_id not present in search hit")
			} else {
				if idString, ok := id.(string); !ok {
					return nil, errors.New("todo_id not in search hit is not a string")
				} else {
					pb.Id = idString
				}
			}

			// Get the title
			if title, ok := hitMap["title"]; !ok {
				return nil, errors.New("title not present in search hit")
			} else {
				if titleString, ok := title.(string); !ok {
					return nil, errors.New("title not in search hit is not a string")
				} else {
					pb.Title = titleString
				}
			}

		}

		out = append(out, pb)
	}
	return out, nil
}
