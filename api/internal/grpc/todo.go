package grpc

import (
	"context"

	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
)

func (s Server) CreateTodo(ctx context.Context, request *todoV1.CreateTodoRequest) (*todoV1.CreateTodoResponse, error) {
	return s.service.CreateTodo(ctx, request)
}

func (s Server) GetTodo(ctx context.Context, request *todoV1.GetTodoRequest) (*todoV1.GetTodoResponse, error) {
	return s.service.GetTodo(ctx, request)
}

func (s Server) GetTodos(ctx context.Context, request *todoV1.GetTodosRequest) (*todoV1.GetTodosResponse, error) {
	return s.service.GetTodos(ctx, request)
}

func (s Server) SearchTodos(ctx context.Context, request *todoV1.SearchTodosRequest) (*todoV1.SearchTodosResponse, error) {
	return s.service.SearchTodos(ctx, request)
}

func (s Server) UpdateTodo(ctx context.Context, request *todoV1.UpdateTodoRequest) (*todoV1.UpdateTodoResponse, error) {
	return s.service.UpdateTodo(ctx, request)
}

func (s Server) DeleteTodo(ctx context.Context, request *todoV1.DeleteTodoRequest) (*todoV1.DeleteTodoResponse, error) {
	return s.service.DeleteTodo(ctx, request)
}
