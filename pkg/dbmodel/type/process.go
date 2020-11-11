package mtype

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/jaegertracing/jaeger/model"
)

type Process struct {
	ServiceName string     `json:"service_name"`
	Tags        []KeyValue `json:"tags"`
}

func (d *Process) TypeName() string {
	return "process"
}

func (d *Process) FromModel(proc model.Process) {
	d.ServiceName = proc.ServiceName
	d.Tags = FromKeyValueModels(proc.Tags)
}

func FromProcessModels(procs []model.Process) []Process {
	retMe := make([]Process, len(procs))
	for i := range procs {
		proc := Process{}
		proc.FromModel(procs[i])
		retMe[i] = proc
	}
	return retMe
}

func (g Process) Value() (driver.Value, error) {
	return json.Marshal(g)
}

func (g *Process) Scan(src interface{}) error {
	bytes := src.([]byte)
	return json.Unmarshal(bytes, g)
}
