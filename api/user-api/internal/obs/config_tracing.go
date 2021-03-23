package obs

import (
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	flagForTraceSampleRatio = "trace_sample_ratio"
	flagForTraceNoop        = "trace_noop"
)

type TraceConfig struct {
	SampleRatio float64
	Noop        bool
}

func LoadTraceConfig() TraceConfig {
	c := TraceConfig{
		SampleRatio: 1,
	}

	flag.Float64(flagForTraceSampleRatio, c.SampleRatio, "Probabilistic sample ratio")
	flag.Bool(flagForTraceNoop, c.Noop, "Whether tracing uses noops or not")

	flag.Parse()

	viper.BindPFlag(flagForTraceSampleRatio, flag.Lookup(flagForTraceSampleRatio))
	viper.BindPFlag(flagForTraceNoop, flag.Lookup(flagForTraceNoop))

	viper.AutomaticEnv()

	c.SampleRatio = viper.GetFloat64(flagForTraceSampleRatio)
	c.Noop = viper.GetBool(flagForTraceNoop)

	return c
}
