package search

import (
	"context"
	"fmt"

	"github.com/meilisearch/meilisearch-go"
)

func CreateIndexes(ctx context.Context, client meilisearch.ClientInterface) error {
	// Create an index if your index does not already exist
	if _, err := client.Indexes().Create(meilisearch.CreateIndexRequest{
		UID: indexForTodos,
	}); err != nil {
		return fmt.Errorf("failed to create search index: %w", err)
	}

	// Not all attributes should be searchable (i.e., matched for query words), e.g., the primary key.
	if _, err := client.Settings(indexForTodos).UpdateSearchableAttributes([]string{
		attributeTitle,
	}); err != nil {
		return fmt.Errorf("failed to update searchable attributes: %w", err)
	}

	return nil
}
