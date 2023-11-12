package provider

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(client *redis.Client) *Redis {
	return &Redis{
		client: client,
	}
}

func (r *Redis) Has(ctx context.Context, key string) (bool, error) {
	count, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *Redis) Get(ctx context.Context, key string) (*string, error) {
	value, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &value, nil
}

func (r *Redis) Set(ctx context.Context, key string, value string) error {
	_, err := r.client.Set(ctx, key, value, time.Duration(0)).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) Delete(ctx context.Context, key string) error {
	_, err := r.client.Del(ctx, key, key).Result()
	if err != nil {
		return err
	}

	return nil
}
