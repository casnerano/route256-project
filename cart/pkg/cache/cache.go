package cache

//go:generate mockgen -destination=mock/cache.go -source=cache.go

import "context"

type Cache interface {
	Has(context.Context, string) (bool, error)
	Get(context.Context, string) (*string, error)
	Set(context.Context, string, string) error
	Delete(context.Context, string) error
}
