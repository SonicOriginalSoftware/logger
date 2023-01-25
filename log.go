//revive:disable:package-comments

package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const flags = log.Ldate | log.Ltime | log.Lmsgprefix

// DefaultSeverity shows error and warn log messages
const DefaultSeverity int64 = Error | Warn

// DefaultLogger is an unprefixed logger using the default severity
var DefaultLogger *Logger

const (
	Error = 1 << iota // Error shows error log messages
	Warn              // Warn shows warning log messages
	Info              // Info shows info log messages
	Debug             // Debug shows debug log messages
)

// Log defines a general logger
type Log interface {
	Info(format string, v ...any)
	Debug(format string, v ...any)
	Warn(format string, v ...any)
	Error(format string, v ...any)
}

func init() {
	severity := DefaultSeverity

	determineSeverity(&severity)

	DefaultLogger = New("", severity)
}

// Logger is used to log to appropriate levels
type Logger struct {
	warn  *log.Logger
	info  *log.Logger
	debug *log.Logger
	err   *log.Logger

	severity int64
}

func determineSeverity(severity *int64) {
	var level int64
	var err error

	rawLevel, defined := os.LookupEnv("LOGLEVEL")
	if defined {
		level, err = strconv.ParseInt(rawLevel, 10, 32)
	}
	if !defined || err != nil {
		level = DefaultSeverity
	}

	*severity = level
}

// New returns a valid instantiated logger
func New(prefix string, severity int64) *Logger {
	determineSeverity(&severity)

	return &Logger{
		warn:     new(prefix, "[WARN] ", os.Stdout),
		info:     new(prefix, "[INFO] ", os.Stdout),
		debug:    new(prefix, "[DEBUG] ", os.Stdout),
		err:      new(prefix, "[ERROR] ", os.Stderr),
		severity: severity,
	}
}

func new(prefix, defaultPrefix string, writer io.Writer) *log.Logger {
	if prefix != "" {
		defaultPrefix = fmt.Sprintf("%v[%v] ", defaultPrefix, prefix)
	}
	return log.New(writer, defaultPrefix, flags)
}

// Info a message
func (logger *Logger) Info(format string, v ...any) {
	if logger.severity&Info == 0 {
		return
	}
	logger.info.Printf(format, v...)
}

// Debug a message
func (logger *Logger) Debug(format string, v ...any) {
	if logger.severity&Debug == 0 {
		return
	}
	logger.debug.Printf(format, v...)
}

// Warn a message
func (logger *Logger) Warn(format string, v ...any) {
	if logger.severity&Warn == 0 {
		return
	}
	logger.warn.Printf(format, v...)
}

// Error a message
func (logger *Logger) Error(format string, v ...any) {
	if logger.severity&Error == 0 {
		return
	}
	logger.err.Printf(format, v...)
}
