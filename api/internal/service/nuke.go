package service

import (
	"context"

	"github.com/MyOrg/go-dgraph-starter/internal/db"
)

func (s Service) DropAllData(ctx context.Context) error {
	return db.NukeData(ctx, s.dbClient.GetClient())
}
