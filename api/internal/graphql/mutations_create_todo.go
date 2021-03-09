package graphql

import (
	"fmt"

	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/graphql-go/graphql"
)

func (s Server) buildFieldForCreateTodo(todoType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        todoType,
		Description: "Create new Todo",
		Args: graphql.FieldConfigArgument{
			argTitle: &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ctx := p.Context
			logger := obs.ToLogger(ctx)

			args := p.Args
			logger.Info().Msgf("Received GraphQL request to %s with args: %v", p.Info.FieldName, args)

			// Build the request protobuf from the GraphQL args
			request, err := buildCreateTodoRequestFromArgs(args)
			if err != nil {
				return nil, err
			}

			// Call the service
			res, err := s.service.CreateTodo(ctx, request)
			if err != nil {
				return nil, err
			}

			// Build the response protobuf and return it
			return buildTodo(res.Todo)
		},
	}
}

func buildCreateTodoRequestFromArgs(args map[string]interface{}) (*todoV1.CreateTodoRequest, error) {
	if value, ok := args[argTitle]; ok {
		if val, ok := value.(string); ok {
			return &todoV1.CreateTodoRequest{
				Title: val,
			}, nil
		} else {
			return nil, fmt.Errorf("'%s' not a string", argTitle)
		}
	} else {
		return nil, fmt.Errorf("must specify '%s'", argTitle)
	}
}
