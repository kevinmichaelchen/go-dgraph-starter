package obs

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/label"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

// InitTracer creates a new trace provider instance and registers it as global trace provider.
func InitTracer(config TraceConfig, appName, appID string) func() {
	log.Info().Msg("Initializing tracer...")
	if config.Noop {
		log.Info().Msg("Using noop tracing!")
		otel.SetTracerProvider(trace.NewNoopTracerProvider())
		return func() {}
	}

	// return initGoogleTraceTracer(config)
	return initJaegerTracer(config, appName, appID)
}

func initJaegerTracer(config TraceConfig, appName, appID string) func() {
	sampler := getSampler(config)

	// TODO what should this be?
	var jaegerHost string

	// Create and install Jaeger export pipeline
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint(fmt.Sprintf("http://%s:14268/api/traces", jaegerHost)),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: appName,
			Tags: []label.KeyValue{
				label.String("exporter", "jaeger"),
				label.String("app_id", appID),
			},
		}),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sampler}),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create Jaeger trace export pipeline")
	}

	return flush
}

func getSampler(config TraceConfig) sdktrace.Sampler {
	var sampler sdktrace.Sampler
	if ratio := config.SampleRatio; ratio == 1 {
		log.Info().Msg("Creating trace sampler with AlwaysSample")
		sampler = sdktrace.AlwaysSample()
	} else {
		log.Info().Msgf("Creating trace sampler with ratio: %v", config.SampleRatio)
		sampler = sdktrace.TraceIDRatioBased(ratio)
	}
	return sampler
}
