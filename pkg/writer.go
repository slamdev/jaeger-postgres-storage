package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/jaegertracing/jaeger/model"
	"github.com/jaegertracing/jaeger/storage/spanstore"
	"github.com/slamdev/jaeger-postgres-storage/pkg/dbmodel/table"
)

type writer struct {
	repo Repository
	log  hclog.Logger
}

func NewWriter(repo Repository, log hclog.Logger) spanstore.Writer {
	return &writer{
		repo: repo,
		log:  log.Named("writer"),
	}
}

func (w *writer) WriteSpan(span *model.Span) error {
	log := w.log.Named("WriteSpan")
	// save trace
	trace := table.Trace{}
	trace.FromSpan(span)
	b, _ := json.Marshal(trace)
	log.Debug("inputs", "trace", string(b))
	if err := w.repo.SaveTrace(trace); err != nil {
		return fmt.Errorf("failed to save trace; %w", err)
	}
	// save service name
	serviceName := table.ServiceName{}
	serviceName.FromSpan(span)
	b, _ = json.Marshal(serviceName)
	log.Debug("inputs", "serviceName", string(b))
	if err := w.repo.SaveServiceName(serviceName); err != nil {
		return fmt.Errorf("failed to save service name; %w", err)
	}
	// save operation name
	operationName := table.OperationName{}
	operationName.FromSpan(span)
	b, _ = json.Marshal(operationName)
	log.Debug("inputs", "operationName", string(b))
	if err := w.repo.SaveOperationName(operationName); err != nil {
		return fmt.Errorf("failed to save operation name; %w", err)
	}
	//
	return nil
}
