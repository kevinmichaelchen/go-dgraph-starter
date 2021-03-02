package graphql

import (
	"time"

	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	"github.com/graphql-go/graphql"
)

func (s Server) buildFieldForGetTodo(todoType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: todoType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ctx := p.Context
			logger := obs.ToLogger(ctx)
			logger.Info().Msg("Resolving GraphQL field: todo")
			return Todo{"1", time.Now(), "Title", false}, nil
		},
		Description: "Retrieve a Todo object",
	}
}
