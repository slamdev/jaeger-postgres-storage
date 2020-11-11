package table

import (
	"time"
)

type DurationIndex struct {
	ServiceName   string    `db:"service_name"`
	OperationName string    `db:"operation_name"`
	Bucket        time.Time `db:"bucket"`
	Duration      int64     `db:"duration"`
	StartTime     int64     `db:"start_time"`
	TraceID       []byte    `db:"trace_id"`
}

func (d *DurationIndex) TableName() string {
	return "duration_index"
}
