package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/MyOrg/go-dgraph-starter/internal/configuration"
	"github.com/MyOrg/go-dgraph-starter/internal/obs"
)

type TransactionFunc func(context.Context, Transaction) error

type Client interface {
	Ping(ctx context.Context) error
	RunInTransaction(ctx context.Context, fn TransactionFunc, opts *sql.TxOptions) error
}

type clientImpl struct {
	db          *sql.DB
	config      configuration.Config
	redisClient RedisClient
}

func NewClient(db *sql.DB, redisClient RedisClient, config configuration.Config) Client {
	return &clientImpl{
		db:          db,
		config:      config,
		redisClient: redisClient,
	}
}

func (c *clientImpl) Ping(ctx context.Context) error {
	return c.db.PingContext(ctx)
}

func (c *clientImpl) EnsureConnectionIsOpen(ctx context.Context) {
	// TODO if the connection is closed, re-open it
}

func (c *clientImpl) RunInTransaction(ctx context.Context, fn TransactionFunc, opts *sql.TxOptions) error {
	// Create a new tracing span
	ctx, span := obs.NewSpan(ctx, "RunInTransaction")
	defer span.End()

	logger := obs.ToLogger(ctx)

	c.EnsureConnectionIsOpen(ctx)

	// Run function
	tx, err := c.db.BeginTx(ctx, opts)
	if err != nil {
		return err
	}

	if err := fn(ctx, newTransaction(tx, c.redisClient, c.config)); err != nil {
		if err := tx.Rollback(); err != nil {
			logger.Warn().Err(err).Msg("rollback failed")
			// We don't really need to return the rollback error if one occurs;
			// the client is more interested in the underlying database layer error.
		}
		return err
	} else {
		if opts == nil || !opts.ReadOnly {
			if err := tx.Commit(); err != nil {
				return fmt.Errorf("transaction failed to commit: %w", err)
			}
		}
	}

	return nil
}
