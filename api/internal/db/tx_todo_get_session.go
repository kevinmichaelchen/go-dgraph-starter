package db

import (
	"context"

	"github.com/MyOrg/go-dgraph-starter/internal/models"
	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

func (tx *todoTransactionImpl) GetTodoByID(ctx context.Context, id string) (*models.Todo, error) {
	ctx, span := obs.NewSpan(ctx, "GetTodoByID")
	defer span.End()

	// Check cache before checking database
	redisKey := redisKeyForTodo(id)
	var cachedItem models.Todo
	if err := tx.redisClient.Get(ctx, redisKey, &cachedItem); err != nil {
		// Tag span with error
		obs.SetError(span, err)

		// Cache miss
		if err == redis.Nil {
			log.Warn().Msgf("CACHE MISS: Todo '%s' was not cached.", id)
		} else {
			// Internal error
			return nil, err
		}
	} else {
		log.Trace().Msgf("CACHE HIT: Todo '%s' was cached!", id)
		return &cachedItem, nil
	}

	// Find item by ID
	// Cache it
	return nil, nil
}
