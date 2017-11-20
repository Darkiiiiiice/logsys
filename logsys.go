package logsys

import (
	"fmt"
	"io"
	"runtime"
	"sync"
	"time"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

const (
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
	color bool
	level int

	out io.Writer
	buf []byte

	lock sync.Mutex
}

var logger *Logger

func Init(out io.Writer, level int, color bool) {
	if logger == nil {
		logger = &Logger{
			color: color,
			level: level,
			out:   out,
			buf:   make([]byte, 0),
		}
	}
}

func itoa(buf *[]byte, i int) {

	var b [20]byte
	var bp = len(b) - 1

	for i >= 10 {
		q := i / 10
		b[bp] = byte('0' + (i - q*10))
		bp--
		i = q
	}
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

func (l *Logger) formatTime(buf *[]byte, now *time.Time) {

	//
	year, month, day := now.Date()
	itoa(buf, year)
	*buf = append(*buf, '-')
	itoa(buf, int(month))
	*buf = append(*buf, '-')
	itoa(buf, day)
	*buf = append(*buf, ' ')
	//
	hour, min, sec := now.Clock()
	itoa(buf, hour)
	*buf = append(*buf, ':')
	itoa(buf, min)
	*buf = append(*buf, ':')
	itoa(buf, sec)
	*buf = append(*buf, '.')
	//
	msec := now.Nanosecond() / 1e3
	itoa(buf, msec)
	*buf = append(*buf, ' ')
}

func (l *Logger) formatPath(buf *[]byte, file string, line int) {
	if l.level > DEBUG {
		return
	}

	*buf = append(*buf, file...)
	*buf = append(*buf, ':')
	itoa(buf, line)
	*buf = append(*buf, ' ')

}

func (l *Logger) formatHeader(buf *[]byte) {
	var now = time.Now()

	l.formatTime(&l.buf, &now)

	if _, file, line, ok := runtime.Caller(4); ok {
		l.formatPath(&l.buf, file, line)
	}

}

func (l *Logger) Output(str string) {

	l.lock.Lock()
	defer l.lock.Unlock()
	l.buf = l.buf[:0]
	l.formatHeader(&l.buf)

	l.buf = append(l.buf, str...)
	if len(str) == 0 || str[len(str)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}

	l.out.Write(l.buf)
}

func (l *Logger) Debug(str string) {
	var s string
	if l.color {
		s = fmt.Sprintf("\x1b[0;%dm[\x1b[0m\x1b[%d;%dm %s \x1b[0m\x1b[0;%dm]\x1b[0m \x1b[0;%dm %s \x1b[0m", white, white+10, black, debug, white, white, str)

	} else {
		s = fmt.Sprintf("[ %s ] %s", debug, str)
	}
	l.Output(s)
}

func (l *Logger) Info(str string) {
	var s string
	if l.color {
		s = fmt.Sprintf("\x1b[0;%dm[\x1b[0m\x1b[%d;%dm %s \x1b[0m\x1b[0;%dm]\x1b[0m \x1b[0;%dm %s \x1b[0m", green, green+10, black, info, green, green, str)

	} else {
		s = fmt.Sprintf("[ %s ] %s", info, str)
	}
	l.Output(s)
}

func (l *Logger) Warn(str string) {
	var s string
	if l.color {
		s = fmt.Sprintf("\x1b[0;%dm[\x1b[0m\x1b[%d;%dm %s \x1b[0m\x1b[0;%dm]\x1b[0m \x1b[0;%dm %s \x1b[0m", yellow, yellow+10, black, warn, yellow, yellow, str)

	} else {
		s = fmt.Sprintf("[ %s ] %s", warn, str)
	}
	l.Output(s)
}

func (l *Logger) Error(str string) {
	var s string
	if l.color {
		s = fmt.Sprintf("\x1b[0;%dm[\x1b[0m\x1b[%d;%dm %s \x1b[0m\x1b[0;%dm]\x1b[0m \x1b[0;%dm %s \x1b[0m", red, red+10, black, error, red, red, str)

	} else {
		s = fmt.Sprintf("[ %s ] %s", error, str)
	}
	l.Output(s)
}

func Debug(format string, v ...interface{}) {
	if logger.level <= DEBUG {
		logger.Debug(fmt.Sprintf(format, v...))
	}
}
func Warn(format string, v ...interface{}) {
	if logger.level <= WARN {
		logger.Warn(fmt.Sprintf(format, v...))
	}
}
func Info(format string, v ...interface{}) {
	if logger.level <= INFO {
		logger.Info(fmt.Sprintf(format, v...))
	}
}
func Error(format string, v ...interface{}) {
	if logger.level <= ERROR {
		logger.Error(fmt.Sprintf(format, v...))
	}
}

//type Logger struct {
//	logger *log.Logger
//	level  int
//	flags  int
//}
//
//var logger *Logger

//func Init(out io.Writer, prefix string, flags int, level int) *Logger {
//
//	//runtime.Caller(1)
//
//	if LFileOut == flags {
//		prefix = fmt.Sprintf(" %s ", prefix)
//	} else if LStdOut == flags {
//		prefix = fmt.Sprintf("\x1b[0;%dm %s \x1b[0m ", 33, prefix)
//	}
//
//	logger := log.New(out, prefix, log.LstdFlags|log.Lmicroseconds)
//	return &Logger{
//		logger: logger,
//		flags:  flags,
//		level:  mask - level,
//	}
//}
//
//func (l *Logger) Trace(format string, v ...interface{}) {
//	s := fmt.Sprintf(format, v...)
//	l.output(s, white, TRACE, trace)
//}
//
//func (l *Logger) Debug(format string, v ...interface{}) {
//	s := fmt.Sprintf(format, v...)
//	l.output(s, yellow, DEBUG, debug)
//}
//
//func (l *Logger) Info(format string, v ...interface{}) {
//	s := fmt.Sprintf(format, v...)
//	l.output(s, green, INFO, info)
//}
//
//func (l *Logger) Warn(format string, v ...interface{}) {
//	s := fmt.Sprintf(format, v...)
//	l.output(s, lightblue, WARN, warn)
//}
//
//func (l *Logger) Error(format string, v ...interface{}) {
//	s := fmt.Sprintf(format, v...)
//	l.output(s, red, ERROR, error)
//}
//
//func (l *Logger) Fatal(format string, v ...interface{}) {
//	s := fmt.Sprintf(format, v...)
//	l.output(s, purple, FATAL, fatal)
//}
//
//func (l *Logger) output(s string, color int, level int, flag string) {
//	if l.level != l.level|level {
//		if LFileOut == l.flags {
//			l.logger.Printf("[ %s ] %s", flag, s)
//		} else if LStdOut == l.flags {
//			l.logger.Printf("\x1b[0;%dm[\x1b[0m\x1b[%d;%dm %s \x1b[0m\x1b[0;%dm]\x1b[0m \x1b[0;%dm %s \x1b[0m", color, color+10, black, flag, color, color, s)
//		}
//	}
//}
