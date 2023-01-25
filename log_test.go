package logger_test

import (
	"testing"

	logger "git.sonicoriginal.software/logger"
)

func TestDefaultLoggerError(t *testing.T) {
	logger.DefaultLogger.Error("Test error message")
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
