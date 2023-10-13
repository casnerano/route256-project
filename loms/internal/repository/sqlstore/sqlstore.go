package sqlstore

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"log"
	"route256/loms/internal/repository"
)

type store interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)

	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Provider interface {
	Store(ctx context.Context) store
}

type Transactor struct {
	store store
}

func NewTransactor(store store) *Transactor {
	return &Transactor{store: store}
}

func (m *Transactor) RunRepeatableRead(ctx context.Context, fn func(ctx context.Context) error) error {
	return m.runTransaction(ctx, pgx.RepeatableRead, fn)
}

func (m *Transactor) runTransaction(ctx context.Context, level pgx.TxIsoLevel, fn func(ctx context.Context) error) error {
	tx, err := m.store.BeginTx(ctx, pgx.TxOptions{IsoLevel: level})
	if err != nil {
		return err
	}
	defer func() {
		if err = tx.Rollback(ctx); err != nil {
			log.Println("Failed rollback transaction: ", err)
		}
	}()

	if err = fn(context.WithValue(ctx, repository.ContextTxKey, tx)); err != nil {
		return fmt.Errorf("failed exec transaction callback: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed commit transaction: %w", err)
	}

	return nil
}

func (m *Transactor) Store(ctx context.Context) store {
	if tx, ok := ctx.Value(repository.ContextTxKey).(store); ok && tx != nil {
		return tx
	}
	return m.store
}
