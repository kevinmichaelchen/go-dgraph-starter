package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/dgraph-io/dgo/v200/protos/api"
)

const (
	eventTypeCreate = "create"
	eventTypeUpdate = "update"
	eventTypeDelete = "delete"
)

type Transaction interface {
	UserTransaction
}

type txImpl struct {
	tx *sql.Tx
	userTransactionImpl
}

func newTransaction(tx *sql.Tx, redisClient RedisClient) Transaction {
	return &txImpl{
		tx:                  tx,
		userTransactionImpl: userTransactionImpl{tx: tx, redisClient: redisClient},
	}
}

func latency(res *api.Response) string {
	elapsedMilliseconds := float64(res.Latency.TotalNs) / float64(time.Millisecond)
	elapsedMicroseconds := int64(res.Latency.TotalNs) / int64(time.Microsecond)
	return fmt.Sprintf("%.2f ms === %d μs === %d ns", elapsedMilliseconds, elapsedMicroseconds, res.Latency.TotalNs)
}
