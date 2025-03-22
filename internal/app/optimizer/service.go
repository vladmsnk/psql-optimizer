package optimizer

import (
	"psql-optimizer/internal/usecases/indexes"
	desc "psql-optimizer/pkg/optimizer"
)

type Service struct {
	desc.OptimizerAPIServer
	indexInfoGetter indexes.IndexInfoGetter
}

func New(indexInfoGetter indexes.IndexInfoGetter) *Service {
	return &Service{
		indexInfoGetter: indexInfoGetter,
	}
}
