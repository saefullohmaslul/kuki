package jeager

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jeagerCfg "github.com/uber/jaeger-client-go/config"
)

func InitializeJaeger(serviceName string) (io.Closer, error) {
	cfg:= jeagerCfg.Configuration{
        ServiceName: serviceName,
        Sampler: &jeagerCfg.SamplerConfig{
            Type:  "const",
            Param: 1,
        },
        Reporter: &jeagerCfg.ReporterConfig{
            LogSpans: true,
            CollectorEndpoint: fmt.Sprintf("http://%s:14268/api/traces", "jaeger"),
        },
    }

    tracer, closer, err := cfg.NewTracer(
        jeagerCfg.Logger(jaeger.StdLogger),
    )
    if err != nil {
        return nil, err
    }

    opentracing.SetGlobalTracer(tracer)

    return closer, nil
}