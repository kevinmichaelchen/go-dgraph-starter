package db

import (
	"context"
	"time"

	"github.com/MyOrg/go-dgraph-starter/internal/configuration"
	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type RedisClient interface {
	Get(ctx context.Context, key string, value interface{}) error
	Set(ctx context.Context, key string, value interface{}, longevity time.Duration) error
}

type redisClientImpl struct {
	cache *cache.Cache
}

func NewRedisClient(config configuration.RedisConfig) RedisClient {
	ring := redis.NewRing(config.GetRingOptions())

	mycache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return &redisClientImpl{cache: mycache}
}

func (r *redisClientImpl) Get(ctx context.Context, key string, value interface{}) error {
	ctx, span := obs.NewSpan(ctx, "redis-get")
	defer span.End()
	obs.WithString(span, "redis.key", key)

	return r.cache.Get(ctx, key, value)
}

func (r *redisClientImpl) Set(ctx context.Context, key string, value interface{}, longevity time.Duration) error {
	ctx, span := obs.NewSpan(ctx, "redis-set")
	defer span.End()
	obs.WithString(span, "redis.key", key)

	return r.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   longevity,
	})
}
