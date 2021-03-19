package search

import (
	"context"
	"fmt"

	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/meilisearch/meilisearch-go"
)

const (
	indexForTodos = "todos"
)

type Client interface {
	AddOrUpdate(ctx context.Context, todo *todoV1.Todo) error
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

func todoToDocument(ctx context.Context, todo *todoV1.Todo) []map[string]interface{} {
	return []map[string]interface{}{
		{"todo_id": todo.Id, "title": todo.Title},
	}
}
