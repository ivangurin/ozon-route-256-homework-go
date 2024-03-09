package logger

import (
	"bufio"
	"log"
	"os"
)

type tLogger struct {
	out *bufio.Writer
}

var logger = tLogger{
	out: bufio.NewWriter(os.Stdout),
}

func (l *tLogger) info(m string) {
	log.Printf("[info]: %s", m)
}

func (l *tLogger) warn(m string) {
	log.Printf("[warn]: %s", m)
}

func (l *tLogger) error(m string, err error) {
	log.Printf("[error] %s: %s", m, err)
}

func (l *tLogger) fatal(m string, err error) {
	log.Fatalf("[fatal] %s: %s", m, err)
}

func Info(m string) {
	logger.info(m)
}

func Warn(m string) {
	logger.warn(m)
}

func Error(m string, err error) {
	logger.error(m, err)
}

func Fatal(m string, err error) {
	logger.fatal(m, err)
}
