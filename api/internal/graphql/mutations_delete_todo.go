package graphql

import (
	"fmt"

	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/graphql-go/graphql"
)

func (s Server) buildFieldForDeleteTodo(todoType *graphql.Object) *graphql.Field {
	deleteResponseType := graphql.NewObject(graphql.ObjectConfig{
		Name: "DeleteTodo",
		Fields: graphql.Fields{
			"success": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "Always true.",
			},
		},
	})
	return &graphql.Field{
		Type:        deleteResponseType,
		Description: "Delete Todo",
		Args: graphql.FieldConfigArgument{
			argID: &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
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
				return nil, err
			}

			// Call the service
			_, err = s.service.DeleteTodo(ctx, request)
			if err != nil {
				return nil, err
			}

			type deleteResponse struct {
				Success bool `json:"success"`
			}

			// Build the response protobuf and return it
			return deleteResponse{Success: true}, nil
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
