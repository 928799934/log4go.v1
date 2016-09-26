package log4go

type level int

const (
	DEBUG level = iota
	TRACE
	INFO
	WARNING
	ERROR
)

var (
	levelStrings = [...]string{"DEBG", "TRAC", "INFO", "WARN", "EROR"}
)
