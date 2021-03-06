package log

import (
	"log"
)

type StdoutLogger struct {
	tag       string
	priority  Priority
	priorityS string
}

func (p *StdoutLogger) write(priority Priority, format string, args ...interface{}) {
	if p.priority >= priority {
		log.Printf(p.priorityS+" "+format, args...)
	}
}

func (p *StdoutLogger) Debug(format string, args ...interface{}) {
	p.write(DEBUG, format, args...)
}
func (p *StdoutLogger) Error(format string, args ...interface{}) {
	p.write(ERROR, format, args...)
}
func (p *StdoutLogger) Info(format string, args ...interface{}) {
	p.write(INFO, format, args...)
}
func (p *StdoutLogger) Warning(format string, args ...interface{}) {
	p.write(WARNING, format, args...)
}
func (p *StdoutLogger) Notice(format string, args ...interface{}) {
	p.write(NOTICE, format, args...)
}

func NewStdoutLogger(tag string, priority Priority) Logger {
	var priorityS string
	switch priority {
	case DEBUG:
		priorityS = "DEBUG"
	case INFO:
		priorityS = "INFO"
	case WARNING:
		priorityS = "WARNING"
	case NOTICE:
		priorityS = "NOTICE"
	case ERROR:
		priorityS = "ERROR"
	}
	return &StdoutLogger{tag: tag, priority: priority, priorityS: priorityS}
}
