package entities

type IndexStat struct {
	TableName     string
	IndexName     string
	IndexScans    float32
	TuplesRead    float32
	TuplesFetched float32
}
