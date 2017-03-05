package log_test

import (
	"os"
	"testing"

	"github.com/altairsix/log"
)

func TestFunc(t *testing.T) {
	//log.DefaultTracer = log.NewTracer(os.Stdout)
	log.DefaultTracer = log.NewTracer(os.Stdout, log.Text())
	log.DefaultTracer.Debug()

	_, span := log.Start("root", log.String("hello", "world"))
	span.Debug("woot!")
	span.Close(nil, "ok", log.String("a", "b"))
	span.Close(nil, "ok")
	span.Close(nil, "ok")
	span.Close(nil, "ok")
}
