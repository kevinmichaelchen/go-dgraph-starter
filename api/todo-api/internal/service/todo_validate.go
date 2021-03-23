package service

import (
	"context"

	todoV1 "github.com/MyOrg/todo-api/pkg/pb/myorg/todo/v1"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func validateUpdateUserRequest(ctx context.Context, r *todoV1.UpdateTodoRequest) error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Id, validation.Required),
	)
}
