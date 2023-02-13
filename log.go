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

const (
	Error Severity = 1 << iota // Error shows error log messages
	Warn                       // Warn shows warning log messages
	Info                       // Info shows info log messages
	Debug                      // Debug shows debug log messages

	// DefaultSeverity shows error, warn, and info log messages
	DefaultSeverity Severity = Error | Warn | Info
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

var severityMap = map[string]Severity{
	ChannelLabelError: Error,
	ChannelLabelWarn:  Warn,
	ChannelLabelInfo:  Info,
	ChannelLabelDebug: Debug,
}

// DefaultLogger is an unprefixed logger using the default severity
var DefaultLogger = New("", DefaultSeverity, os.Stdout, os.Stderr)

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

func (logger *Logger) setLoggerChannelState(logLevel, state string) {
	channel, found := severityMap[logLevel]

	if !found {
		return
	}

	if state == "0" {
		// Disable that channel
		logger.Severity &= ^channel
	} else {
		// Enable that channel
		logger.Severity |= channel
	}
}

func (logger *Logger) handleLogLevel(prefix, logLevel string) {
	envVariable := fmt.Sprintf("%v_LOG_LEVEL_%v", prefix, logLevel)
	state, defined := os.LookupEnv(envVariable)

	if !defined {
		envVariable = fmt.Sprintf("LOG_LEVEL_%v", logLevel)
		state, defined = os.LookupEnv(envVariable)
	}

	if !defined {
		return
	}

	logger.setLoggerChannelState(logLevel, state)
}

func new(prefix, defaultPrefix string, writer io.Writer) *log.Logger {
	const defaultFlags = log.Ldate | log.Ltime | log.Lmsgprefix

	if prefix != "" {
		prefix = fmt.Sprintf("%v[%v] ", defaultPrefix, prefix)
	}

	return log.New(writer, prefix, defaultFlags)
}

// New returns a valid logger ready for use
func New(prefix string, severity Severity, stdoutWriter io.Writer, stderrWriter io.Writer) (logger *Logger) {
	logger = &Logger{
		warn:     new(prefix, fmt.Sprintf("[%v] ", ChannelLabelWarn), stdoutWriter),
		info:     new(prefix, fmt.Sprintf("[%v] ", ChannelLabelInfo), stdoutWriter),
		debug:    new(prefix, fmt.Sprintf("[%v] ", ChannelLabelDebug), stdoutWriter),
		err:      new(prefix, fmt.Sprintf("[%v] ", ChannelLabelError), stderrWriter),
		Severity: severity,
	}

	logger.handleLogLevel(prefix, ChannelLabelError)
	logger.handleLogLevel(prefix, ChannelLabelWarn)
	logger.handleLogLevel(prefix, ChannelLabelInfo)
	logger.handleLogLevel(prefix, ChannelLabelDebug)

	return
}

// ChannelEnabled returns whether the severity is enabled (prints to the log)
func (logger *Logger) ChannelEnabled(channel Severity) bool {
	return logger.Severity&channel != 0
}

// Debug a message
func (logger *Logger) Debug(format string, v ...any) {
	if !logger.ChannelEnabled(Debug) {
		return
	}
	logger.debug.Printf(format, v...)
}

// Info a message
func (logger *Logger) Info(format string, v ...any) {
	if !logger.ChannelEnabled(Info) {
		return
	}
	logger.info.Printf(format, v...)
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
