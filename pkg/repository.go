package pkg

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/jaegertracing/jaeger/storage/spanstore"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/slamdev/jaeger-postgres-storage/assets/migrations"
	"github.com/slamdev/jaeger-postgres-storage/pkg/dbmodel/table"
)

const (
	insertTraceQuery = `
insert into %s 
  (trace_id, span_id, span_hash, parent_id, operation_name, flags, start_time, duration, tags, logs, refs, process)
values 
  (:trace_id, :span_id, :span_hash, :parent_id, :operation_name, :flags, :start_time, :duration, :tags, :logs, :refs, :process)
	`
	insertServiceNameQuery = `
insert into %s 
  (service_name)
values 
  (:service_name)
on conflict do nothing
	`
	selectServiceNameQuery = `
select * from %s
	`
	insertOperationNameQuery = `
insert into %s 
  (service_name, span_kind, operation_name)
values 
  (:service_name, :span_kind, :operation_name)
on conflict do nothing
	`
	selectOperationNameQuery = `
select * from %s 
where
  service_name = $1
	`
	selectOperationNameQueryByKind = `
select * from %s 
where
  service_name = $1 and span_kind = $2
	`
)

type repository struct {
	db *sqlx.DB
}

type Repository interface {
	Close() error
	Migrate() error
	SaveTrace(trace table.Trace) error
	SaveServiceName(serviceName table.ServiceName) error
	SaveOperationName(operationName table.OperationName) error
	GetServiceNames(ctx context.Context) ([]table.ServiceName, error)
	GetOperationNames(ctx context.Context, query spanstore.OperationQueryParameters) ([]table.OperationName, error)
}

func NewRepository(config Config) (Repository, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s",
		config.Username, config.Password, config.Host, config.Port, config.Database, config.ConnectionOptions)
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres instance at %s:***@%s:%d/%s?%s; %w",
			config.Username, config.Host, config.Port, config.Database, config.ConnectionOptions, err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping postgres instance at %s:***@%s:%d/%s?%s; %w",
			config.Username, config.Host, config.Port, config.Database, config.ConnectionOptions, err)
	}
	return &repository{db: db}, nil
}

func (r *repository) Migrate() error {
	s := bindata.Resource(migrations.AssetNames(), func(name string) ([]byte, error) {
		return migrations.Asset(name)
	})
	d, err := bindata.WithInstance(s)
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(r.db.DB, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("bindata", d, "postgres", driver)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		return err
	}
	return nil
}

func (r *repository) Close() error {
	return r.db.Close()
}

func (r *repository) SaveTrace(trace table.Trace) error {
	q := fmt.Sprintf(insertTraceQuery, trace.TableName())
	_, err := r.db.NamedExec(q, trace)
	return err
}

func (r *repository) SaveServiceName(serviceName table.ServiceName) error {
	q := fmt.Sprintf(insertServiceNameQuery, serviceName.TableName())
	_, err := r.db.NamedExec(q, serviceName)
	return err
}

func (r *repository) GetServiceNames(ctx context.Context) ([]table.ServiceName, error) {
	q := fmt.Sprintf(selectServiceNameQuery, (&table.ServiceName{}).TableName())
	var serviceNames []table.ServiceName
	err := r.db.SelectContext(ctx, &serviceNames, q)
	return serviceNames, err
}

func (r *repository) SaveOperationName(operationName table.OperationName) error {
	q := fmt.Sprintf(insertOperationNameQuery, operationName.TableName())
	_, err := r.db.NamedExec(q, operationName)
	return err
}

func (r *repository) GetOperationNames(ctx context.Context, query spanstore.OperationQueryParameters) ([]table.OperationName, error) {
	var operationNames []table.OperationName
	var err error
	if query.SpanKind == "" {
		q := fmt.Sprintf(selectOperationNameQuery, (&table.OperationName{}).TableName())
		err = r.db.SelectContext(ctx, &operationNames, q, query.ServiceName)
	} else {
		q := fmt.Sprintf(selectOperationNameQueryByKind, (&table.OperationName{}).TableName())
		err = r.db.SelectContext(ctx, &operationNames, q, query.ServiceName, query.SpanKind)
	}
	return operationNames, err
}
