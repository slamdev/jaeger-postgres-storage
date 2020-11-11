package mtype

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/jaegertracing/jaeger/model"
	"github.com/slamdev/jaeger-postgres-storage/pkg/dbmodel"
)

const (
	childOf     = "child-of"
	followsFrom = "follows-from"
)

var domainToDBRefMap = map[model.SpanRefType]string{
	model.SpanRefType_CHILD_OF:     childOf,
	model.SpanRefType_FOLLOWS_FROM: followsFrom,
}

type SpanRef struct {
	RefType string   `json:"ref_type"`
	TraceID [16]byte `json:"trace_id"`
	SpanID  int64    `json:"span_id"`
}

func (d *SpanRef) TypeName() string {
	return "span_ref"
}

func (d *SpanRef) FromModel(ref model.SpanRef) {
	d.RefType = domainToDBRefMap[ref.RefType]
	d.TraceID = dbmodel.TraceIDFromDomain(ref.TraceID)
	d.SpanID = int64(ref.SpanID)
}

func FromSpanRefModels(refs []model.SpanRef) []SpanRef {
	retMe := make([]SpanRef, len(refs))
	for i := range refs {
		ref := SpanRef{}
		ref.FromModel(refs[i])
		retMe[i] = ref
	}
	return retMe
}

type SpanRefArray []SpanRef

func (g SpanRefArray) Value() (driver.Value, error) {
	return json.Marshal(g)
}

func (g *SpanRefArray) Scan(src interface{}) error {
	bytes := src.([]byte)
	return json.Unmarshal(bytes, g)
}
