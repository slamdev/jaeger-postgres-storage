package table

type TagIndex struct {
	ServiceName string `db:"service_name"`
	TagKey      string `db:"tag_key"`
	TagValue    string `db:"tag_value"`
	StartTime   int64  `db:"start_time"`
	TraceID     []byte `db:"trace_id"`
	SpanID      int64  `db:"span_id"`
}

func (t *TagIndex) TableName() string {
	return "tag_index"
}
