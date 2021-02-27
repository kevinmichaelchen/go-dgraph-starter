package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
)

func (s Service) CreateTodo(ctx context.Context, request *todoV1.CreateTodoRequest) (*todoV1.CreateTodoResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func (s Service) GetTodo(ctx context.Context, request *todoV1.GetTodoRequest) (*todoV1.GetTodoResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func (s Service) GetTodos(ctx context.Context, request *todoV1.GetTodosRequest) (*todoV1.GetTodosResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func (s Service) UpdateTodo(ctx context.Context, request *todoV1.UpdateTodoRequest) (*todoV1.UpdateTodoResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func (s Service) DeleteTodo(ctx context.Context, request *todoV1.DeleteTodoRequest) (*todoV1.DeleteTodoResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}
