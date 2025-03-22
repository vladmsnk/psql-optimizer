package indexes

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"psql-optimizer/adapter/postgres"
	"psql-optimizer/internal/usecases/entities"
)

type IndexInfoGetter interface {
	ListIndexesStat(ctx context.Context) ([]entities.IndexStat, error)
}

type PostgresInfoProvider interface {
	ListIndexesStat(ctx context.Context) ([]postgres.IndexStat, error)
}

type Implementation struct {
	provider PostgresInfoProvider
}

func New(provider PostgresInfoProvider) IndexInfoGetter {
	return &Implementation{
		provider: provider,
	}
}

func (i Implementation) ListIndexesStat(ctx context.Context) ([]entities.IndexStat, error) {
	postgresIndexStat, err := i.provider.ListIndexesStat(ctx)
	if err != nil {
		return nil, fmt.Errorf("provider.ListIndexesStat: %w", err)
	}

	result := lo.Map(postgresIndexStat, func(item postgres.IndexStat, _ int) entities.IndexStat {
		return item.ToEntity()
	})

	return result, nil
}
