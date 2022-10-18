package log

type LevelLogging int

const (
	// iota constant increasing sequences
	Trace LevelLogging = iota
	Debug
	Info
	Warn
	Fatal
	none
)

// Interface for different level severity
type DataLogger interface {
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
