package log

import (
	"io"
	"time"

	"github.com/uber-go/zap"
)

const (
	dateFormat = "2006-01-02T15:04:05.999999Z07:00"
)

// Tracer contains the root of the logging structure
type Tracer struct {
	logger zap.Logger
	level  zap.AtomicLevel
}

func (t *Tracer) Debug() {
	t.level.SetLevel(zap.DebugLevel)
}

type timestamp struct {
	format string
}

func (t *timestamp) String() string {
	return time.Now().UTC().Format(t.format)
}

type config struct {
	encoder zap.Encoder
	options []zap.Option
}

type Option func(*config)

// NewTracer constructs a new tracer that logs to the specified writer
func NewTracer(out io.Writer, opts ...Option) *Tracer {
	var w zap.WriteSyncer
	if v, ok := out.(zap.WriteSyncer); ok {
		w = v
	} else {
		w = zap.AddSync(out)
	}

	skipField := zap.Skip()
	noLevel := zap.LevelFormatter(func(zap.Level) zap.Field { return skipField })

	c := &config{
		encoder: zap.NewJSONEncoder(zap.NoTime(), noLevel),
		options: []zap.Option{
			zap.Fields(zap.Stringer("timestamp", &timestamp{format: dateFormat})),
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	atomicLevel := zap.DynamicLevel()

	return &Tracer{
		logger: zap.New(c.encoder, append(c.options, zap.Output(w), atomicLevel)...),
		level:  atomicLevel,
	}
}

func Text() Option {
	return func(c *config) {
		c.encoder = zap.NewTextEncoder(zap.TextNoTime())
		c.options = []zap.Option{}
	}
}
