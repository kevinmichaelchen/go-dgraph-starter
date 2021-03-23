package obs

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	flag "github.com/spf13/pflag"

	"github.com/spf13/viper"
)

const (
	flagForPrettyLogs = "log_pretty"
	flagForLogLevel   = "log_level"
)

type LoggingConfig struct {
	Pretty bool
	Level  string
}

func LoadLoggingConfig() LoggingConfig {
	c := LoggingConfig{
		Pretty: true,
		Level:  zerolog.DebugLevel.String(),
	}

	flag.Bool(flagForPrettyLogs, c.Pretty, "Pretty logs")
	flag.String(flagForLogLevel, c.Level, "Log level")

	flag.Parse()

	viper.BindPFlag(flagForPrettyLogs, flag.Lookup(flagForPrettyLogs))
	viper.BindPFlag(flagForLogLevel, flag.Lookup(flagForLogLevel))

	viper.AutomaticEnv()

	ConfigureLogger(c)

	return c
}

func ConfigureLogger(c LoggingConfig) {
	// UNIX Time is faster and smaller than most timestamps
	// If you set zerolog.TimeFieldFormat to an empty string,
	// logs will write with UNIX time
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if level, err := zerolog.ParseLevel(c.Level); err != nil {
		log.Fatal().Err(err).Msg("Failed to parse log level")
	} else {
		zerolog.SetGlobalLevel(level)
	}

	if c.Pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}

func GetLogger() zerolog.Logger {
	return log.Logger
}
