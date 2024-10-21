package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	GoRoutineErrorLevel
)

type Logger struct {
	MinLevel LogLevel
	Writer   io.Writer
	Flag     int
}

func NewLogger(minLevel LogLevel, writer io.Writer, flag int) *Logger {
	log.SetOutput(writer)
	log.SetFlags(flag)

	return &Logger{MinLevel: minLevel, Writer: writer, Flag: flag}
}

func (l *Logger) log(level LogLevel, message string) {
	if l.Writer == nil {
		fmt.Fprintln(os.Stdout, "Logger STDOUT is nil")
		return
	}

	levelName := "DEBUG"

	switch level {
	case DebugLevel:
		levelName = "DEBUG"
	case InfoLevel:
		levelName = "INFO"
	case WarnLevel:
		levelName = "WARN"
	case ErrorLevel:
		levelName = "ERROR"
	case FatalLevel:
		levelName = "FATAL"
	case GoRoutineErrorLevel:
		levelName = "GOROUTINE_ERROR"
	}

	if level >= l.MinLevel {
		oRaw := fmt.Sprintf("[%s] %s\n", levelName, message)
		log.Print(oRaw)
	}
}

func (l *Logger) Debug(format string) {
	l.log(DebugLevel, format)
}

func (l *Logger) Debugf(format string, a ...any) {
	l.log(DebugLevel, fmt.Sprintf(format, a...))
}

func (l *Logger) Info(format string) {
	l.log(InfoLevel, format)
}

func (l *Logger) Infof(format string, a ...any) {
	l.log(InfoLevel, fmt.Sprintf(format, a...))
}

func (l *Logger) Warn(format string) {
	l.log(WarnLevel, format)
}

func (l *Logger) Warnf(format string, a ...any) {
	l.log(WarnLevel, fmt.Sprintf(format, a...))
}

func (l *Logger) Error(format string) {
	l.log(ErrorLevel, format)
}

func (l *Logger) Errorf(format string, a ...any) {
	l.log(ErrorLevel, fmt.Sprintf(format, a...))
}

func (l *Logger) Fatal(format string) {
	l.log(FatalLevel, format)
	// os.Exit(1)
	l.wait()
}

func (l *Logger) Fatalf(format string, a ...any) {
	l.log(FatalLevel, fmt.Sprintf(format, a...))
	// os.Exit(1)
	l.wait()
}

func (l *Logger) wait() {
	l.log(FatalLevel, "Press CTRL+C to Exit.")

	_, _ = fmt.Scanln()
}

var SharedLogger *Logger //nolint:gochecknoglobals // wontfix
