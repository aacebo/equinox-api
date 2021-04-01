package log

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	_info  *log.Logger
	_warn  *log.Logger
	_error *log.Logger
}

func New(namespace string) *Logger {
	var l = new(Logger)

	l._info = log.New(os.Stderr, color.CyanString(fmt.Sprintf("[INFO][%s] ", namespace)), log.Ldate|log.Ltime)
	l._warn = log.New(os.Stderr, color.YellowString(fmt.Sprintf("[WARN][%s] ", namespace)), log.Ldate|log.Ltime)
	l._error = log.New(os.Stderr, color.RedString(fmt.Sprintf("[ERROR][%s] ", namespace)), log.Ldate|log.Ltime)

	return l
}

func (l *Logger) Info(args ...interface{}) {
	l._info.Print(args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l._info.Printf(format, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l._warn.Print(args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l._warn.Printf(format, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l._error.Print(args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l._error.Printf(format, args...)
}
