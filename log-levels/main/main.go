package main

import (
	"fmt"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

// See https://github.com/go-logr/zapr
func main() {
	// Has DEBUG as default log-levels level
	config := zap.NewDevelopmentConfig()

	// Set log-levels level as INFO
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	zapLogger, err := config.Build()

	if err != nil {
		panic(fmt.Sprintf("who watches the watchmen (%v)?", err))
	}

	logger := zapr.NewLogger(zapLogger)

	logger.Info("This is a info message")

	// Log a message at the debug level. Should not be logged.
	debugLogger := logger.V(1)
	debugLogger.Info("This is a debug message")

	// Log a message at the trace level (verbosity level 2). Should not be logged.
	traceLogger := logger.V(2)
	traceLogger.Info("This is a trace message")

	// You can also use logger.V() with a context to specify the verbosity level.
	// This will logger messages at the specified level. Should not be logged.
	logger.V(3).Info("This is a custom verbosity level message")
}
