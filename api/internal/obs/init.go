package obs

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

// InitTracer creates a new trace provider instance and registers it as global trace provider.
func InitTracer(config TraceConfig) func() {
	log.Info().Msg("Initializing tracer...")
	if config.Noop {
		log.Info().Msg("Using noop tracing!")
		otel.SetTracerProvider(trace.NewNoopTracerProvider())
		return func() {}
	}

	// return initGoogleTraceTracer(config)
	return initJaegerTracer(config)
}

func initJaegerTracer(config TraceConfig) func() {
	sampler := getSampler(config)

	jaegerHost := config.JaegerConfig.Host
	log.Info().Msgf("Exporting to Jaeger host: %s", jaegerHost)

	// Create and install Jaeger export pipeline
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint(fmt.Sprintf("http://%s:14268/api/traces", jaegerHost)),
		jaeger.WithProcessFromEnv(),
		jaeger.WithSDKOptions(sdktrace.WithSampler(sampler)),
	)

	// Handler error
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
