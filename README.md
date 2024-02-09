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
	"fmt"
	"github.com/your-username/go-logger"
)

func main() {
	// Create logger with JSON output 
    logger := logger.NewSlog(
        true,    // Use JSON handler
        true,    // Add source information
        slog.InfoLevel, // Log level (e.g., Info, Debug, Error)
        nil,     // Use default output (os.Stdout)
    )

	// Add program information
	logger = logger.WithBuildInfo()

	// Log messages with different levels
	logger.Debug("This is a debug message")
	logger.Info("This is an informational message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")

	// Log with additional fields
	logger.Info("Custom message", "key", "value")

	fmt.Println("Logging completed")
}
```

 

### Additional Notes

- This package leverages the `log/slog` library for underlying functionality.
- Consider adding documentation for each function and exported variable in the package.

I hope this helps!
