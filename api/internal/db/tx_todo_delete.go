package db

import (
	"context"
	"time"

	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/dgraph-io/dgo/v200/protos/api"
)

func (tx *todoTransactionImpl) DeleteTodo(ctx context.Context, id string) (*todoV1.DeleteTodoResponse, error) {
	ctx, span := obs.NewSpan(ctx, "DeleteTodo")
	defer span.End()

	logger := obs.ToLogger(ctx)

	nowStr := time.Now().UTC().Format(time.RFC3339)

	// Perform deletion
	res, err := tx.tx.Do(ctx, &api.Request{
		Query: `
		query getTodo($id: string) {
			todo as var(func: eq(id, $id)) {
				todo_id as id
				todo_title as title
				todo_is_done as is_done
				creator {
					todo_creator_id as id
				}
			}
		}
		`,
		Mutations: []*api.Mutation{
			{
				// Only delete if we find one Todo with that ID.
				// Otherwise, we'll assume there are no Todos with that ID,
				// and we don't perform the Deletion, thus ensuring Idempotence.
				Cond: `@if(eq(len(todo), 1))`,

				// Instruct Dgraph to delete the Todo
				Del: []*api.NQuad{
					// The pattern S * * deletes all the known edges out of a node,
					// any reverse edges corresponding to the removed edges,
					// and any indexing for the removed data.
					nquadAll("uid(todo)"),
				},

				// Insert event
				Set: []*api.NQuad{
					nquadStr("_:todoEvent", fieldDgraphType, dgraphTypeTodo),
					nquadStr("_:todoEvent", fieldEventType, eventTypeCreate),
					nquadStr("_:todoEvent", fieldEventAt, nowStr),
					nquadBool("_:todoEvent", fieldEventPublishedToSearchIndex, false),
					nquadRel("_:todoEvent", fieldTodoID, "val(todo_id)"),
					nquadRel("_:todoEvent", fieldTitle, "val(todo_title)"),
					nquadStr("_:todoEvent", fieldCreatedAt, nowStr),
					nquadRel("_:todoEvent", fieldDone, "val(todo_is_done)"),
					nquadRel("_:todoEvent", fieldCreatorID, "val(todo_creator_id)"),
				},
			},
		},
		Vars: map[string]string{
			"$id": id,
		},
	})

	// Handle error
	if err != nil {
		return nil, err
	}

	// TODO Delete from cache

	logger.Info().Msgf("Deleted Todo in %s", latency(res))

	return &todoV1.DeleteTodoResponse{}, nil
}
