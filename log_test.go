package logger_test

import (
	"fmt"
	"strings"
	"testing"

	logger "git.sonicoriginal.software/logger"
)

const (
	prefix           = "Test"
	testErrorMessage = "Test error message"
	testWarnMessage  = "Test warn message"
	testInfoMessage  = "Test info message"
	testDebugMessage = "Test debug message"
)

type callback func(f string, v ...any)

func prepare(severity logger.Severity) (writer *strings.Builder, testLogger *logger.Logger) {
	writer = &strings.Builder{}
	testLogger = logger.New(prefix, severity, writer)
	return
}

func runTest(
	t *testing.T,
	testLogger *logger.Logger,
	testCall callback,
	channel logger.Severity,
	message, channelLabel string,
	writer *strings.Builder,
) {
	testCall(message)

	if !testLogger.ChannelEnabled(channel) {
		return
	}

	receivedMessage := writer.String()
	expectedSuffix := fmt.Sprintf("[%v] [%v] %v\n", channelLabel, prefix, message)

	if !strings.HasSuffix(receivedMessage, expectedSuffix) {
		t.Errorf("'%v' != '%v'", message, receivedMessage)
	}
}

func defaultError(t *testing.T, channel logger.Severity, channelLabel, prefix, message string) {
	loggerSeverity := logger.DefaultSeverity
	writer, testLogger := prepare(loggerSeverity)
	testFunction := testLogger.Error

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func onlyError(t *testing.T, channel logger.Severity, channelLabel, prefix, message string) {
	severity := logger.Error
	writer, testLogger := prepare(severity)
	testFunction := testLogger.Error

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func exceptError(t *testing.T, channel logger.Severity, channelLabel, prefix, message string) {
	severity := logger.Warn | logger.Info | logger.Debug
	writer, testLogger := prepare(severity)
	testFunction := testLogger.Debug

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func defaultWarn(t *testing.T, channel logger.Severity, channelLabel, prefix, message string) {
	severity := logger.DefaultSeverity
	writer, testLogger := prepare(severity)
	testFunction := testLogger.Warn

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func onlyWarn(t *testing.T, channel logger.Severity, channelLabel, prefix, message string) {
	severity := logger.Warn
	writer, testLogger := prepare(severity)
	testFunction := testLogger.Warn

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func exceptWarn(t *testing.T, channel logger.Severity, channelLabel, prefix, message string) {
	severity := logger.Error | logger.Info | logger.Debug
	writer, testLogger := prepare(severity)
	testFunction := testLogger.Debug

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func defaultInfo(t *testing.T, channel logger.Severity, channelLabel, prefix, message string) {
	severity := logger.DefaultSeverity
	writer, testLogger := prepare(severity)
	testFunction := testLogger.Info

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func onlyInfo(t *testing.T, channel logger.Severity, channelLabel, prefix, message string) {
	severity := logger.Info
	writer, testLogger := prepare(severity)
	testFunction := testLogger.Info

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func exceptInfo(t *testing.T, channel logger.Severity, channelLabel, prefix, message string) {
	severity := logger.Error | logger.Warn | logger.Debug
	writer, testLogger := prepare(severity)
	testFunction := testLogger.Debug

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func defaultDebug(t *testing.T, channel logger.Severity, channelLabel, prefix, message string) {
	severity := logger.DefaultSeverity
	writer, testLogger := prepare(severity)
	testFunction := testLogger.Debug

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func onlyDebug(t *testing.T, channel logger.Severity, channelLabel, prefix, message string) {
	severity := logger.Debug
	writer, testLogger := prepare(severity)
	testFunction := testLogger.Debug

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func exceptDebug(t *testing.T, channel logger.Severity, channelLabel, prefix, message string) {
	severity := logger.Error | logger.Warn | logger.Info
	writer, testLogger := prepare(severity)
	testFunction := testLogger.Debug

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func TestError(t *testing.T) {
	channel := logger.Error
	message := testErrorMessage
	channelLabel := logger.ErrorChannelLabel

	t.Run("Default Error", func(t *testing.T) {
		defaultError(t, channel, channelLabel, prefix, message)
	})
	t.Run("Only Error", func(t *testing.T) {
		onlyError(t, channel, channelLabel, prefix, message)
	})

	t.Run("Except Error", func(t *testing.T) {
		exceptError(t, channel, channelLabel, prefix, message)
	})
}

func TestWarn(t *testing.T) {
	channel := logger.Error
	message := testWarnMessage
	channelLabel := logger.WarnChannelLabel

	t.Run("Default Warn", func(t *testing.T) {
		defaultWarn(t, channel, channelLabel, prefix, message)
	})
	t.Run("Only Warn", func(t *testing.T) {
		onlyWarn(t, channel, channelLabel, prefix, message)
	})
	t.Run("Except Warn", func(t *testing.T) {
		exceptWarn(t, channel, channelLabel, prefix, message)
	})
}

func TestInfo(t *testing.T) {
	channel := logger.Info
	message := testInfoMessage
	channelLabel := logger.InfoChannelLabel

	t.Run("Default Info", func(t *testing.T) {
		defaultInfo(t, channel, channelLabel, prefix, message)
	})
	t.Run("Only Info", func(t *testing.T) {
		onlyInfo(t, channel, channelLabel, prefix, message)
	})
	t.Run("Except Info", func(t *testing.T) {
		exceptInfo(t, channel, channelLabel, prefix, message)
	})
}

func TestDebug(t *testing.T) {
	channel := logger.Debug
	message := testDebugMessage
	channelLabel := logger.DebugChannelLabel

	t.Run("Default Debug", func(t *testing.T) {
		defaultDebug(t, channel, channelLabel, prefix, message)
	})
	t.Run("Only Debug", func(t *testing.T) {
		onlyDebug(t, channel, channelLabel, prefix, message)
	})
	t.Run("Except Debug", func(t *testing.T) {
		exceptDebug(t, channel, channelLabel, prefix, message)
	})
}
