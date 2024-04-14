package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"route256.ozon.ru/project/notifier/internal/config"
)

type Logger struct {
	logger *zap.Logger
}

var logger = NewLogger(WithDebugLevel(), WithOutputStdout())

func NewLogger(opts ...ConfigOption) *Logger {
	logger, err := NewConfig().Build()
	if err != nil {
		panic(err)
	}
	logger = logger.With(zap.String("app", config.AppName))
	return &Logger{
		logger: logger,
	}
}

func Sync() {
	if logger != nil && logger.logger != nil {
		logger.logger.Sync()
	}
}

func (l *Logger) Info(ctx context.Context, m string) {
	l.logger.Info(m, getFields(ctx)...)
}

func (l *Logger) Infof(ctx context.Context, m string, args ...any) {
	l.logger.Info(fmt.Sprintf(m, args...), getFields(ctx)...)
}

func (l *Logger) Warn(ctx context.Context, m string) {
	l.logger.Warn(m, getFields(ctx)...)
}

func (l *Logger) Warnf(ctx context.Context, m string, args ...any) {
	l.logger.Warn(fmt.Sprintf(m, args...), getFields(ctx)...)
}

func (l *Logger) Error(ctx context.Context, m string) {
	l.logger.Error(m, getFields(ctx)...)
}

func (l *Logger) Errorf(ctx context.Context, m string, args ...any) {
	l.logger.Warn(fmt.Sprintf(m, args...), getFields(ctx)...)
}

func (l *Logger) Panic(ctx context.Context, m string) {
	l.logger.Panic(m, getFields(ctx)...)
}

func (l *Logger) Panicf(ctx context.Context, m string, args ...any) {
	l.logger.Panic(fmt.Sprintf(m, args...), getFields(ctx)...)
}

func (l *Logger) Fatal(ctx context.Context, m string) {
	l.logger.Fatal(m, getFields(ctx)...)
}

func (l *Logger) Fatalf(ctx context.Context, m string, args ...any) {
	l.logger.Fatal(fmt.Sprintf(m, args...), getFields(ctx)...)
}

func Info(ctx context.Context, m string) {
	logger.Info(ctx, m)
}

func Infof(ctx context.Context, m string, args ...any) {
	logger.Infof(ctx, m, args...)
}

func Warn(ctx context.Context, m string) {
	logger.Info(ctx, m)
}

func Warnf(ctx context.Context, m string, args ...any) {
	logger.Warnf(ctx, m, args...)
}

func Error(ctx context.Context, m string) {
	logger.Error(ctx, m)
}

func Errorf(ctx context.Context, m string, args ...any) {
	logger.Errorf(ctx, m, args...)
}

func Panic(ctx context.Context, m string) {
	logger.Panic(ctx, m)
}

func Panicf(ctx context.Context, m string, args ...any) {
	logger.Panicf(ctx, m, args...)
}

func Fatal(ctx context.Context, m string) {
	logger.Fatal(ctx, m)
}

func Fatalf(ctx context.Context, m string, args ...any) {
	logger.Fatalf(ctx, m, args...)
}

func getFields(ctx context.Context) []zap.Field {
	fields := make([]zap.Field, 0, 2)
	value := ctx.Value("trace_id")
	if value != nil {
		fields = append(fields, zap.String("trace_id", value.(string)))
	}
	value = ctx.Value("span_id")
	if value != nil {
		fields = append(fields, zap.String("span_id", value.(string)))
	}
	return fields
}
