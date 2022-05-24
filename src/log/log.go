package log

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger zerolog.Logger

func Init(logLevel string) {
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		log.Fatal().Err(err).Msgf("Fatal: failed to parse log level: %s", logLevel)
	}
	if level > zerolog.FatalLevel {
		log.Fatal().Msgf("zerolog level cannot be set to '%s', ending app here..."+
			" (tip: did you forget to set log level?)", level)
	}

	zerolog.SetGlobalLevel(level)

	// Shorten keys
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "m"
	Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
}
