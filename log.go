//revive:disable:package-comments

package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Severity is an alias for an int64
type Severity = int64

const (
	Error Severity = 1 << iota // Error shows error log messages
	Warn                       // Warn shows warning log messages
	Info                       // Info shows info log messages
	Debug                      // Debug shows debug log messages

	// ErrorChannelLabel is the prefix used for error messages
	ErrorChannelLabel = "ERROR"
	// WarnChannelLabel is the prefix used for warning messages
	WarnChannelLabel = "WARN"
	// InfoChannelLabel is the prefix used for informational messages
	InfoChannelLabel = "INFO"
	// DebugChannelLabel is the prefix used for debug messages
	DebugChannelLabel = "DEBUG"

	flags = log.Ldate | log.Ltime | log.Lmsgprefix

	// DefaultSeverity shows error, warn, and info log messages
	DefaultSeverity Severity = Error | Warn | Info
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

	Severity Severity
}

func new(prefix, defaultPrefix string, writer io.Writer) *log.Logger {
	if prefix != "" {
		prefix = fmt.Sprintf("%v[%v] ", defaultPrefix, prefix)
	}
	return log.New(writer, prefix, flags)
}

func (logger *Logger) determineSeverity() {
	var level Severity
	var err error

	rawLevel, defined := os.LookupEnv("LOGLEVEL")
	if defined {
		level, err = strconv.ParseInt(rawLevel, 10, 32)
	}
	if !defined || err != nil {
		level = logger.Severity
	}

	logger.Severity = level
}

// New returns a valid logger ready for use
func New(prefix string, severity Severity, writer io.Writer) (logger *Logger) {
	logger = &Logger{
		warn:     new(prefix, fmt.Sprintf("[%v] ", WarnChannelLabel), writer),
		info:     new(prefix, fmt.Sprintf("[%v] ", InfoChannelLabel), writer),
		debug:    new(prefix, fmt.Sprintf("[%v] ", DebugChannelLabel), writer),
		err:      new(prefix, fmt.Sprintf("[%v] ", ErrorChannelLabel), writer),
		Severity: severity,
	}

	logger.determineSeverity()

	return
}

// ChannelEnabled returns whether the severity is enabled (prints to the log)
func (logger *Logger) ChannelEnabled(channel Severity) bool {
	return logger.Severity&channel != 0
}

// Info a message
func (logger *Logger) Info(format string, v ...any) {
	if !logger.ChannelEnabled(Info) {
		return
	}
	logger.info.Printf(format, v...)
}

// Debug a message
func (logger *Logger) Debug(format string, v ...any) {
	if !logger.ChannelEnabled(Debug) {
		return
	}
	logger.debug.Printf(format, v...)
}

// Warn a message
func (logger *Logger) Warn(format string, v ...any) {
	if !logger.ChannelEnabled(Warn) {
		return
	}
	logger.warn.Printf(format, v...)
}

// Error a message
func (logger *Logger) Error(format string, v ...any) {
	if !logger.ChannelEnabled(Error) {
		return
	}
	logger.err.Printf(format, v...)
}
