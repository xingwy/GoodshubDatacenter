package dao

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisMangement interface {
	Subscribe(ctx context.Context, c ...string) *redis.PubSub
}

type RedisMangementInstance struct {
	D *Dao
}

func (i *RedisMangementInstance) Subscribe(ctx context.Context, c ...string) *redis.PubSub {
	return i.D.Redis().Subscribe(ctx, c...)
}
