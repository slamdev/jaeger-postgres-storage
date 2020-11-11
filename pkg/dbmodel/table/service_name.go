package table

import (
	"github.com/jaegertracing/jaeger/model"
)

type ServiceName struct {
	ServiceName string `db:"service_name"`
}

func (s *ServiceName) TableName() string {
	return "service_names"
}

func (s *ServiceName) FromSpan(span *model.Span) {
	s.ServiceName = span.Process.ServiceName
}
