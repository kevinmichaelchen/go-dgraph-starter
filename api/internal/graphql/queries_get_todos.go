package graphql

import (
	"errors"
	"fmt"
	"time"

	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"

	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	"github.com/graphql-go/graphql"
)

const (
	argWhere   = "where"
	argOrderBy = "orderBy"
	argFirst   = "first"
	argAfter   = "after"
	argLast    = "last"
	argBefore  = "before"
)

func (s Server) buildFieldForGetTodos(todoType *graphql.Object) *graphql.Field {
	typeEdge := graphql.NewObject(graphql.ObjectConfig{
		Name: "TodosEdge",
		Fields: graphql.Fields{
			// TODO add fields
		},
	})
	typePageInfo := graphql.NewObject(graphql.ObjectConfig{
		Name: "TodosPageInfo",
		Fields: graphql.Fields{
			// TODO add fields
		},
	})
	typeTodosPage := graphql.NewObject(graphql.ObjectConfig{
		Name: "TodosPage",
		Fields: graphql.Fields{
			"totalCount": &graphql.Field{
				Type: graphql.Int,
				Description: "Total number of items in database matching filter",
			},
			"edges": &graphql.Field{
				Type: graphql.NewList(typeEdge),
				Description: "The items in the page",
			},
			"pageInfo": &graphql.Field{
				Type: typePageInfo,
				Description: "Information about the page",
			},
		},
	})

	graphql.NewList(todoType)

	return &graphql.Field{
		Type: typeTodosPage,
		Args: buildArgsForGetTodos(),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ctx := p.Context
			logger := obs.ToLogger(ctx)
			logger.Info().Msg("Resolving GraphQL field: hello")

			args := p.Args

			logger.Info().Msgf("args %v", args)

			// Build the request protobuf from the GraphQL args
			request, err := buildGetTodosRequestFromArgs(args)
			if err != nil {
				return nil, err
			}

			// Call the service
			res, err := s.service.GetTodos(ctx, request)
			if err != nil {
				return nil, err
			}

			// Build the response protobuf and return it
			return buildResponseForGetTodos(res)
		},
		Description: "Retrieve a page of Todo objects",
	}
}

func buildResponseForGetTodos(in *todoV1.GetTodosResponse) ([]Todo, error) {
	// totalCount := in.TotalCount
	// TODO build response
	return []Todo{
		{"1", time.Now(), "Title 1", false},
		{"2", time.Now(), "Title 2", false},
	}, nil
}

func buildArgsForGetTodos() graphql.FieldConfigArgument {
	orderByEnum := buildOrderByEnum()

	whereType := graphql.NewObject(graphql.ObjectConfig{
		Name: "TodosWhere",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.DateTime,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"done": &graphql.Field{
				Type: graphql.Boolean,
			},
			// TODO id_not, id_in, id_not_in, etc. see GraphCMS
		},
	})

	return graphql.FieldConfigArgument{
		argWhere: &graphql.ArgumentConfig{
			Type:        whereType,
			Description: "Used for filtering",
		},
		argOrderBy: &graphql.ArgumentConfig{
			Type:        orderByEnum,
			Description: "Used for ordering/sorting items",
		},
		argFirst: &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "Set a page size limit for forward pagination",
		},
		argAfter: &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "A cursor to be used for forward pagination",
		},
		argLast: &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "Set a page size limit for backward pagination",
		},
		argBefore: &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "A cursor to be used for backward pagination",
		},
	}
}

func buildGetTodosRequestFromArgs(args map[string]interface{}) (*todoV1.GetTodosRequest, error) {
	request := &todoV1.GetTodosRequest{
		PaginationRequest: &todoV1.PaginationRequest{},
		OrderBy:           todoV1.OrderTodosBy_ORDER_TODOS_BY_ID_ASC,
		Where:             &todoV1.TodosWhere{},
	}

	// TODO parse WHERE args

	if value, ok := args[argOrderBy]; ok {
		if val, ok := value.(int32); ok {
			request.OrderBy = todoV1.OrderTodosBy(val)
			if request.OrderBy == todoV1.OrderTodosBy_ORDER_TODOS_BY_UNSPECIFIED {
				request.OrderBy = todoV1.OrderTodosBy_ORDER_TODOS_BY_ID_ASC
			}
		} else {
			return nil, fmt.Errorf("'%s' not an integer", argOrderBy)
		}
	}

	if value, ok := args[argFirst]; ok {
		if val, ok := value.(int); ok {
			if r := request.PaginationRequest.GetForwardPaginationInfo(); r != nil {
				r.First = int32(val)
			} else {
				request.PaginationRequest.Request = &todoV1.PaginationRequest_ForwardPaginationInfo{
					ForwardPaginationInfo: &todoV1.ForwardPaginationRequest{
						First: int32(val),
					},
				}
			}
		} else {
			return nil, errors.New("'first' not an integer")
		}
	}

	if value, ok := args[argAfter]; ok {
		if val, ok := value.(string); ok {
			if r := request.PaginationRequest.GetForwardPaginationInfo(); r != nil {
				r.After = val
			} else {
				request.PaginationRequest.Request = &todoV1.PaginationRequest_ForwardPaginationInfo{
					ForwardPaginationInfo: &todoV1.ForwardPaginationRequest{
						After: val,
					},
				}
			}
		} else {
			return nil, errors.New("'first' not a string")
		}
	}

	if value, ok := args[argLast]; ok {
		if val, ok := value.(int); ok {
			if r := request.PaginationRequest.GetBackwardPaginationInfo(); r != nil {
				r.Last = int32(val)
			} else {
				request.PaginationRequest.Request = &todoV1.PaginationRequest_BackwardPaginationInfo{
					BackwardPaginationInfo: &todoV1.BackwardPaginationRequest{
						Last: int32(val),
					},
				}
			}
		} else {
			return nil, errors.New("'last' not an integer")
		}
	}

	if value, ok := args[argBefore]; ok {
		if val, ok := value.(string); ok {
			if r := request.PaginationRequest.GetBackwardPaginationInfo(); r != nil {
				r.Before = val
			} else {
				request.PaginationRequest.Request = &todoV1.PaginationRequest_BackwardPaginationInfo{
					BackwardPaginationInfo: &todoV1.BackwardPaginationRequest{
						Before: val,
					},
				}
			}
		} else {
			return nil, errors.New("'first' not a string")
		}
	}

	return request, nil
}

func buildOrderByEnum() *graphql.Enum {
	return graphql.NewEnum(graphql.EnumConfig{
		Name:        "OrderBy",
		Description: "Order Todo Objects By",
		Values: graphql.EnumValueConfigMap{
			"ID_ASC": &graphql.EnumValueConfig{
				Value:       1,
				Description: "Order by ID ascending",
			},
			"ID_DESC": &graphql.EnumValueConfig{
				Value:       2,
				Description: "Order by ID descending",
			},
			"CREATED_AT_ASC": &graphql.EnumValueConfig{
				Value:       3,
				Description: "Order by createdAt ascending",
			},
			"CREATED_AT_DESC": &graphql.EnumValueConfig{
				Value:       4,
				Description: "Order by createdAt descending",
			},
			"TITLE_ASC": &graphql.EnumValueConfig{
				Value:       5,
				Description: "Order by title ascending",
			},
			"TITLE_DESC": &graphql.EnumValueConfig{
				Value:       6,
				Description: "Order by title descending",
			},
		},
	})
}
