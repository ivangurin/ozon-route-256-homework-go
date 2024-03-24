package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
)

type Pool interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	Ping(ctx context.Context) error
	Close() error
}

type pool struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func NewPool(ctx context.Context, dsn string) (Pool, error) {
	p, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create db pool for %s: %v", dsn, err)
	}
	err = p.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping db for %s: %v", dsn, err)
	}

	return &pool{
		ctx:  ctx,
		pool: p,
	}, nil
}

func (p *pool) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return p.pool.Exec(ctx, sql, args...)
}

func (p *pool) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return p.pool.Query(ctx, sql, args...)
}

func (p *pool) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return p.pool.QueryRow(ctx, sql, args...)
}

func (p *pool) BeginFunc(ctx context.Context, f func(pgx.Tx) error) error {
	conn, err := p.pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %v", err)
	}

	tx, err := conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	var rollbackErr error
	defer func(tx pgx.Tx, ctx context.Context) {
		rollbackErr = tx.Rollback(ctx)
		if rollbackErr != nil {
			logger.Errorf(ctx, "failed to rollback transaction: %v", rollbackErr)
		}
	}(tx, ctx)

	err = f(tx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (p *pool) Ping(ctx context.Context) error {
	return p.pool.Ping(ctx)
}

func (p *pool) Close() error {
	p.pool.Close()
	return nil
}