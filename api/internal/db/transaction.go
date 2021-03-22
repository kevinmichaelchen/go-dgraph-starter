package db

import (
	"fmt"
	"time"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"

	"github.com/MyOrg/go-dgraph-starter/internal/configuration"
)

const (
	eventTypeCreate = "create"
	eventTypeUpdate = "update"
	eventTypeDelete = "delete"

	dgraphTypeUser      = "User"
	dgraphTypeTodo      = "Todo"
	dgraphTypeTodoEvent = "TodoEvent"

	fieldDgraphType = "dgraph.type"

	fieldID        = "id"
	fieldCreatedAt = "created_at"
	fieldCreator   = "creator"

	fieldEventAt                     = "event_at"
	fieldEventType                   = "event_type"
	fieldTodoID                      = "todo_id"
	fieldEventPublishedToSearchIndex = "is_published_to_search_index"

	fieldName      = "name"
	fieldTitle     = "title"
	fieldDone      = "is_done"
	fieldCreatorID = "creator_id"
)

type Transaction interface {
	TodoTransaction
}

type txImpl struct {
	tx *dgo.Txn
	todoTransactionImpl
}

func newTransaction(tx *dgo.Txn, redisClient RedisClient, config configuration.Config) Transaction {
	return &txImpl{
		tx:                  tx,
		todoTransactionImpl: todoTransactionImpl{tx: tx, redisClient: redisClient},
	}
}

func latency(res *api.Response) string {
	elapsedMilliseconds := float64(res.Latency.TotalNs) / float64(time.Millisecond)
	elapsedMicroseconds := int64(res.Latency.TotalNs) / int64(time.Microsecond)
	return fmt.Sprintf("%.2f ms === %d μs === %d ns", elapsedMilliseconds, elapsedMicroseconds, res.Latency.TotalNs)
}
