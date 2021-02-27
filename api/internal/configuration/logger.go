package configuration

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func GetLogger() zerolog.Logger {
	return log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
