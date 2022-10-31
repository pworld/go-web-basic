package logs

type LevelLogging int

const (
	// iota constant increasing sequences
	Trace LevelLogging = iota
	Debug
	Info
	Warn
	Fatal
	None
)

// Interface for different level severity
type Logger interface {
	Trace(string)
	Tracef(string, ...interface{})

	Info(string)
	Infof(string, ...interface{})

	Debug(string)
	Debugf(string, ...interface{})

	Warn(string)
	Warnf(string, ...interface{})

	Panic(string)
	Panicf(string, ...interface{})
}
