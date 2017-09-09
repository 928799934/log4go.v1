package log4go

type level int

const (
	DEBUG level = iota
	TRACE
	INFO
	WARN
	ERROR
	MARK
)

var (
	levelStrings = [...]string{"DEBG", "TRAC", "INFO", "WARN", "EROR","MARK"}
)
