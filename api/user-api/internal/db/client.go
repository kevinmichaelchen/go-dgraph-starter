package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/MyOrg/user-api/internal/obs"
)

type TransactionFunc func(context.Context, Transaction) error

type Client interface {
	Ping(ctx context.Context) error
	RunInTransaction(ctx context.Context, fn TransactionFunc) error
	RunInReadOnlyTransaction(ctx context.Context, fn TransactionFunc) error
}

type clientImpl struct {
	db          *sql.DB
	redisClient RedisClient
}

func NewClient(db *sql.DB, redisClient RedisClient) Client {
	return &clientImpl{
		db:          db,
		redisClient: redisClient,
	}
}

func (c *clientImpl) Ping(ctx context.Context) error {
	return c.db.Ping()
}

func (c *clientImpl) RunInReadOnlyTransaction(ctx context.Context, fn TransactionFunc) error {
	// Create a new tracing span
	ctx, span := obs.NewSpan(ctx, "RunInReadOnlyTransaction")
	defer span.End()

	// Create logger from context
	logger := obs.ToLogger(ctx)

	// Create read-only transaction
	txn, err := c.db.BeginTx(ctx, &sql.TxOptions{
		ReadOnly: true,
	})

	// Handle error
	if err != nil {
		return fmt.Errorf("failed to open read-only database transaction: %w", err)
	}
	defer txn.Rollback()

	// Run function
	if err := fn(ctx, newTransaction(txn, c.redisClient)); err != nil {
		if rollbackErr := txn.Rollback(); rollbackErr != nil {
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
	txn, err := c.db.BeginTx(ctx, nil)

	// Handle error
	if err != nil {
		return fmt.Errorf("failed to open database transaction: %w", err)
	}

	// Run function
	if err := fn(ctx, newTransaction(txn, c.redisClient)); err != nil {
		if rollbackErr := txn.Rollback(); rollbackErr != nil {
			logger.Warn().Err(rollbackErr).Msg("transaction rollback failed")
			// We don't really need to return the rollback error if one occurs;
			// the client is more interested in the underlying database layer error.
		}
		return err
	}

	// Commit the transaction
	if err := txn.Commit(); err != nil {
		return fmt.Errorf("transaction failed to commit: %w", err)
	}

	return nil
}
