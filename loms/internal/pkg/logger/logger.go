package logger

import (
	"bufio"
	"context"
	"log"
	"os"
)

type Logger struct {
	out *bufio.Writer
}

var logger = Logger{
	out: bufio.NewWriter(os.Stdout),
}

func (l *Logger) Info(ctx context.Context, m string) {
	log.Printf("[info]: %s", m)
}

func (l *Logger) Infof(ctx context.Context, m string, args ...any) {
	log.Printf("[info]:"+m, args...)
}

func (l *Logger) Warn(ctx context.Context, m string) {
	log.Printf("[warn]: %s", m)
}

func (l *Logger) Warnf(ctx context.Context, m string, args ...any) {
	log.Printf("[warn]:"+m, args...)
}

func (l *Logger) Error(ctx context.Context, m string) {
	log.Printf("[error] %s", m)
}

func (l *Logger) Errorf(ctx context.Context, m string, args ...any) {
	log.Printf("[error] "+m, args...)
}

func (l *Logger) Panic(ctx context.Context, m string) {
	log.Panic("[panic] %s", m)
}

func (l *Logger) Panicf(ctx context.Context, m string, args ...any) {
	log.Panicf("[fatal] "+m, args...)
}

func (l *Logger) Fatal(ctx context.Context, m string) {
	log.Fatalf("[fatal] %s", m)
}

func (l *Logger) Fatalf(ctx context.Context, m string, args ...any) {
	log.Fatalf("[fatal] "+m, args...)
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
