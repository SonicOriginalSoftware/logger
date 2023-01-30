package logger_test

import (
	"fmt"
	"strings"
	"testing"

	logger "git.sonicoriginal.software/logger"
)

const (
	prefix  = "Test"
	message = "Test error message"
)

func setup() *logger.Logger {
	writer := &strings.Builder{}

	return logger.New(prefix, logger.DefaultSeverity, writer)
}

func assert(channel string, writer fmt.Stringer) bool {
	expectedSuffix := fmt.Sprintf("[%v] [%v] %v\n", channel, prefix, message)
	receivedMessage := writer.String()

	return strings.HasSuffix(receivedMessage, expectedSuffix)
}

func TestDefaultLoggerError(t *testing.T) {
	testLogger := setup()

	testLogger.Error(message)

	if !assert(logger.ErrorPrefix, writer) {
		t.Errorf("'%v' != '%v'", receivedMessage, message)
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
