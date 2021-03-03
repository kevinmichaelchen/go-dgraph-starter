package db

import (
	"fmt"
	"time"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"

	"github.com/MyOrg/go-dgraph-starter/internal/configuration"
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
