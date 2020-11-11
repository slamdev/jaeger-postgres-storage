package table

import (
	"github.com/jaegertracing/jaeger/model"
	"github.com/slamdev/jaeger-postgres-storage/pkg/dbmodel"
	mtype "github.com/slamdev/jaeger-postgres-storage/pkg/dbmodel/type"
)

type Trace struct {
	TraceID       []byte              `db:"trace_id" json:"trace_id"`
	SpanID        int64               `db:"span_id" json:"span_id"`
	SpanHash      int64               `db:"span_hash" json:"span_hash"`
	ParentID      int64               `db:"parent_id" json:"parent_id"` // deprecate
	OperationName string              `db:"operation_name" json:"operation_name"`
	Flags         int32               `db:"flags" json:"flags"`
	StartTime     int64               `db:"start_time" json:"start_time"`
	Duration      int64               `db:"duration" json:"duration"`
	Tags          mtype.KeyValueArray `db:"tags" json:"tags"`
	Logs          mtype.LogArray      `db:"logs" json:"logs"`
	Refs          mtype.SpanRefArray  `db:"refs" json:"refs"`
	Process       mtype.Process       `db:"process" json:"process"`
}

func (t *Trace) TableName() string {
	return "traces"
}

func (t *Trace) FromSpan(span *model.Span) {
	tags := mtype.FromKeyValueModels(span.Tags)
	logs := mtype.FromLogModels(span.Logs)
	refs := mtype.FromSpanRefModels(span.References)
	process := mtype.Process{}
	process.FromModel(*span.Process)
	spanHash, _ := model.HashCode(span)

	traceID := dbmodel.TraceIDFromDomain(span.TraceID)
	t.TraceID = traceID[:]
	t.SpanID = int64(span.SpanID)
	t.SpanHash = int64(spanHash)
	t.OperationName = span.OperationName
	t.Flags = int32(span.Flags)
	t.StartTime = int64(model.TimeAsEpochMicroseconds(span.StartTime))
	t.Duration = int64(model.DurationAsMicroseconds(span.Duration))
	t.Tags = tags
	t.Logs = logs
	t.Refs = refs
	t.Process = process
}
