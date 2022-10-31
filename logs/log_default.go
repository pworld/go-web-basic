package logs

import (
	"go-web-platform/config"
	"log"
	"os"
	"strings"
)

func NewDefaultLog(cfg config.Config) Logger {

	var level = Debug
	if configLevelString, found := cfg.GetString("logging:level"); found {
		level = LogLevelFromString(configLevelString)
	}

	flags := log.Lmsgprefix | log.Ltime

	return &DefaultLog{
		minLevel: level,
		loggers: map[LevelLogging]*log.Logger{
			Trace: log.New(os.Stdout, "TRACE ", flags),
			Debug: log.New(os.Stdout, "DEBUG ", flags),
			Info:  log.New(os.Stdout, "INFO ", flags),
			Warn:  log.New(os.Stdout, "WARN ", flags),
			Fatal: log.New(os.Stdout, "FATAL ", flags),
		},
		triggerPanic: true,
	}
}

func LogLevelFromString(val string) (level LevelLogging) {
	switch strings.ToLower(val) {
	case "debug":
		level = Debug
	case "information":
		level = Info
	case "warning":
		level = Warn
	case "fatal":
		level = Fatal
	case "none":
		level = None
	default:
		level = Debug
	}
	return
}
