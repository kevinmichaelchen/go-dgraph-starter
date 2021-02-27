package db

import (
	"database/sql"

	"github.com/MyOrg/go-dgraph-starter/internal/configuration"
)

const (
	SpanPrefix = "db."
)

type Transaction interface {
	TodoTransaction
	GetTx() *sql.Tx
}

type txImpl struct {
	tx *sql.Tx
	todoTransactionImpl
}

func (tx txImpl) GetTx() *sql.Tx {
	return tx.tx
}

func newTransaction(tx *sql.Tx, redisClient RedisClient, config configuration.Config) Transaction {
	return &txImpl{
		tx:                  tx,
		todoTransactionImpl: todoTransactionImpl{tx: tx, redisClient: redisClient},
	}
}
