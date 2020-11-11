package table

import (
	"github.com/jaegertracing/jaeger/model"
	"github.com/jaegertracing/jaeger/storage/spanstore"
)

type OperationName struct {
	ServiceName   string `db:"service_name"`
	SpanKind      string `db:"span_kind"`
	OperationName string `db:"operation_name"`
}

func (o *OperationName) TableName() string {
	return "operation_names_v2"
}

func (o *OperationName) FromSpan(span *model.Span) {
	spanKind, _ := span.GetSpanKind()
	o.SpanKind = spanKind
	o.ServiceName = span.Process.ServiceName
	o.OperationName = span.OperationName
}

func (o *OperationName) ToSpanOperation() spanstore.Operation {
	return spanstore.Operation{
		Name:     o.OperationName,
		SpanKind: o.SpanKind,
	}
}

func ToSpanOperations(operationNames []OperationName) []spanstore.Operation {
	spanOperations := make([]spanstore.Operation, len(operationNames))
	for i := range operationNames {
		spanOperations[i] = operationNames[i].ToSpanOperation()
	}
	return spanOperations
}
