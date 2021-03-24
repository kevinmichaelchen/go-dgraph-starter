package graphql

import (
	"fmt"

	"github.com/MyOrg/todo-api/internal/obs"
	todoV1 "github.com/MyOrg/todo-api/pkg/pb/myorg/todo/v1"
	"github.com/graphql-go/graphql"
)

func (s Server) buildFieldForGetTodo(todoType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type: todoType,
		Args: graphql.FieldConfigArgument{
			argID: &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "The Todo's ID",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ctx := p.Context
			logger := obs.ToLogger(ctx)

			args := p.Args
			logger.Info().Msgf("Received GraphQL request to %s with args: %v", p.Info.FieldName, args)

			// Build the request protobuf from the GraphQL args
			request, err := buildGetTodoRequestFromArgs(args)
			if err != nil {
				return nil, err
			}

			// Call the service
			res, err := s.service.GetTodo(ctx, request)
			if err != nil {
				return nil, err
			}

			// Build the response protobuf and return it
			return buildTodo(res.Todo)
		},
		Description: "Retrieve a Todo object",
	}
}

func buildGetTodoRequestFromArgs(args map[string]interface{}) (*todoV1.GetTodoRequest, error) {
	if value, ok := args[argID]; ok {
		if val, ok := value.(string); ok {
			return &todoV1.GetTodoRequest{
				Id: val,
			}, nil
		} else {
			return nil, fmt.Errorf("'%s' not a string", argID)
		}
	} else {
		return nil, fmt.Errorf("must specify '%s'", argID)
	}
}
