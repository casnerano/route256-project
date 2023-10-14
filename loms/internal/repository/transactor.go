package repository

import "context"

type contextTxKeyType string

const ContextTxKey contextTxKeyType = "repository_tx"

type Transactor interface {
	RunRepeatableRead(context.Context, func(context.Context) error) error
}
