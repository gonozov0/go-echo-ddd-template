package logger_test

import (
	"context"
	"log/slog"
	"testing"

	"go-echo-template/pkg/logger"
)

func TestSetup(_ *testing.T) {
	logger.Setup()
	slog.InfoContext(context.Background(), "test logging")
}
