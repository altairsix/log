package log

import (
	"context"
	"io/ioutil"
	"time"
)

var DefaultTracer = NewTracer(ioutil.Discard)

func Start(operationName string, fields ...Field) (context.Context, *SpanLog) {
	return StartContext(context.Background(), operationName, fields...)
}

func StartContext(parent context.Context, opName string, fields ...Field) (context.Context, *SpanLog) {
	span := &SpanLog{
		opName:    opName,
		fields:    fields,
		startedAt: time.Now(),
		logger:    DefaultTracer.logger,
	}

	return parent, span
}
