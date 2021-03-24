package service

import (
	"context"
	"fmt"
	"github.com/MyOrg/todo-api/internal/db"
	"github.com/MyOrg/todo-api/internal/obs"
	todoV1 "github.com/MyOrg/todo-api/pkg/pb/myorg/todo/v1"
	"github.com/golang/protobuf/ptypes"
	"github.com/rs/xid"
)

func (s Service) CreateTodo(ctx context.Context, request *todoV1.CreateTodoRequest) (*todoV1.CreateTodoResponse, error) {
	logger := obs.ToLogger(ctx)

	requesterID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}

	// TODO add validation

	todo := &todoV1.Todo{
		Id:        xid.New().String(),
		CreatedAt: ptypes.TimestampNow(),
		Title:     request.Title,
		Done:      false,
		CreatorId: requesterID,
	}

	if err := s.dbClient.RunInTransaction(ctx, func(ctx context.Context, tx db.Transaction) error {
		return tx.CreateTodo(ctx, todo)
	}); err != nil {
		return nil, err
	}

	// TODO use Transactional Outbox pattern instead
	if err := s.searchClient.AddOrUpdate(ctx, todo); err != nil {
		return nil, err
	}

	// Request information about the user
	logger.Info().Msgf("Fetching info for user: %s", requesterID)
	//userClient := userV1.NewUserServiceClient(s.usersConn)
	//if out, err := userClient.GetUser(ctx, &userV1.GetUserRequest{
	//	Id: requesterID,
	//}); err != nil {
	//	return nil, fmt.Errorf("failed to obtain user info: %w", err)
	//} else {
	//	logger.Info().Msgf("Fetched info for user: %s: %s", requesterID, out.User.Name)
	//}

	return &todoV1.CreateTodoResponse{
		Todo: todo,
	}, nil
}

func (s Service) UpdateTodo(ctx context.Context, request *todoV1.UpdateTodoRequest) (*todoV1.UpdateTodoResponse, error) {
	var response *todoV1.UpdateTodoResponse

	// TODO add validation

	if err := s.dbClient.RunInTransaction(ctx, func(ctx context.Context, tx db.Transaction) error {
		if res, err := tx.UpdateTodo(ctx, request); err != nil {
			return err
		} else {
			response = res
		}
		return nil
	}); err != nil {
		return nil, err
	}

	// TODO use Transactional Outbox pattern instead
	todoPB := &todoV1.Todo{
		Id:    request.Id,
		Title: request.Title,
	}
	if err := s.searchClient.AddOrUpdate(ctx, todoPB); err != nil {
		return nil, err
	}

	return response, nil
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
	// Validate inputs
	if request.OrderBy == todoV1.OrderTodosBy_ORDER_TODOS_BY_UNSPECIFIED {
		// Old to new
		request.OrderBy = todoV1.OrderTodosBy_ORDER_TODOS_BY_CREATED_AT_ASC
	}
	if f := request.PaginationRequest.GetForwardPaginationInfo(); f != nil {
		if f.First < db.MinPageSize || f.First > db.MaxPageSize {
			f.First = db.DefaultPageSize
		}
	} else if b := request.PaginationRequest.GetBackwardPaginationInfo(); b != nil {
		if b.Last < db.MinPageSize || b.Last > db.MaxPageSize {
			b.Last = db.DefaultPageSize
		}
	} else {
		request.PaginationRequest = &todoV1.PaginationRequest{
			Request: &todoV1.PaginationRequest_ForwardPaginationInfo{
				ForwardPaginationInfo: &todoV1.ForwardPaginationRequest{
					First: db.DefaultPageSize,
				},
			},
		}
	}

	var response *todoV1.GetTodosResponse
	if err := s.dbClient.RunInReadOnlyTransaction(ctx, func(ctx context.Context, tx db.Transaction) error {
		if res, err := tx.GetTodos(ctx, request); err != nil {
			return err
		} else {
			response = res
		}

		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s Service) DeleteTodo(ctx context.Context, request *todoV1.DeleteTodoRequest) (*todoV1.DeleteTodoResponse, error) {
	var response *todoV1.DeleteTodoResponse
	if err := s.dbClient.RunInTransaction(ctx, func(ctx context.Context, tx db.Transaction) error {
		if res, err := tx.DeleteTodo(ctx, request.Id); err != nil {
			return err
		} else {
			response = res
		}
		return nil
	}); err != nil {
		return nil, err
	}

	// TODO use Transactional Outbox pattern instead
	// Delete document from search index
	if err := s.searchClient.Delete(ctx, request.Id); err != nil {
		return nil, err
	}

	return response, nil
}

func (s Service) SearchTodos(ctx context.Context, request *todoV1.SearchTodosRequest) (*todoV1.SearchTodosResponse, error) {
	// Perform search query
	ids, err := s.searchClient.Query(ctx, request.Query)
	if err != nil {
		return nil, err
	}

	// TODO implement batch-lookup so we can minimize DB trips
	// Look up entities from database
	var todos []*todoV1.Todo
	for _, id := range ids {
		if res, err := s.GetTodo(ctx, &todoV1.GetTodoRequest{Id: string(id)}); err != nil {
			return nil, fmt.Errorf("failed to find Todo by id '%s': %w", string(id), err)
		} else {
			todos = append(todos, res.Todo)
		}
	}

	return &todoV1.SearchTodosResponse{
		Todos: todos,
	}, nil
}
