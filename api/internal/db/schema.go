package db

import (
	"context"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
)

const schema = `
  id: string @index(exact) .
  title: string @index(term) .
  name: string @index(term) .
  created_at: datetime @index(day) .
  is_done: bool .
  creator: uid .

  type Todo {
    id
    created_at
    title
    is_done
    creator
  }

  type User {
    id
    name
    created_at
  }

  event_type: string @index(exact) .
  event_at: datetime @index(hour) .
  is_published_to_search_index: bool .

  type TodoEvent {
    event_type
    event_at
    is_published_to_search_index
    id
    created_at
    title
    is_done
    creator
  }
`

func BuildSchema(ctx context.Context, dgraphClient *dgo.Dgraph) error {
	// Add types and indexes to schema
	op := &api.Operation{
		Schema: schema,
	}
	return dgraphClient.Alter(ctx, op)
}
