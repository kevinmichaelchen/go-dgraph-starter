package db

import (
	"github.com/dgraph-io/dgo/v200"

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
