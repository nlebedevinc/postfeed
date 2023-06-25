package models

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Timeline struct {
	rdb *redis.Client
}

func NewTimeline(rdb *redis.Client) Timeline {
	return Timeline{rdb: rdb}
}

func (t Timeline) Push(ctx context.Context, user string, post ...interface{}) error {
	return t.rdb.RPush(ctx, "timeline:"+user, post...).Err()
}

func (t Timeline) Latest(ctx context.Context, user string, count int64) ([]string, error) {
	return t.rdb.LRange(ctx, "timeline:"+user, -1*count, -1).Result()
}
