package obs

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ToLogger(ctx context.Context) zerolog.Logger {
	// Create logger
	logContext := log.With()

	// TODO add requester fields as structured fields in logger
	//logContext = logContext.Str(key, val)

	return logContext.Logger()
}
