package pool

import (
	"context"

	"route256/cart/internal/service/cache/pool/shard_strategy"
	"route256/cart/pkg/cache"
)

type Pool struct {
	providers        []cache.Cache
	shardingStrategy shard_strategy.ShardingStrategy
}

func New(shardingStrategy shard_strategy.ShardingStrategy, provider ...cache.Cache) *Pool {
	return &Pool{
		providers:        provider,
		shardingStrategy: shardingStrategy,
	}
}

func (p *Pool) Has(ctx context.Context, key string) (bool, error) {
	return p.getProvider(key).Has(ctx, key)
}

func (p *Pool) Get(ctx context.Context, key string) (string, error) {
	return p.getProvider(key).Get(ctx, key)
}

func (p *Pool) Set(ctx context.Context, key string, value string) error {
	return p.getProvider(key).Set(ctx, key, value)
}

func (p *Pool) Delete(ctx context.Context, key string) error {
	return p.getProvider(key).Delete(ctx, key)
}

func (p *Pool) getProvider(key string) cache.Cache {
	shard := p.shardingStrategy.Shard(key)
	if shard < 0 || shard >= len(p.providers) {
		shard = 0
	}
	return p.providers[shard]
}
