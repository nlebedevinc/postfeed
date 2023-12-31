package services

import (
	"context"
	"encoding/json"

	"github.com/nlebedevinc/postfeed/internal/models"
	"github.com/redis/go-redis/v9"
)

type Redis[T models.Keyer] struct {
	rdb *redis.Client
}

func (r Redis[T]) Save(ctx context.Context, k T) error {
	b, _ := json.Marshal(k)
	return r.rdb.Set(ctx, k.Key(), b, 0).Err()
}

func (r Redis[T]) Get(ctx context.Context, key string) (T, error) {
	var t T
	b, err := r.rdb.Get(ctx, key).Bytes()
	if err != nil {
		return t, err
	}
	json.Unmarshal(b, &t)
	return t, nil
}

func (r Redis[T]) MGet(ctx context.Context, key ...string) ([]T, error) {
	bb, err := r.rdb.MGet(ctx, key...).Result()
	if err != nil {
		return nil, err
	}

	result := make([]T, len(key))
	for i, b := range bb {
		json.Unmarshal([]byte(b.(string)), &result[i])
	}
	return result, nil
}

func NewRedis[T models.Keyer](rdb *redis.Client) Redis[T] {
	r := Redis[T]{rdb: rdb}
	return r
}
