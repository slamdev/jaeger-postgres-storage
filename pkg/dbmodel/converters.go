package dbmodel

import (
	"encoding/binary"
	"github.com/jaegertracing/jaeger/model"
)

func TraceIDFromDomain(traceID model.TraceID) [16]byte {
	dbTraceID := [16]byte{}
	binary.BigEndian.PutUint64(dbTraceID[:8], traceID.High)
	binary.BigEndian.PutUint64(dbTraceID[8:], traceID.Low)
	return dbTraceID
}
