package pkg

import (
	"context"
	"encoding/json"
	"github.com/jaegertracing/jaeger/storage/spanstore"
	"github.com/slamdev/jaeger-postgres-storage/pkg/dbmodel/table"
	"testing"
)

const traceJson = `
{
  "trace_id":[
    0,
    0,
    0,
    0,
    0,
    0,
    0,
    0,
    19,
    67,
    5,
    116,
    187,
    224,
    140,
    17
  ],
  "span_id":6455130307893583360,
  "span_hash":6184321674429755536,
  "operation_name":"GetDriver",
  "flags":1,
  "start_time":1587212688759257,
  "duration":37811,
  "tags":[
    {
      "key":"param.driverID",
      "value_type":"string",
      "value_string":"T757044C",
      "value_bool":false,
      "value_long":0,
      "value_double":0,
      "value_binary":null
    },
    {
      "key":"span.kind",
      "value_type":"string",
      "value_string":"client",
      "value_bool":false,
      "value_long":0,
      "value_double":0,
      "value_binary":null
    },
    {
      "key":"error",
      "value_type":"bool",
      "value_string":"",
      "value_bool":true,
      "value_long":0,
      "value_double":0,
      "value_binary":null
    },
    {
      "key":"internal.span.format",
      "value_type":"string",
      "value_string":"proto",
      "value_bool":false,
      "value_long":0,
      "value_double":0,
      "value_binary":null
    }
  ],
  "logs":[
    {
      "ts":1587212688796760,
      "fields":[
        {
          "key":"event",
          "value_type":"string",
          "value_string":"redis timeout",
          "value_bool":false,
          "value_long":0,
          "value_double":0,
          "value_binary":null
        },
        {
          "key":"level",
          "value_type":"string",
          "value_string":"error",
          "value_bool":false,
          "value_long":0,
          "value_double":0,
          "value_binary":null
        },
        {
          "key":"driver_id",
          "value_type":"string",
          "value_string":"T757044C",
          "value_bool":false,
          "value_long":0,
          "value_double":0,
          "value_binary":null
        },
        {
          "key":"error",
          "value_type":"string",
          "value_string":"redis timeout",
          "value_bool":false,
          "value_long":0,
          "value_double":0,
          "value_binary":null
        }
      ]
    }
  ],
  "refs":[
    {
      "ref_type":"child-of",
      "trace_id":[
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        19,
        67,
        5,
        116,
        187,
        224,
        140,
        17
      ],
      "span_id":5117007496139890642
    }
  ],
  "process":{
    "service_name":"redis",
    "tags":[
      {
        "key":"jaeger.version",
        "value_type":"string",
        "value_string":"Go-2.22.1",
        "value_bool":false,
        "value_long":0,
        "value_double":0,
        "value_binary":null
      },
      {
        "key":"hostname",
        "value_type":"string",
        "value_string":"caa9afce8db3",
        "value_bool":false,
        "value_long":0,
        "value_double":0,
        "value_binary":null
      },
      {
        "key":"ip",
        "value_type":"string",
        "value_string":"172.18.0.4",
        "value_bool":false,
        "value_long":0,
        "value_double":0,
        "value_binary":null
      },
      {
        "key":"client-uuid",
        "value_type":"string",
        "value_string":"654aa60811b6ae1c",
        "value_bool":false,
        "value_long":0,
        "value_double":0,
        "value_binary":null
      }
    ]
  }
}
`

func TestRepository_Should_Save_Trace(t *testing.T) {
	trace := table.Trace{}
	json.Unmarshal([]byte(traceJson), &trace)
	r, err := NewRepository(Config{
		Host:              "localhost",
		Port:              5432,
		Username:          "postgres",
		Password:          "postgres",
		Database:          "jaeger",
		ConnectionOptions: "sslmode=disable",
	})
	if err := r.Migrate(); err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	s, err := r.GetServiceNames(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	println(s)
	o, err := r.GetOperationNames(context.Background(), spanstore.OperationQueryParameters{
		ServiceName: s[0].ServiceName,
	})
	println(o)
	//if err := r.SaveTrace(trace); err != nil {
	//	t.Fatal(err)
	//}
}
