package configuration

import (
	"encoding/json"

	"github.com/rs/zerolog/log"

	"github.com/rs/xid"

	flag "github.com/spf13/pflag"

	"github.com/spf13/viper"
)

const (
	flagForGrpcPort = "grpc_port"
	flagForHTTPPort = "http_port"
)

type Config struct {
	// AppName is a low cardinality identifier for this service.
	AppName string

	// AppID is a unique identifier for the instance (pod) running this app.
	AppID string

	// GrpcPort controls what port our gRPC server runs on.
	GrpcPort int

	// HTTPPort controls what port our HTTP server runs on.
	HTTPPort int

	SQLConfig    SQLConfig
	RedisConfig  RedisConfig
	DgraphConfig DgraphConfig

	// TraceConfig contains config info for how we do tracing.
	TraceConfig TraceConfig
}

func (c Config) String() string {
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not marshal config to string")
	}
	return string(b)
}

func LoadConfig() Config {
	c := Config{
		AppName:  "api-monorepo",
		AppID:    xid.New().String(),
		GrpcPort: 8084,
		HTTPPort: 8085,
	}

	c.SQLConfig = LoadSQLConfig()
	c.RedisConfig = LoadRedisConfig()
	c.DgraphConfig = LoadDgraphConfig()
	c.TraceConfig = LoadTraceConfig()

	flag.Int(flagForGrpcPort, c.GrpcPort, "gRPC port")
	flag.Int(flagForHTTPPort, c.HTTPPort, "HTTP port")

	flag.Parse()

	viper.BindPFlag(flagForGrpcPort, flag.Lookup(flagForGrpcPort))
	viper.BindPFlag(flagForHTTPPort, flag.Lookup(flagForHTTPPort))

	viper.AutomaticEnv()

	c.GrpcPort = viper.GetInt(flagForGrpcPort)
	c.HTTPPort = viper.GetInt(flagForHTTPPort)

	return c
}
