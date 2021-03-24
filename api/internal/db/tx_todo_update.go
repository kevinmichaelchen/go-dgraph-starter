package db

import (
	"context"
	"time"

	"github.com/MyOrg/todo-api/internal/obs"
	todoV1 "github.com/MyOrg/todo-api/pkg/pb/myorg/todo/v1"
	"github.com/dgraph-io/dgo/v200/protos/api"
)

func pathMap(in []string) map[string]bool {
	out := make(map[string]bool)
	for _, e := range in {
		out[e] = true
	}
	return out
}

func (tx *todoTransactionImpl) UpdateTodo(ctx context.Context, request *todoV1.UpdateTodoRequest) (*todoV1.UpdateTodoResponse, error) {
	ctx, span := obs.NewSpan(ctx, "UpdateTodo")
	defer span.End()

	logger := obs.ToLogger(ctx)

	id := request.Id

	nowStr := time.Now().UTC().Format(time.RFC3339)

	query := `
		query getTodo($id: string) {
			todo as var(func: eq(id, $id)) {
				todo_id as id
				todo_created_at as created_at
				todo_title as title
				todo_is_done as is_done
				creator {
					todo_creator_id as id
				}
			}
		}
	`

	nquads := []*api.NQuad{
		// Insert event
		nquadStr("_:todoEvent", fieldDgraphType, dgraphTypeTodoEvent),
		nquadStr("_:todoEvent", fieldEventType, eventTypeUpdate),
		nquadStr("_:todoEvent", fieldEventAt, nowStr),
		nquadBool("_:todoEvent", fieldEventPublishedToSearchIndex, false),
		nquadRel("_:todoEvent", fieldTodoID, "val(todo_id)"),
		nquadRel("_:todoEvent", fieldTitle, "val(todo_title)"),
		nquadRel("_:todoEvent", fieldCreatedAt, "val(todo_created_at)"),
		nquadRel("_:todoEvent", fieldDone, "val(todo_is_done)"),
		nquadRel("_:todoEvent", fieldCreatorID, "val(todo_creator_id)"),
	}

	// Check our FieldMask to see which fields we want to update
	fields := pathMap(request.FieldMask.Paths)
	if _, ok := fields["title"]; ok {
		nquads = append(nquads, nquadStr("uid(todo)", fieldTitle, request.Title))
	}
	if _, ok := fields["done"]; ok {
		nquads = append(nquads, nquadBool("uid(todo)", fieldDone, request.Done))
	}

	if len(nquads) > 0 {
		req := &api.Request{
			Query: query,
			Vars:  map[string]string{"$id": id},
			Mutations: []*api.Mutation{
				{
					// Only mutate if we find one entity
					Cond: `@if(eq(len(todo), 1))`,
					Set:  nquads,
				},
			},
		}

		// Perform query
		res, err := tx.tx.Do(ctx, req)
		if err != nil {
			return nil, err
		}

		// Log latency
		logger.Info().Msgf("Updated Todo in %s", latency(res))

		logger.Info().Msgf("Update response: %s", res.Json)
	}

	var todoPB *todoV1.Todo
	if res, err := tx.GetTodoByID(ctx, id); err != nil {
		return nil, err
	} else {
		todoPB = res
	}

	return &todoV1.UpdateTodoResponse{
		Todo: todoPB,
	}, nil
}
