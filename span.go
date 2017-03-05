package log

import (
	"time"

	"github.com/uber-go/zap"
)

// SpanLog handles log messages for a given span
type SpanLog struct {
	opName    string
	fields    []Field
	startedAt time.Time
	logger    zap.Logger
}

// Close should be called when when exiting the span
func (s *SpanLog) Close(err error, msg string, fields ...Field) error {
	elapsed := time.Now().Sub(s.startedAt) / time.Millisecond
	f := zapFields(s.fields, fields, Error(err), Int64("elapsed", int64(elapsed)))
	s.logger.Info(msg, f...)
	return err
}

// Debug just prints output
func (s *SpanLog) Debug(msg string, fields ...Field) {
	f := zapFields(s.fields, fields)
	s.logger.Debug(msg, f...)
}

// Info just prints output
func (s *SpanLog) Info(msg string, fields ...Field) {
	f := zapFields(s.fields, fields)
	s.logger.Info(msg, f...)
}

func (s *SpanLog) zapFields(fields []Field, additional ...Field) []zap.Field {
	zapFields := make([]zap.Field, 0, len(s.fields)+len(fields)+len(additional))

	for _, field := range additional {
		zapFields = append(zapFields, zap.Field(field))
	}

	for _, field := range s.fields {
		zapFields = append(zapFields, zap.Field(field))
	}

	for _, field := range fields {
		zapFields = append(zapFields, zap.Field(field))
	}

	return zapFields
}

func zapFields(parent, local []Field, additional ...Field) []zap.Field {
	zapFields := make([]zap.Field, 0, len(parent)+len(local)+len(additional))

	for _, field := range additional {
		zapFields = append(zapFields, zap.Field(field))
	}

	for _, field := range parent {
		zapFields = append(zapFields, zap.Field(field))
	}

	for _, field := range local {
		zapFields = append(zapFields, zap.Field(field))
	}

	return zapFields
}
