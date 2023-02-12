package logger_test

import (
	"fmt"
	"strings"
	"testing"

	logger "git.sonicoriginal.software/logger"
)

const (
	prefix           = "Test"
	testMessageError = "Test error message"
	testMessageWarn  = "Test warn message"
	testMessageInfo  = "Test info message"
	testMessageDebug = "Test debug message"

	enabledValue      = "1"
	undeterminedValue = "-1"
	bogusValue        = "bogus"
	disabledValue     = "0"
)

type callback func(f string, v ...any)

func prepare(severity logger.Severity) (writer *strings.Builder, testLogger *logger.Logger) {
	writer = &strings.Builder{}
	testLogger = logger.New(prefix, severity, writer, writer)
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

func testError(t *testing.T, loggerSeverity logger.Severity, channel logger.Severity, channelLabel, prefix, message string) {
	writer, testLogger := prepare(loggerSeverity)
	testFunction := testLogger.Error

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func testWarn(t *testing.T, loggerSeverity logger.Severity, channel logger.Severity, channelLabel, prefix, message string) {
	writer, testLogger := prepare(loggerSeverity)
	testFunction := testLogger.Warn

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func testInfo(t *testing.T, loggerSeverity logger.Severity, channel logger.Severity, channelLabel, prefix, message string) {
	writer, testLogger := prepare(loggerSeverity)
	testFunction := testLogger.Info

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func testDebug(t *testing.T, loggerSeverity logger.Severity, channel logger.Severity, channelLabel, prefix, message string) {
	writer, testLogger := prepare(loggerSeverity)
	testFunction := testLogger.Debug

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func logLevelError(t *testing.T, channel logger.Severity, channelLabel, prefix, message string) {
	writer, testLogger := prepare(logger.DefaultSeverity)
	testFunction := testLogger.Error

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func logLevelWarn(t *testing.T, channel logger.Severity, channelLabel, prefix, message string) {
	writer, testLogger := prepare(logger.DefaultSeverity)
	testFunction := testLogger.Warn

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func logLevelInfo(t *testing.T, channel logger.Severity, channelLabel, prefix, message string) {
	writer, testLogger := prepare(logger.DefaultSeverity)
	testFunction := testLogger.Info

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func logLevelDebug(t *testing.T, channel logger.Severity, channelLabel, prefix, message string) {
	writer, testLogger := prepare(logger.DefaultSeverity)
	testFunction := testLogger.Debug

	runTest(t, testLogger, testFunction, channel, message, channelLabel, writer)
}

func TestError(t *testing.T) {
	channel := logger.Error
	channelLabel := logger.ChannelLabelError
	message := testMessageError

	t.Run("Default Error", func(t *testing.T) {
		testError(t, logger.DefaultSeverity, channel, channelLabel, prefix, message)
	})
	t.Run("Only Error", func(t *testing.T) {
		testError(t, logger.Error, channel, channelLabel, prefix, message)
	})

	t.Run("Except Error", func(t *testing.T) {
		testError(t, logger.Warn|logger.Info|logger.Debug, channel, channelLabel, prefix, message)
	})
}

func TestWarn(t *testing.T) {
	channel := logger.Warn
	channelLabel := logger.ChannelLabelWarn
	message := testMessageWarn

	t.Run("Default Warn", func(t *testing.T) {
		testWarn(t, logger.DefaultSeverity, channel, channelLabel, prefix, message)
	})
	t.Run("Only Warn", func(t *testing.T) {
		testWarn(t, logger.Warn, channel, channelLabel, prefix, message)
	})
	t.Run("Except Warn", func(t *testing.T) {
		testWarn(t, logger.Error|logger.Info|logger.Debug, channel, channelLabel, prefix, message)
	})
}

func TestInfo(t *testing.T) {
	channel := logger.Info
	channelLabel := logger.ChannelLabelInfo
	message := testMessageInfo

	t.Run("Default Info", func(t *testing.T) {
		testInfo(t, logger.DefaultSeverity, channel, channelLabel, prefix, message)
	})
	t.Run("Only Info", func(t *testing.T) {
		testInfo(t, logger.Info, channel, channelLabel, prefix, message)
	})
	t.Run("Except Info", func(t *testing.T) {
		testInfo(t, logger.Error|logger.Warn|logger.Debug, channel, channelLabel, prefix, message)
	})
}

func TestDebug(t *testing.T) {
	channel := logger.Debug
	channelLabel := logger.ChannelLabelDebug
	message := testMessageDebug

	t.Run("Default Debug", func(t *testing.T) {
		testDebug(t, logger.DefaultSeverity, channel, channelLabel, prefix, message)
	})
	t.Run("Only Debug", func(t *testing.T) {
		testDebug(t, logger.Debug, channel, channelLabel, prefix, message)
	})
	t.Run("Except Debug", func(t *testing.T) {
		testDebug(t, logger.Error|logger.Warn|logger.Info, channel, channelLabel, prefix, message)
	})
}

func TestLogLevelError(t *testing.T) {
	channel := logger.Error
	channelLabel := logger.ChannelLabelError
	message := testMessageError

	t.Run("Default Severity Error LogLevel Disabled", func(t *testing.T) {
		t.Setenv(string(logger.LogLevelWarn), disabledValue)
		logLevelError(t, channel, channelLabel, prefix, message)
	})

	t.Run("Default Severity Error LogLevel Enabled", func(t *testing.T) {
		t.Setenv(string(logger.LogLevelWarn), enabledValue)
		logLevelError(t, channel, channelLabel, prefix, message)
	})

	t.Run("Default Severity Error LogLevel Undetermined", func(t *testing.T) {
		t.Setenv(string(logger.LogLevelWarn), undeterminedValue)
		logLevelError(t, channel, channelLabel, prefix, message)
	})

	t.Run("Default Severity Error LogLevel Bogus", func(t *testing.T) {
		t.Setenv(string(logger.LogLevelWarn), bogusValue)
		logLevelError(t, channel, channelLabel, prefix, message)
	})
}

func TestLogLevelWarn(t *testing.T) {
	channel := logger.Warn
	channelLabel := logger.ChannelLabelWarn
	message := testMessageWarn

	t.Run("Default Severity Warn LogLevel Disabled", func(t *testing.T) {
		t.Setenv(string(logger.LogLevelWarn), disabledValue)
		logLevelWarn(t, channel, channelLabel, prefix, message)
	})

	t.Run("Default Severity Warn LogLevel Enabled", func(t *testing.T) {
		t.Setenv(string(logger.LogLevelWarn), enabledValue)
		logLevelWarn(t, channel, channelLabel, prefix, message)
	})

	t.Run("Default Severity Warn LogLevel Undetermined", func(t *testing.T) {
		t.Setenv(string(logger.LogLevelWarn), undeterminedValue)
		logLevelWarn(t, channel, channelLabel, prefix, message)
	})

	t.Run("Default Severity Warn LogLevel Bogus", func(t *testing.T) {
		t.Setenv(string(logger.LogLevelWarn), bogusValue)
		logLevelWarn(t, channel, channelLabel, prefix, message)
	})
}

func TestLogLevelInfo(t *testing.T) {
	channel := logger.Info
	channelLabel := logger.ChannelLabelInfo
	message := testMessageInfo

	t.Run("Default Severity Info LogLevel Disabled", func(t *testing.T) {
		t.Setenv(string(logger.LogLevelInfo), disabledValue)
		logLevelInfo(t, channel, channelLabel, prefix, message)
	})

	t.Run("Default Severity Info LogLevel Enabled", func(t *testing.T) {
		t.Setenv(string(logger.LogLevelInfo), enabledValue)
		logLevelInfo(t, channel, channelLabel, prefix, message)
	})

	t.Run("Default Severity Info LogLevel Undetermined", func(t *testing.T) {
		t.Setenv(string(logger.LogLevelInfo), undeterminedValue)
		logLevelInfo(t, channel, channelLabel, prefix, message)
	})

	t.Run("Default Severity Info LogLevel Bogus", func(t *testing.T) {
		t.Setenv(string(logger.LogLevelInfo), bogusValue)
		logLevelInfo(t, channel, channelLabel, prefix, message)
	})
}

func TestLogLevelDebug(t *testing.T) {
	channel := logger.Debug
	channelLabel := logger.ChannelLabelDebug
	message := testMessageDebug

	t.Run("Default Severity Debug LogLevel Disabled", func(t *testing.T) {
		t.Setenv(string(logger.LogLevelDebug), disabledValue)
		logLevelDebug(t, channel, channelLabel, prefix, message)
	})

	t.Run("Default Severity Debug LogLevel Enabled", func(t *testing.T) {
		t.Setenv(string(logger.LogLevelDebug), enabledValue)
		logLevelDebug(t, channel, channelLabel, prefix, message)
	})

	t.Run("Default Severity Debug LogLevel Undetermined", func(t *testing.T) {
		t.Setenv(string(logger.LogLevelDebug), undeterminedValue)
		logLevelDebug(t, channel, channelLabel, prefix, message)
	})

	t.Run("Default Severity Debug LogLevel Bogus", func(t *testing.T) {
		t.Setenv(string(logger.LogLevelDebug), bogusValue)
		logLevelDebug(t, channel, channelLabel, prefix, message)
	})
}
