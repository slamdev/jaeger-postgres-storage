package mtype

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/jaegertracing/jaeger/model"
)

const (
	stringType  = "string"
	boolType    = "bool"
	int64Type   = "int64"
	float64Type = "float64"
	binaryType  = "binary"
)

var domainToDBValueTypeMap = map[model.ValueType]string{
	model.StringType:  stringType,
	model.BoolType:    boolType,
	model.Int64Type:   int64Type,
	model.Float64Type: float64Type,
	model.BinaryType:  binaryType,
}

type KeyValue struct {
	Key         string  `json:"key"`
	ValueType   string  `json:"value_type"`
	ValueString string  `json:"value_string"`
	ValueBool   bool    `json:"value_bool"`
	ValueLong   int64   `json:"value_long"`
	ValueDouble float64 `json:"value_double"`
	ValueBinary []byte  `json:"value_binary"`
}

func (d *KeyValue) TypeName() string {
	return "keyvalue"
}

func (d *KeyValue) FromModel(kv model.KeyValue) {
	d.Key = kv.Key
	d.ValueType = domainToDBValueTypeMap[kv.VType]
	d.ValueString = kv.VStr
	d.ValueBool = kv.Bool()
	d.ValueLong = kv.Int64()
	d.ValueDouble = kv.Float64()
	d.ValueBinary = kv.Binary()
}

func FromKeyValueModels(kvs []model.KeyValue) []KeyValue {
	retMe := make([]KeyValue, len(kvs))
	for i := range kvs {
		kv := KeyValue{}
		kv.FromModel(kvs[i])
		retMe[i] = kv
	}
	return retMe
}

type KeyValueArray []KeyValue

func (g KeyValueArray) Value() (driver.Value, error) {
	return json.Marshal(g)
}

func (g *KeyValueArray) Scan(src interface{}) error {
	bytes := src.([]byte)
	return json.Unmarshal(bytes, g)
}
