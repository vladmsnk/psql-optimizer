package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InfoProvider interface {
	ListIndexesStat(ctx context.Context) ([]IndexStat, error)
}

type Adapter struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Adapter {
	return &Adapter{
		pool: pool,
	}
}

func (a *Adapter) ListIndexesStat(_ context.Context) ([]IndexStat, error) {
	return nil, nil
}
