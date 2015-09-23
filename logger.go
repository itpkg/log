package log

type Logger interface {
	Debug(format string, args ...interface{})
	Error(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Notice(format string, args ...interface{})
}

type Priority int

const (
	WARNING Priority = iota
	NOTICE
	INFO
	DEBUG
	ERROR
)
