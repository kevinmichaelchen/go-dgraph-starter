package graphql

import (
	"fmt"

	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/graphql-go/graphql"
)

func (s Server) buildFieldForUpdateTodo(todoType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        todoType,
		Description: "Update Todo",
		Args: graphql.FieldConfigArgument{
			argID: &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			argTitle: &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			argDone: &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ctx := p.Context
			logger := obs.ToLogger(ctx)

			args := p.Args
			logger.Info().Msgf("Received GraphQL request to %s with args: %v", p.Info.FieldName, args)

			// TODO get selection set and use FieldMask

			// Build the request protobuf from the GraphQL args
			request, err := buildUpdateTodoRequestFromArgs(args)
			if err != nil {
				return nil, err
			}

			// Call the service
			res, err := s.service.UpdateTodo(ctx, request)
			if err != nil {
				return nil, err
			}

			// Build the response protobuf and return it
			return buildTodo(res.Todo)
		},
	}
}

func buildUpdateTodoRequestFromArgs(args map[string]interface{}) (*todoV1.UpdateTodoRequest, error) {
	request := &todoV1.UpdateTodoRequest{}

	if value, ok := args[argTitle]; ok {
		if val, ok := value.(string); ok {
			request.Title = val
		} else {
			return nil, fmt.Errorf("'%s' not a string", argTitle)
		}
	} else {
		return nil, fmt.Errorf("must specify '%s'", argTitle)
	}

	if value, ok := args[argDone]; ok {
		if val, ok := value.(string); ok {
			request.Title = val
		} else {
			return nil, fmt.Errorf("'%s' not a string", argDone)
		}
	} else {
		return nil, fmt.Errorf("must specify '%s'", argDone)
	}

	return request, nil
}
