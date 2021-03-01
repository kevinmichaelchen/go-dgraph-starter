package db

import (
	"context"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
)

const schema = `
  title: string @index(term) .
  name: string @index(term) .
  created_at: datetime @index(day) .
  is_done: bool .
  creator: uid .

  type Todo {
    created_at
    title
    is_done
    creator
  }

  type User {
    name
    created_at
  }
`

func BuildSchema(ctx context.Context, dgraphClient *dgo.Dgraph) error {
	// Add types and indexes to schema
	op := &api.Operation{
		Schema: schema,
	}
	return dgraphClient.Alter(ctx, op)
}
