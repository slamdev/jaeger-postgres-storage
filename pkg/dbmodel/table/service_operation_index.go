package table

type ServiceOperationIndex struct {
	ServiceName   string `db:"service_name"`
	OperationName string `db:"operation_name"`
	StartTime     int64  `db:"start_time"`
	TraceID       []byte `db:"trace_id"`
}

func (s *ServiceOperationIndex) TableName() string {
	return "service_operation_index"
}
