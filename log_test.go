package logger_test

import (
	"testing"

	logger "git.sonicoriginal.software/logger"
)

func TestDefaultLoggerError(t *testing.T) {
	message := "Test error message"

	logger.DefaultLogger.Error(message)

	receivedMessage := ""

	if receivedMessage != message {
		t.Logf(
			"Piped message not equal to test message: '%v' != '%v'",
			receivedMessage,
			message,
		)
		t.FailNow()
	}
}

func TestDefaultLoggerWarn(t *testing.T) {
	logger.DefaultLogger.Warn("Test warn message")
}

func TestDefaultLoggerInfo(t *testing.T) {
	logger.DefaultLogger.Info("Test info message")
}

func TestDefaultLoggerDebug(t *testing.T) {
	logger.DefaultLogger.Debug("Test debug message")
}
