package g

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
)

func TraceInit(serviceName string, samplerType string, samplerParam float64) (opentracing.Tracer, io.Closer) {
	cfg := &jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  samplerType,
			Param: samplerParam,
		},
		Reporter: &jaegercfg.ReporterConfig{
			//LocalAgentHostPort: "81.70.77.142:6831",
			LogSpans: true,
		},
	}

	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger), jaegercfg.MaxTagValueLength(10000))
	if err != nil {
		panic(fmt.Sprintf("Init failed: %v\n", err))
	}

	return tracer, closer
}
