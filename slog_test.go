package logger

import (
	"log/slog"
	"os"
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSlog(t *testing.T) {
	tests := []struct {
		name               string
		useJsonHandlerType bool
		addSource          bool
		level              slog.Leveler
		wantOutput         string
	}{
		{
			name:               "text handler, no source",
			useJsonHandlerType: false,
			addSource:          false,
			level:              slog.LevelError,
		},
		{
			name:               "json handler, added source",
			useJsonHandlerType: true,
			addSource:          true,
			level:              slog.LevelError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				output = os.Stdout
				opts   = slog.HandlerOptions{
					AddSource: tt.addSource,
					Level:     tt.level,
				}
			)

			var want *slog.Logger
			{
				want = slog.New(slog.NewTextHandler(output, &opts))
				if tt.useJsonHandlerType {
					want = slog.New(slog.NewJSONHandler(output, &opts))
				}
			}

			got := NewSlog(tt.useJsonHandlerType, tt.addSource, tt.level, os.Stdout)

			assert.Equal(t, want, got)
		})
	}
}

func TestWithBuildInfo(t *testing.T) {
	// want
	opts := slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelError,
	}

	buildInfo, _ := debug.ReadBuildInfo()

	want := slog.
		New(slog.NewJSONHandler(os.Stdout, &opts)).
		With(
			slog.Group("program_info",
				slog.Int("pid", os.Getpid()),
				slog.String("go_version", buildInfo.GoVersion),
			),
		)

	// got
	got := NewSlog(true, false, slog.LevelError, os.Stdout)
	got = WithBuildInfo(got)

	// assert
	assert.Equal(t, want, got)
}
