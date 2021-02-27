package obs

import (
	"fmt"

	"github.com/MyOrg/go-dgraph-starter/internal/configuration"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/label"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

// InitTracer creates a new trace provider instance and registers it as global trace provider.
func InitTracer(config configuration.Config) func() {
	log.Info().Msg("Initializing tracer...")
	if config.TraceConfig.Noop {
		log.Info().Msg("Using noop tracing!")
		otel.SetTracerProvider(trace.NewNoopTracerProvider())
		return func() {}
	}

	// return initGoogleTraceTracer(config)
	return initJaegerTracer(config)
}

func initJaegerTracer(config configuration.Config) func() {
	sampler := getSampler(config)

	// TODO what should this be?
	var jaegerHost string

	// Create and install Jaeger export pipeline
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint(fmt.Sprintf("http://%s:14268/api/traces", jaegerHost)),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: config.AppName,
			Tags: []label.KeyValue{
				label.String("exporter", "jaeger"),
				label.String("app_id", config.AppID),
			},
		}),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sampler}),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create Jaeger trace export pipeline")
	}

	return flush
}

// func initGoogleTraceTracer(config configuration.Config) func() {
// 	projectID := config.GCPProject

// 	sampler := getSampler(config)

// 	// Create Google Cloud Trace exporter to be able to retrieve the collected spans.
// 	// InstallNewPipeline instantiates a NewExportPipeline and registers it globally.
// 	_, flush, err := cloudtrace.InstallNewPipeline(
// 		[]cloudtrace.Option{cloudtrace.WithProjectID(projectID)},
// 		sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sampler}),
// 	)

// 	if err != nil {
// 		log.Fatal().Err(err).Msg("failed to create trace export pipeline")
// 	}

// 	return flush
// }

func getSampler(config configuration.Config) sdktrace.Sampler {
	var sampler sdktrace.Sampler
	if ratio := config.TraceConfig.SampleRatio; ratio == 1 {
		log.Info().Msg("Creating trace sampler with AlwaysSample")
		sampler = sdktrace.AlwaysSample()
	} else {
		log.Info().Msgf("Creating trace sampler with ratio: %v", config.TraceConfig.SampleRatio)
		sampler = sdktrace.TraceIDRatioBased(ratio)
	}
	return sampler
}
