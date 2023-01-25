package logger_test

import (
	"testing"

	logger "git.sonicoriginal.software/logger"
)

func TestDefaultLoggerError(t *testing.T) {
	logger.DefaultLogger.Error("Test error message")
}

func TestDefaultLoggerWarn(t *testing.T) {
	logger.DefaultLogger.Warn("Test error message")
}

func TestDefaultLoggerInfo(t *testing.T) {
	logger.DefaultLogger.Info("Test error message")
}

func TestDefaultLoggerDebug(t *testing.T) {
	logger.DefaultLogger.Debug("Test error message")
}
