package postgres

import "psql-optimizer/internal/usecases/entities"

type IndexStat struct {
	TableName     string
	IndexName     string
	IndexScans    float32
	TuplesRead    float32
	TuplesFetched float32
}

func (i *IndexStat) ToEntity() entities.IndexStat {
	return entities.IndexStat{
		TableName:     i.TableName,
		IndexName:     i.IndexName,
		IndexScans:    i.IndexScans,
		TuplesRead:    i.TuplesRead,
		TuplesFetched: i.TuplesFetched,
	}
}
