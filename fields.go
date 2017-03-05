package log

import (
	"fmt"

	"github.com/uber-go/zap"
)

var (
	skipField = Field(zap.Skip())
)

type Field zap.Field

func String(k, v string) Field {
	return Field(zap.String(k, v))
}

func Int(k string, v int) Field {
	return Field(zap.Int(k, v))
}

func Int64(k string, v int64) Field {
	return Field(zap.Int64(k, v))
}

func Float64(k string, v float64) Field {
	return Field(zap.Float64(k, v))
}

func Stringer(k string, v fmt.Stringer) Field {
	return Field(zap.Stringer(k, v))
}

func Error(err error) Field {
	if err == nil {
		return skipField
	}

	return String("err", err.Error())
}

