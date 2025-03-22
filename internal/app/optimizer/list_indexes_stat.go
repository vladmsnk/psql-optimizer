package optimizer

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	
	"psql-optimizer/internal/usecases/entities"
	desc "psql-optimizer/pkg/optimizer"
)

func (s *Service) ListIndexesStat(ctx context.Context, _ *desc.ListIndexesStatRequest) (*desc.ListIndexesStatResponse, error) {
	indexesStat, err := s.indexInfoGetter.ListIndexesStat(ctx)
	if err != nil {
		return nil, fmt.Errorf("indexInfoGetter.ListIndexesStat: %w", err)
	}

	stats := lo.Map(indexesStat, func(stat entities.IndexStat, _ int) *desc.ListIndexesStatResponse_IndexStat {
		return toDescIndexesStat(stat)
	})

	return &desc.ListIndexesStatResponse{IndexesStat: stats}, nil
}

func toDescIndexesStat(stat entities.IndexStat) *desc.ListIndexesStatResponse_IndexStat {
	return &desc.ListIndexesStatResponse_IndexStat{
		TableName:     stat.TableName,
		IndexName:     stat.IndexName,
		IndexScans:    stat.IndexScans,
		TuplesRead:    stat.TuplesRead,
		TuplesFetched: stat.TuplesFetched,
	}
}
