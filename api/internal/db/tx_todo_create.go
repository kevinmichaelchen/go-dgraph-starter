package db

import (
	"context"
	"encoding/json"
	"time"

	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/golang/protobuf/ptypes"
	"github.com/rs/zerolog/log"
)

func (tx *todoTransactionImpl) CreateTodo(ctx context.Context, item *todoV1.Todo) error {
	ctx, span := obs.NewSpan(ctx, "CreateTodo")
	defer span.End()

	logger := obs.ToLogger(ctx)

	now, err := ptypes.Timestamp(item.CreatedAt)
	if err != nil {
		return err
	}
	nowStr := now.Format(time.RFC3339)

	type User struct {
		Uid string `json:"uid,omitempty"`
	}
	type Creator struct {
		Users []User `json:"creator"`
	}

	var creatorUID string

	// Find the creator's UID
	query := `
		query getCreator($creatorID: string) {
			creator(func: eq(dgraph.type, "User")) @filter(eq(id, $creatorID)) {
				uid
			}
		}
	`

	if res, err := tx.tx.QueryWithVars(ctx, query, map[string]string{"$creatorID": item.CreatorId}); err != nil {
		return err
	} else {
		var c Creator
		if err := json.Unmarshal(res.Json, &c); err != nil {
			return err
		}
		if len(c.Users) == 0 {
			log.Info().Msg("No user existed... creating one...")
			// Insert user into database
			if res, err := tx.tx.Mutate(ctx, &api.Mutation{
				Set: []*api.NQuad{
					nquadStr("_:newUser", fieldDgraphType, dgraphTypeUser),
					nquadStr("_:newUser", fieldID, item.CreatorId),
					nquadStr("_:newUser", fieldName, "Alice"),
					nquadStr("_:newUser", fieldCreatedAt, nowStr),
				},
			}); err != nil {
				return err
			} else {
				creatorUID = res.Uids["newUser"]
				logger.Info().Msgf("Created new user with UID: %s in %s", creatorUID, latency(res))
			}
		} else {
			creatorUID = c.Users[0].Uid
			logger.Info().Msgf("Found existing user with UID: %s", creatorUID)
		}
	}

	// Insert into database
	res, err := tx.tx.Mutate(ctx, &api.Mutation{
		Set: []*api.NQuad{
			nquadStr("_:todo", fieldDgraphType, dgraphTypeTodo),
			nquadStr("_:todo", fieldID, item.Id),
			nquadStr("_:todo", fieldTitle, item.Title),
			nquadStr("_:todo", fieldCreatedAt, nowStr),
			nquadBool("_:todo", fieldDone, item.Done),
			nquadRel("_:todo", fieldCreator, creatorUID),

			nquadStr("_:todoEvent", fieldDgraphType, dgraphTypeTodo),
			nquadStr("_:todoEvent", fieldEventType, eventTypeCreate),
			nquadStr("_:todoEvent", fieldEventAt, nowStr),
			nquadBool("_:todoEvent", fieldEventPublishedToSearchIndex, false),
			nquadStr("_:todoEvent", fieldTodoID, item.Id),
			nquadStr("_:todoEvent", fieldTitle, item.Title),
			nquadStr("_:todoEvent", fieldCreatedAt, nowStr),
			nquadBool("_:todoEvent", fieldDone, item.Done),
			nquadRel("_:todoEvent", fieldCreator, creatorUID),
		},
	})

	// Handle error
	if err != nil {
		return err
	}

	// TODO Insert into cache

	logger.Info().Msgf("Created new Todo in %s", latency(res))

	return nil
}
