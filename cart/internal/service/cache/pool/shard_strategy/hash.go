package shard_strategy

import (
	"hash/crc32"
)

type hash struct {
	shardCount int
}

func NewHash(shardCount int) *hash {
	return &hash{
		shardCount: shardCount,
	}
}

func (h *hash) Shard(key string) int {
	return int(crc32.ChecksumIEEE([]byte(key))) % h.shardCount
}
