package db

import (
	"context"
	"encoding/json"
	"fmt"
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

	now, err := ptypes.Timestamp(item.CreatedAt)
	if err != nil {
		return err
	}
	nowStr := now.Format(time.RFC3339)

	logger := obs.ToLogger(ctx)

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

	if res, err := tx.tx.QueryWithVars(ctx, query, map[string]string{"$creatorID": item.AuthorId}); err != nil {
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
					nquadStr("_:newUser", "dgraph.type", "User"),
					nquadStr("_:newUser", "id", item.AuthorId),
					nquadStr("_:newUser", "name", "Alice"),
					nquadStr("_:newUser", "created_at", nowStr),
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
			nquadStr("_:todo", "dgraph.type", "Todo"),
			nquadStr("_:todo", "id", item.Id),
			nquadStr("_:todo", "title", item.Title),
			nquadStr("_:todo", "created_at", nowStr),
			nquadBool("_:todo", "is_done", item.Done),
			nquadRel("_:todo", "creator", creatorUID),
		},
	})

	if err != nil {
		return err
	}

	// TODO Insert into cache

	logger.Info().Msgf("Created new Todo in %s", latency(res))

	return nil
}

func latency(res *api.Response) string {
	elapsedMilliseconds := float64(res.Latency.TotalNs) / float64(time.Millisecond)
	elapsedMicroseconds := float64(res.Latency.TotalNs) / float64(time.Millisecond)
	return fmt.Sprintf("%f ms = %f μs = %d ns", elapsedMilliseconds, elapsedMicroseconds, res.Latency.TotalNs)
}
