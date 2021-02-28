package db

import (
	"context"
	"fmt"
	"github.com/dgraph-io/dgo/v200"

	"github.com/MyOrg/go-dgraph-starter/internal/configuration"
	"github.com/MyOrg/go-dgraph-starter/internal/obs"
)

type TransactionFunc func(context.Context, Transaction) error

type Client interface {
	Ping(ctx context.Context) error
	RunInTransaction(ctx context.Context, fn TransactionFunc) error
	RunInReadOnlyTransaction(ctx context.Context, fn TransactionFunc) error
}

type clientImpl struct {
	db          *dgo.Dgraph
	config      configuration.Config
	redisClient RedisClient
}

func NewClient(db *dgo.Dgraph, redisClient RedisClient, config configuration.Config) Client {
	return &clientImpl{
		db:          db,
		config:      config,
		redisClient: redisClient,
	}
}

func (c *clientImpl) Ping(ctx context.Context) error {
	// TODO how do you ping dgraph?
	return nil
}

func (c *clientImpl) RunInReadOnlyTransaction(ctx context.Context, fn TransactionFunc) error {
	// Create a new tracing span
	ctx, span := obs.NewSpan(ctx, "RunInReadOnlyTransaction")
	defer span.End()

	// Create logger from context
	logger := obs.ToLogger(ctx)

	// Create read-only transaction
	txn := c.db.NewReadOnlyTxn()
	defer txn.Discard(ctx)

	// Run function
	if err := fn(ctx, newTransaction(txn, c.redisClient, c.config)); err != nil {
		if rollbackErr := txn.Discard(ctx); rollbackErr != nil {
			logger.Warn().Err(rollbackErr).Msg("transaction rollback failed")
			// We don't really need to return the rollback error if one occurs;
			// the client is more interested in the underlying database layer error.
		}
		return err
	}

	return nil
}

func (c *clientImpl) RunInTransaction(ctx context.Context, fn TransactionFunc) error {
	// Create a new tracing span
	ctx, span := obs.NewSpan(ctx, "RunInTransaction")
	defer span.End()

	// Create logger from context
	logger := obs.ToLogger(ctx)

	// Create transaction
	txn := c.db.NewTxn()
	defer txn.Discard(ctx)

	// Run function
	if err := fn(ctx, newTransaction(txn, c.redisClient, c.config)); err != nil {
		if rollbackErr := txn.Discard(ctx); rollbackErr != nil {
			logger.Warn().Err(rollbackErr).Msg("transaction rollback failed")
			// We don't really need to return the rollback error if one occurs;
			// the client is more interested in the underlying database layer error.
		}
		return err
	}

	// Commit the transaction
	if err := txn.Commit(ctx); err != nil {
		return fmt.Errorf("transaction failed to commit: %w", err)
	}

	return nil
}
