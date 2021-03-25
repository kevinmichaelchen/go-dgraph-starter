package obs

import (
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	flagForJaegerHost = "jaeger_host"
)

type JaegerConfig struct {
	Host string
}

func LoadJaegerConfig() JaegerConfig {
	c := JaegerConfig{
		Host: "jaeger",
	}

	flag.String(flagForJaegerHost, c.Host, "Jaeger host")

	flag.Parse()

	viper.BindPFlag(flagForJaegerHost, flag.Lookup(flagForJaegerHost))

	viper.AutomaticEnv()

	c.Host = viper.GetString(flagForJaegerHost)

	return c
}
