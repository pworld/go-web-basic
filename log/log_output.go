package log

import (
	"fmt"
	"log"
)

// DefaultLog Implement Log interface using the features provided by the log package in the standard library
type DefaultLog struct {
	minLevel     LevelLogging
	loggers      map[LevelLogging]*log.Logger
	triggerPanic bool
}

func (l *DefaultLog) MinLogLevel() LevelLogging {
	return l.minLevel
}

func (l *DefaultLog) write(level LevelLogging, message string) {
	if l.minLevel <= level {
		l.loggers[level].Output(2, message)
	}
}

func (l *DefaultLog) Trace(message string) {
	l.write(Trace, message)
}

func (l *DefaultLog) Tracef(template string, vals ...interface{}) {
	l.write(Trace, fmt.Sprintf(template, vals...))
}

func (l *DefaultLog) Info(message string) {
	l.write(Info, message)
}

func (l *DefaultLog) Infof(template string, vals ...interface{}) {
	l.write(Info, fmt.Sprintf(template, vals...))
}

func (l *DefaultLog) Debug(message string) {
	l.write(Debug, message)
}

func (l *DefaultLog) Debugf(template string, vals ...interface{}) {
	l.write(Debug, fmt.Sprintf(template, vals...))
}

func (l *DefaultLog) Warn(message string) {
	l.write(Warn, message)
}

func (l *DefaultLog) Warnf(template string, vals ...interface{}) {
	l.write(Warn, fmt.Sprintf(template, vals...))
}

func (l *DefaultLog) Panic(msg string) {
	l.write(Fatal, msg)
	if l.triggerPanic {
		panic(msg)
	}

}
func (l *DefaultLog) Panicf(template string, vals ...interface{}) {
	formattedMsg := fmt.Sprintf(template, vals...)
	l.write(Fatal, formattedMsg)
	if l.triggerPanic {
		panic(formattedMsg)
	}
}
