package service

import (
	"context"

	"github.com/MyOrg/todo-api/internal/db"
	"github.com/MyOrg/todo-api/internal/search"
)

func (s Service) DropAllData(ctx context.Context) error {
	if _, err := s.searchClient.GetClient().Documents(search.IndexForTodos).DeleteAllDocuments(); err != nil {
		return err
	}
	return db.NukeDataButNotSchema(ctx, s.dbClient.GetClient())
}
