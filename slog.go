package logger

import (
	"io"
	"log/slog"
	"os"
	"runtime/debug"
)

func NewSlog(useJsonHandlerType bool, addSource bool, level slog.Leveler, output io.Writer) *slog.Logger {
	opts := slog.HandlerOptions{
		AddSource: addSource,
		Level:     level,
	}

	if output == nil {
		output = os.Stdout
	}

	// default handler text
	logger := slog.New(slog.NewTextHandler(output, &opts))

	// if something else overide
	if useJsonHandlerType {
		logger = slog.New(slog.NewJSONHandler(output, &opts))
	}

	slog.SetDefault(logger)
	return logger
}

// WithBuildInfo enriches the logger with program information.
func WithBuildInfo(logger *slog.Logger) *slog.Logger {
	buildInfo, _ := debug.ReadBuildInfo()

	logger = logger.With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
			slog.String("go_version", buildInfo.GoVersion),
		),
	)

	return logger
}
