package shard_strategy

type ShardingStrategy interface {
	Shard(key string) int
}
