package log

import (
	"log"
	"os"
)

func NewDefaultLog(level LevelLogging) DataLogger {
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
