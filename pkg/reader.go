package pkg

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/jaegertracing/jaeger/model"
	"github.com/jaegertracing/jaeger/storage/spanstore"
	"github.com/slamdev/jaeger-postgres-storage/pkg/dbmodel/table"
)

type reader struct {
	repo Repository
	log  hclog.Logger
}

func NewReader(repo Repository, log hclog.Logger) spanstore.Reader {
	return &reader{
		repo: repo,
		log:  log.Named("reader"),
	}
}

func (r *reader) GetTrace(ctx context.Context, traceID model.TraceID) (*model.Trace, error) {
	log := r.log.Named("GetTrace")
	log.Debug("inputs", "traceID", traceID)
	return nil, nil
}

func (r *reader) GetServices(ctx context.Context) ([]string, error) {
	log := r.log.Named("GetServices")
	log.Debug("inputs")
	serviceNames, err := r.repo.GetServiceNames(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get service names; %w", err)
	}
	names := make([]string, len(serviceNames))
	for i := range serviceNames {
		names[i] = serviceNames[i].ServiceName
	}
	return names, nil
}

func (r *reader) GetOperations(ctx context.Context, query spanstore.OperationQueryParameters) ([]spanstore.Operation, error) {
	log := r.log.Named("GetOperations")
	log.Debug("inputs", "query", query)
	operationNames, err := r.repo.GetOperationNames(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get operation names; %w", err)
	}
	return table.ToSpanOperations(operationNames), nil
}

func (r *reader) FindTraces(ctx context.Context, query *spanstore.TraceQueryParameters) ([]*model.Trace, error) {
	log := r.log.Named("FindTraces")
	log.Debug("inputs", "query", query)
	return []*model.Trace{}, nil
}

func (r *reader) FindTraceIDs(ctx context.Context, query *spanstore.TraceQueryParameters) ([]model.TraceID, error) {
	log := r.log.Named("FindTraceIDs")
	log.Debug("inputs", "query", query)
	return []model.TraceID{}, nil
}
