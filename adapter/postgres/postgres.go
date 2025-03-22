package postgres

import (
	"context"
	"fmt"

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

func (a *Adapter) ListIndexesStat(ctx context.Context) ([]IndexStat, error) {
	query := `
select
    s.relname table_name,
    i.relname index_name,
    ix.idx_scan index_scans,
    ix.idx_tup_read tuples_read,
    ix.idx_tup_fetch tuples_fetched
from pg_stat_user_indexes ix
join pg_class s ON s.oid = ix.relid
join pg_class i ON i.oid = ix.indexrelid
`

	rows, err := a.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("pool.Query: %w", err)
	}
	defer rows.Close()

	result := make([]IndexStat, 0)

	for rows.Next() {
		is := IndexStat{}

		err := rows.Scan(&is.TableName, &is.IndexName, &is.IndexScans, &is.TuplesRead, &is.TuplesFetched)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		result = append(result, is)
	}

	return result, nil
}
