package mtype

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/jaegertracing/jaeger/model"
)

type Log struct {
	Ts     int64      `json:"ts"`
	Fields []KeyValue `json:"fields"`
}

func (d *Log) TypeName() string {
	return "log"
}

func (d *Log) FromModel(log model.Log) {
	d.Ts = int64(model.TimeAsEpochMicroseconds(log.Timestamp))
	d.Fields = FromKeyValueModels(log.Fields)
}

func FromLogModels(logs []model.Log) []Log {
	retMe := make([]Log, len(logs))
	for i := range logs {
		log := Log{}
		log.FromModel(logs[i])
		retMe[i] = log
	}
	return retMe
}

type LogArray []Log

func (g LogArray) Value() (driver.Value, error) {
	return json.Marshal(g)
}

func (g *LogArray) Scan(src interface{}) error {
	bytes := src.([]byte)
	return json.Unmarshal(bytes, g)
}
