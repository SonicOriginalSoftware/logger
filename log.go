//revive:disable:package-comments

package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Severity is an alias for an int64
type Severity = int64

type severityLabel string

const (
	Error Severity = 1 << iota // Error shows error log messages
	Warn                       // Warn shows warning log messages
	Info                       // Info shows info log messages
	Debug                      // Debug shows debug log messages
)

const (
	// DefaultSeverity shows error, warn, and info log messages
	DefaultSeverity Severity = Error | Warn | Info

	defaultFlags = log.Ldate | log.Ltime | log.Lmsgprefix
)

const (
	// ChannelLabelError is the prefix used for error messages
	ChannelLabelError string = "ERROR"
	// ChannelLabelWarn is the prefix used for warning messages
	ChannelLabelWarn string = "WARN"
	// ChannelLabelInfo is the prefix used for informational messages
	ChannelLabelInfo string = "INFO"
	// ChannelLabelDebug is the prefix used for debug messages
	ChannelLabelDebug string = "DEBUG"
)

const (
	// LogLevelDefault defines an alias for a default log severity
	LogLevelDefault severityLabel = "LOG_LEVEL_DEFAULT"

	// LogLevelError is the variable for defining the state of the error channel
	LogLevelError severityLabel = "LOG_LEVEL_ERROR"
	// LogLevelWarn is the variable for defining the state of the warn channel
	LogLevelWarn severityLabel = "LOG_LEVEL_WARN"
	// LogLevelInfo is the variable for defining the state of the info channel
	LogLevelInfo severityLabel = "LOG_LEVEL_INFO"
	// LogLevelDebug is the variable for defining the state of the debug channel
	LogLevelDebug severityLabel = "LOG_LEVEL_DEBUG"
)

var severityMap = map[severityLabel]Severity{
	LogLevelError: Error,
	LogLevelWarn:  Warn,
	LogLevelInfo:  Info,
	LogLevelDebug: Debug,
}

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

func (logger *Logger) handleLogLevel(logLevel severityLabel) {
	state, defined := os.LookupEnv(string(logLevel))
	if !defined {
		return
	}

	if state == "0" {
		// Disable that channel
		logger.Severity &= ^severityMap[logLevel]
	} else {
		// Enable that channel
		logger.Severity |= severityMap[logLevel]
	}
}

func (logger *Logger) determineSeverity() {
	logger.handleLogLevel(LogLevelDefault)
	logger.handleLogLevel(LogLevelError)
	logger.handleLogLevel(LogLevelWarn)
	logger.handleLogLevel(LogLevelInfo)
	logger.handleLogLevel(LogLevelDebug)
}

func new(prefix, defaultPrefix string, writer io.Writer) *log.Logger {
	if prefix != "" {
		prefix = fmt.Sprintf("%v[%v] ", defaultPrefix, prefix)
	}
	return log.New(writer, prefix, defaultFlags)
}

// New returns a valid logger ready for use
func New(prefix string, severity Severity, writer io.Writer) (logger *Logger) {
	logger = &Logger{
		warn:     new(prefix, fmt.Sprintf("[%v] ", ChannelLabelWarn), writer),
		info:     new(prefix, fmt.Sprintf("[%v] ", ChannelLabelInfo), writer),
		debug:    new(prefix, fmt.Sprintf("[%v] ", ChannelLabelDebug), writer),
		err:      new(prefix, fmt.Sprintf("[%v] ", ChannelLabelError), writer),
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
