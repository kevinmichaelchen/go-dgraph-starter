package graphql

import (
	"fmt"

	"github.com/MyOrg/todo-api/internal/obs"
	todoV1 "github.com/MyOrg/todo-api/pkg/pb/myorg/todo/v1"
	"github.com/graphql-go/graphql"
)

type SearchResponse struct {
	Todos []Todo `json:"todos"`
}

func (s Server) buildFieldForSearchTodos(todoType *graphql.Object) *graphql.Field {
	responseType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "SearchResponse",
		Description: "A response containing a list of Todos that match the search query",
		Fields: graphql.Fields{
			"todos": &graphql.Field{
				Type:        graphql.NewList(todoType),
				Description: "List of Todos matching search query",
			},
		},
	})
	return &graphql.Field{
		Type:        responseType,
		Description: "Retrieve a page of Todo objects that match the search query",
		Args: graphql.FieldConfigArgument{
			argQuery: &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Used for searching",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ctx := p.Context
			logger := obs.ToLogger(ctx)

			args := p.Args
			logger.Info().Msgf("Received GraphQL request to %s with args: %v", p.Info.FieldName, args)

			// Build the request protobuf from the GraphQL args
			request, err := buildSearchTodosRequestFromArgs(args)
			if err != nil {
				return nil, err
			}

			// Call the service
			res, err := s.service.SearchTodos(ctx, request)
			if err != nil {
				return nil, err
			}

			// Build the response protobuf and return it
			return buildResponseForSearchTodos(res)
		},
	}
}

func buildSearchTodosRequestFromArgs(args map[string]interface{}) (*todoV1.SearchTodosRequest, error) {
	out := &todoV1.SearchTodosRequest{}

	if value, ok := args[argQuery]; ok {
		if val, ok := value.(string); ok {
			out.Query = val
		} else {
			return nil, fmt.Errorf("'%s' not a string: %T", argQuery, value)
		}
	}

	return out, nil
}

func buildResponseForSearchTodos(in *todoV1.SearchTodosResponse) (SearchResponse, error) {
	var todos []Todo
	for _, todoPB := range in.Todos {
		todo, err := buildTodo(todoPB)
		if err != nil {
			return SearchResponse{}, err
		}
		todos = append(todos, todo)
	}
	return SearchResponse{
		Todos: todos,
	}, nil
}
