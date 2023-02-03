//revive:disable:package-comments

package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const (
	Error = 1 << iota // Error shows error log messages
	Warn              // Warn shows warning log messages
	Info              // Info shows info log messages
	Debug             // Debug shows debug log messages

	// ErrorPrefix is the prefix used for error messages
	ErrorPrefix = "ERROR"
	// WarnPrefix is the prefix used for warning messages
	WarnPrefix = "WARN"
	// InfoPrefix is the prefix used for informational messages
	InfoPrefix = "INFO"
	// DebugPrefix is the prefix used for debug messages
	DebugPrefix = "DEBUG"

	flags = log.Ldate | log.Ltime | log.Lmsgprefix

	// DefaultSeverity shows error, warn, and info log messages
	DefaultSeverity int64 = Error | Warn | Info
)

// DefaultLogger is an unprefixed logger using the default severity
var DefaultLogger = New("", DefaultSeverity, os.Stdout)

// Log defines a general logger
type Log interface {
	Info(format string, v ...any)
	Debug(format string, v ...any)
	Warn(format string, v ...any)
	Error(format string, v ...any)
}

// Logger is used to log to appropriate levels
type Logger struct {
	warn  *log.Logger
	info  *log.Logger
	debug *log.Logger
	err   *log.Logger

	severity int64
}

func new(prefix, defaultPrefix string, writer io.Writer) *log.Logger {
	if prefix != "" {
		defaultPrefix = fmt.Sprintf("%v[%v] ", defaultPrefix, prefix)
	}
	return log.New(writer, defaultPrefix, flags)
}

// New returns a valid logger ready for use
func New(prefix string, severity int64, writer io.Writer) (logger *Logger) {
	logger = &Logger{
		warn:     new(prefix, fmt.Sprintf("[%v] ", WarnPrefix), writer),
		info:     new(prefix, fmt.Sprintf("[%v] ", InfoPrefix), writer),
		debug:    new(prefix, fmt.Sprintf("[%v] ", DebugPrefix), writer),
		err:      new(prefix, fmt.Sprintf("[%v] ", ErrorPrefix), writer),
		severity: severity,
	}

	logger.determineSeverity()

	return
}

func (logger *Logger) determineSeverity() {
	var level int64
	var err error

	rawLevel, defined := os.LookupEnv("LOGLEVEL")
	if defined {
		level, err = strconv.ParseInt(rawLevel, 10, 32)
	}
	if !defined || err != nil {
		level = DefaultSeverity
	}

	logger.severity = level
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
