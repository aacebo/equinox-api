package logger

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
	var self = new(Logger)

	self._info = log.New(os.Stderr, color.CyanString(fmt.Sprintf("[INFO][%s] ", namespace)), log.Ldate|log.Ltime)
	self._warn = log.New(os.Stderr, color.YellowString(fmt.Sprintf("[WARN][%s] ", namespace)), log.Ldate|log.Ltime)
	self._error = log.New(os.Stderr, color.RedString(fmt.Sprintf("[ERROR][%s] ", namespace)), log.Ldate|log.Ltime)

	return self
}

func (self *Logger) Info(args ...interface{}) {
	self._info.Print(args...)
}

func (self *Logger) Infof(format string, args ...interface{}) {
	self._info.Printf(format, args...)
}

func (self *Logger) Warn(args ...interface{}) {
	self._warn.Print(args...)
}

func (self *Logger) Warnf(format string, args ...interface{}) {
	self._warn.Printf(format, args...)
}

func (self *Logger) Error(args ...interface{}) {
	self._error.Print(args...)
}

func (self *Logger) Errorf(format string, args ...interface{}) {
	self._error.Printf(format, args...)
}
