package graphql

import (
	"fmt"

	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/graphql-go/graphql"
)

func (s Server) buildFieldForDeleteTodo(todoType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.Boolean,
		Description: "Delete a Todo",
		Args: graphql.FieldConfigArgument{
			argID: &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The ID of the Todo the client wishes to delete",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ctx := p.Context
			logger := obs.ToLogger(ctx)

			args := p.Args
			logger.Info().Msgf("Received GraphQL request to %s with args: %v", p.Info.FieldName, args)

			// Build the request protobuf from the GraphQL args
			request, err := buildDeleteTodoRequestFromArgs(args)
			if err != nil {
				return false, err
			}

			// Call the service
			_, err = s.service.DeleteTodo(ctx, request)
			if err != nil {
				return false, err
			}

			// Build the response protobuf and return it
			return true, nil
		},
	}
}

func buildDeleteTodoRequestFromArgs(args map[string]interface{}) (*todoV1.DeleteTodoRequest, error) {
	if value, ok := args[argID]; ok {
		if val, ok := value.(string); ok {
			return &todoV1.DeleteTodoRequest{
				Id: val,
			}, nil
		} else {
			return nil, fmt.Errorf("'%s' not a string", argID)
		}
	} else {
		return nil, fmt.Errorf("must specify '%s'", argID)
	}
}
