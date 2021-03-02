package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/golang/protobuf/ptypes"
)

func (s Service) CreateTodo(ctx context.Context, request *todoV1.CreateTodoRequest) (*todoV1.CreateTodoResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func (s Service) GetTodo(ctx context.Context, request *todoV1.GetTodoRequest) (*todoV1.GetTodoResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func (s Service) GetTodos(ctx context.Context, request *todoV1.GetTodosRequest) (*todoV1.GetTodosResponse, error) {
	return &todoV1.GetTodosResponse{
		Edges: []*todoV1.TodoEdge{
			{
				Cursor: "1",
				Todo: &todoV1.Todo{
					Id:        "1",
					Title:     "Todo 1",
					CreatedAt: ptypes.TimestampNow(),
				},
			},
			{
				Cursor: "2",
				Todo: &todoV1.Todo{
					Id:        "2",
					Title:     "Todo 2",
					CreatedAt: ptypes.TimestampNow(),
				},
			},
		},
		PageInfo: &todoV1.PageInfo{
			EndCursor: "2",
		},
		TotalCount: 2,
	}, nil
}

func (s Service) UpdateTodo(ctx context.Context, request *todoV1.UpdateTodoRequest) (*todoV1.UpdateTodoResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func (s Service) DeleteTodo(ctx context.Context, request *todoV1.DeleteTodoRequest) (*todoV1.DeleteTodoResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}
