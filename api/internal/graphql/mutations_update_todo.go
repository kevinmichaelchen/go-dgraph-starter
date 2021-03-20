package graphql

import (
	"fmt"

	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/graphql-go/graphql"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (s Server) buildFieldForUpdateTodo(todoType *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Type:        todoType,
		Description: "Update a Todo",
		Args: graphql.FieldConfigArgument{
			argID: &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The ID of the Todo the client wishes to update",
			},
			argTitle: &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "The Todo's new title. This argument is optional.",
			},
			argDone: &graphql.ArgumentConfig{
				Type:        graphql.Boolean,
				Description: "Whether the Todo is being marked as completed. This argument is optional.",
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
	var paths []string
	request := &todoV1.UpdateTodoRequest{}

	if value, ok := args[argID]; ok {
		// TODO do we need these type assertions or does the graphql library take care of that for us?
		if val, ok := value.(string); ok {
			request.Id = val
		} else {
			return nil, fmt.Errorf("'%s' not a string", argID)
		}
	} else {
		return nil, fmt.Errorf("must specify '%s'", argID)
	}

	if value, ok := args[argTitle]; ok {
		if val, ok := value.(string); ok {
			request.Title = val
			paths = append(paths, argTitle)
		} else {
			return nil, fmt.Errorf("'%s' not a string", argTitle)
		}
	} else {
		return nil, fmt.Errorf("must specify '%s'", argTitle)
	}

	if value, ok := args[argDone]; ok {
		if val, ok := value.(bool); ok {
			request.Done = val
			paths = append(paths, argDone)
		} else {
			return nil, fmt.Errorf("'%s' not a string", argDone)
		}
	} else {
		return nil, fmt.Errorf("must specify '%s'", argDone)
	}

	request.FieldMask = &fieldmaskpb.FieldMask{
		Paths: paths,
	}

	return request, nil
}
