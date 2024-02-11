# Go-logger


## Go Logger Package with JSON and Text Handlers

This package provides a simple and customizable logger for Go applications, offering both text and JSON output options.

### Features

- Supports both text and JSON output formats.
- Configurable logger level for filtering messages.
- Ability to add source information to logs (file name and line number).
- Enrichment of logs with program information (PID, Go version).
- Flexible output destination (defaulting to standard output).

### Usage
```go
package main

import (
	"log/slog"

	"github.com/otyang/go-logger"
)

func main() {
	// Create logger with JSON output
	log := logger.NewSlog(
		true,           // Use JSON handler
		true,           // Add source information
		slog.LevelInfo, // Log level (e.g., Info, Debug, Error)
		nil,            // Use default output (os.Stdout)
	)

	// Add program information
	logger := logger.WithBuildInfo(log)

	// Log messages with different levels
	logger.Debug("This is a debug message")
	logger.Info("This is an informational message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")

	// Log with additional fields
	logger.Info("Custom message", "key", "value")
}
```
