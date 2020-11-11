package table

type ServiceNameIndex struct {
	ServiceName string `db:"service_name"`
	Bucket      int    `db:"bucket"`
	StartTime   int64  `db:"start_time"`
	TraceID     []byte `db:"trace_id"`
}

func (s *ServiceNameIndex) TableName() string {
	return "service_name_index"
}
