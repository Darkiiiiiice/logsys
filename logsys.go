package logsys

import (
	"fmt"
	"io"
	"log"
)

const (
	TRACE = 0x01
	DEBUG = 0x02
	INFO  = 0x04
	WARN  = 0x08
	ERROR = 0x10
	FATAL = 0x20

	ALL  = 0x3F
	NONE = 0x00

	mask = 0xFF
)

const (
	trace = "TRACE"
	debug = "DEBUG"
	info  = "INFO "
	warn  = "WARN "
	error = "ERROR"
	fatal = "FATAL"
)

const (
	black     = 30
	red       = 31
	green     = 32
	yellow    = 33
	blue      = 34
	purple    = 35
	lightblue = 36
	white     = 37
)

const (
	LStdOut  = 0x01
	LFileOut = 0x02
)

type Logger struct {
	logger *log.Logger
	level  int
	flags  int
}

func Init(out io.Writer, prefix string, flags int, level int) *Logger {
	log.Fatalln()

	if LFileOut == flags {
		prefix = fmt.Sprintf(" %s ", prefix)
	} else if LStdOut == flags {
		prefix = fmt.Sprintf("\x1b[0;%dm %s \x1b[0m ", 33, prefix)
	}

	logger := log.New(out, prefix, log.LstdFlags|log.Lmicroseconds)
	return &Logger{
		logger: logger,
		flags:  flags,
		level:  mask - level,
	}
}

func (l *Logger) Trace(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.output(s, white, TRACE, trace)
}

func (l *Logger) Debug(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.output(s, yellow, DEBUG, debug)
}

func (l *Logger) Info(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.output(s, green, INFO, info)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.output(s, lightblue, WARN, warn)
}

func (l *Logger) Error(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.output(s, red, ERROR, error)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.output(s, purple, FATAL, fatal)
}

func (l *Logger) output(s string, color int, level int, flag string) {
	if l.level != l.level|level {
		if LFileOut == l.flags {
			l.logger.Printf("[ %s ] %s", flag, s)
		} else if LStdOut == l.flags {
			l.logger.Printf("\x1b[0;%dm[\x1b[0m\x1b[%d;%dm %s \x1b[0m\x1b[0;%dm]\x1b[0m \x1b[0;%dm %s \x1b[0m", color, color+10, black, flag, color, color, s)
		}
	}
}
