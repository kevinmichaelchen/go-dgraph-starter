package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/MyOrg/go-dgraph-starter/internal/db"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/golang/protobuf/ptypes"
	"github.com/rs/xid"
)

func (s Service) CreateTodo(ctx context.Context, request *todoV1.CreateTodoRequest) (*todoV1.CreateTodoResponse, error) {
	requesterID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	todo := &todoV1.Todo{
		Id:        xid.New().String(),
		CreatedAt: ptypes.TimestampNow(),
		Title:     request.Title,
		Done:      false,
		AuthorId:  requesterID,
	}
	err = s.dbClient.RunInTransaction(ctx, func(ctx context.Context, tx db.Transaction) error {
		return tx.CreateTodo(ctx, todo)
	})

	if err != nil {
		return nil, err
	}

	return &todoV1.CreateTodoResponse{
		Todo: todo,
	}, nil
}

func (s Service) GetTodo(ctx context.Context, request *todoV1.GetTodoRequest) (*todoV1.GetTodoResponse, error) {
	// requesterID, err := getUserID(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	var todo *todoV1.Todo

	if err := s.dbClient.RunInReadOnlyTransaction(ctx, func(ctx context.Context, tx db.Transaction) error {
		if t, err := tx.GetTodoByID(ctx, request.Id); err != nil {
			return err
		} else {
			todo = t
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &todoV1.GetTodoResponse{
		Todo: todo,
	}, nil
}

func (s Service) GetTodos(ctx context.Context, request *todoV1.GetTodosRequest) (*todoV1.GetTodosResponse, error) {
	return &todoV1.GetTodosResponse{
		Edges: []*todoV1.TodoEdge{
			{
				Cursor: "1",
				Node: &todoV1.Todo{
					Id:        "1",
					Title:     "Todo 1",
					CreatedAt: ptypes.TimestampNow(),
				},
			},
			{
				Cursor: "2",
				Node: &todoV1.Todo{
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
