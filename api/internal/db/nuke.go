package db

import (
	"context"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
)

func Nuke(ctx context.Context, dgraphClient *dgo.Dgraph) error {
	return dgraphClient.Alter(ctx, &api.Operation{
		DropAll: true,
	})
}

func NukeData(ctx context.Context, dgraphClient *dgo.Dgraph) error {
	return dgraphClient.Alter(ctx, &api.Operation{
		DropOp: api.Operation_DATA,
	})
}
