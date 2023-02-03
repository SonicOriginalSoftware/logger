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

func assert(t *testing.T, testCall callback, message, channel string, writer *strings.Builder) {
	testCall(message)

	expectedSuffix := fmt.Sprintf("[%v] [%v] %v\n", channel, prefix, message)
	receivedMessage := writer.String()
	matches := strings.HasSuffix(receivedMessage, expectedSuffix)

	if !matches {
		t.Errorf("'%v' != '%v'", receivedMessage, message)
	}
}

func prepare(severity int64) (writer *strings.Builder, testLogger *logger.Logger) {
	writer = &strings.Builder{}
	testLogger = logger.New(prefix, severity, writer)
	return
}

func defaultError(t *testing.T) {
	writer, testLogger := prepare(logger.DefaultSeverity)
	assert(t, testLogger.Error, testErrorMessage, logger.ErrorPrefix, writer)
}

func onlyError(t *testing.T) {
	writer, testLogger := prepare(logger.Error)
	assert(t, testLogger.Error, testErrorMessage, logger.ErrorPrefix, writer)
}

func exceptError(t *testing.T) {
	writer, testLogger := prepare(logger.Warn | logger.Info | logger.Debug)
	assert(t, testLogger.Error, testErrorMessage, logger.ErrorPrefix, writer)
}

func defaultWarn(t *testing.T) {
	writer, testLogger := prepare(logger.DefaultSeverity)
	assert(t, testLogger.Warn, testWarnMessage, logger.WarnPrefix, writer)
}

func onlyWarn(t *testing.T) {
	writer, testLogger := prepare(logger.Warn)
	assert(t, testLogger.Warn, testWarnMessage, logger.WarnPrefix, writer)
}

func exceptWarn(t *testing.T) {
	writer, testLogger := prepare(logger.Error | logger.Info | logger.Debug)
	assert(t, testLogger.Warn, testWarnMessage, logger.WarnPrefix, writer)
}

func defaultInfo(t *testing.T) {
	writer, testLogger := prepare(logger.DefaultSeverity)
	assert(t, testLogger.Info, testInfoMessage, logger.InfoPrefix, writer)
}

func onlyInfo(t *testing.T) {
	writer, testLogger := prepare(logger.Info)
	assert(t, testLogger.Info, testInfoMessage, logger.InfoPrefix, writer)
}

func exceptInfo(t *testing.T) {
	writer, testLogger := prepare(logger.Error | logger.Warn | logger.Debug)
	assert(t, testLogger.Info, testInfoMessage, logger.InfoPrefix, writer)
}

func defaultDebug(t *testing.T) {
	writer, testLogger := prepare(logger.DefaultSeverity)
	assert(t, testLogger.Debug, "Test debug message", logger.DebugPrefix, writer)
}

func onlyDebug(t *testing.T) {
	writer, testLogger := prepare(logger.Debug)
	assert(t, testLogger.Debug, "Test debug message", logger.DebugPrefix, writer)
}

func exceptDebug(t *testing.T) {
	writer, testLogger := prepare(logger.Error | logger.Warn | logger.Info)
	assert(t, testLogger.Debug, "Test debug message", logger.DebugPrefix, writer)
}

func TestError(t *testing.T) {
	t.Run("Default Error", defaultError)
	t.Run("Only Error", onlyError)
	t.Run("Except Error", exceptError)
}

func TestWarn(t *testing.T) {
	t.Run("Default Warn", defaultWarn)
	t.Run("Only Warn", onlyWarn)
	t.Run("Except Warn", exceptWarn)
}

func TestInfo(t *testing.T) {
	t.Run("Default Info", defaultInfo)
	t.Run("Only Info", onlyInfo)
	t.Run("Except Info", exceptInfo)
}

func TestDebug(t *testing.T) {
	t.Run("Default Debug", defaultDebug)
	t.Run("Only Debug", onlyDebug)
	t.Run("Except Debug", exceptDebug)
}
