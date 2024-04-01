package logger

import (
	"bufio"
	"log"
	"os"
)

type Logger struct {
	out *bufio.Writer
}

var logger = Logger{
	out: bufio.NewWriter(os.Stdout),
}

func (l *Logger) Info(m string) {
	log.Printf("[info]: %s", m)
}

func (l *Logger) Infof(m string, args ...any) {
	log.Printf("[info]:"+m, args...)
}

func (l *Logger) Warn(m string) {
	log.Printf("[warn]: %s", m)
}

func (l *Logger) Warnf(m string, args ...any) {
	log.Printf("[warn]:"+m, args...)
}

func (l *Logger) Error(m string) {
	log.Printf("[error] %s", m)
}

func (l *Logger) Errorf(m string, args ...any) {
	log.Printf("[error] "+m, args...)
}

func (l *Logger) Panic(m string) {
	log.Panicf("[panic] %s", m)
}

func (l *Logger) Panicf(m string, args ...any) {
	log.Panicf("[panic] "+m, args...)
}

func (l *Logger) Fatal(m string) {
	log.Fatalf("[fatal] %s", m)
}

func (l *Logger) Fatalf(m string, args ...any) {
	log.Fatalf("[fatal] "+m, args...)
}

func Info(m string) {
	logger.Info(m)
}

func Infof(m string, args ...any) {
	logger.Infof(m, args...)
}

func Warn(m string) {
	logger.Info(m)
}

func Warnf(m string, args ...any) {
	logger.Warnf(m, args...)
}

func Error(m string) {
	logger.Error(m)
}

func Errorf(m string, args ...any) {
	logger.Errorf(m, args...)
}

func Panic(m string) {
	logger.Panic(m)
}

func Panicf(m string, args ...any) {
	logger.Panicf(m, args...)
}

func Fatal(m string) {
	logger.Fatal(m)
}

func Fatalf(m string, args ...any) {
	logger.Fatalf(m, args...)
}
