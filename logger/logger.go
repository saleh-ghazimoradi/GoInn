package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	logger    *slog.Logger
	level     slog.Leveler
	addSource bool
	output    *os.File
}

type Option func(*Logger)

func WithLevel(level slog.Level) Option {
	return func(l *Logger) {
		l.level = level
	}
}

func WithSource(enable bool) Option {
	return func(l *Logger) {
		l.addSource = enable
	}
}

func WithOutput(output *os.File) Option {
	return func(l *Logger) {
		l.output = output
	}
}

func NewLogger(opts ...Option) *Logger {
	logger := &Logger{
		level:     slog.LevelInfo,
		addSource: false,
		output:    os.Stderr,
	}

	for _, opt := range opts {
		opt(logger)
	}

	if logger.output == nil {
		logger.output = os.Stderr
	}

	handlerOpts := &slog.HandlerOptions{
		Level:     logger.level,
		AddSource: logger.addSource,
	}

	handler := slog.NewJSONHandler(logger.output, handlerOpts)
	logger.logger = slog.New(handler)

	return logger
}

func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *Logger) Close() error {
	if l.output != os.Stdout && l.output != os.Stderr && l.output != nil {
		return l.output.Close()
	}
	return nil
}
