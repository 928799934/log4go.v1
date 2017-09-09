package log4go

import (
	"fmt"
	"log"
)

var (
	log4go = NewLog4go()
)

func LoadConfiguration(filename string) error {
	return log4go.LoadConfiguration(filename)
}

func Close() {
	log4go.Close()
}

func Logger(l level) *log.Logger {
	return log4go.Logger(l)
}

func Error(format string, v ...interface{}) {
	logger := log4go.Logger(ERROR)
	if logger == nil {
		std.Output(2, fmt.Sprintf(format, v...))
		return
	}
	logger.Output(2, fmt.Sprintf(format, v...))
}

func Warn(format string, v ...interface{}) {
	logger := log4go.Logger(WARN)
	if logger == nil {
		std.Output(2, fmt.Sprintf(format, v...))
		return
	}
	logger.Output(2, fmt.Sprintf(format, v...))
}

func Info(format string, v ...interface{}) {
	logger := log4go.Logger(INFO)
	if logger == nil {
		std.Output(2, fmt.Sprintf(format, v...))
		return
	}
	logger.Output(2, fmt.Sprintf(format, v...))
}

func Trace(format string, v ...interface{}) {
	logger := log4go.Logger(TRACE)
	if logger == nil {
		std.Output(2, fmt.Sprintf(format, v...))
		return
	}
	logger.Output(2, fmt.Sprintf(format, v...))
}

func Debug(format string, v ...interface{}) {
	logger := log4go.Logger(DEBUG)
	if logger == nil {
		std.Output(2, fmt.Sprintf(format, v...))
		return
	}
	logger.Output(2, fmt.Sprintf(format, v...))
}


func Mark(format string, v ...interface{}) {
	logger := log4go.Logger(MARK)
	if logger == nil {
		std.Output(2, fmt.Sprintf(format, v...))
		return
	}
	logger.Output(2, fmt.Sprintf(format, v...))
}