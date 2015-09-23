package log

import (
	"fmt"
	"log/syslog"
)

type SyslogLogger struct {
	writer *syslog.Writer
}

func (p *SyslogLogger) Debug(format string, args ...interface{}) {
	p.writer.Debug(fmt.Sprintf(format, args...))
}
func (p *SyslogLogger) Error(format string, args ...interface{}) {
	p.writer.Err(fmt.Sprintf(format, args...))
}
func (p *SyslogLogger) Info(format string, args ...interface{}) {
	p.writer.Info(fmt.Sprintf(format, args...))
}
func (p *SyslogLogger) Warning(format string, args ...interface{}) {
	p.writer.Warning(fmt.Sprintf(format, args...))
}
func (p *SyslogLogger) Notice(format string, args ...interface{}) {
	p.writer.Notice(fmt.Sprintf(format, args...))
}

func syslogPriority(priority Priority) syslog.Priority {
	switch priority {
	case INFO:
		return syslog.LOG_INFO
	case WARNING:
		return syslog.LOG_WARNING
	case ERROR:
		return syslog.LOG_ERR
	}
	return syslog.LOG_DEBUG
}

func NewSyslogLogger(tag string, priority Priority) (Logger, error) {
	if wrt, err := syslog.New(syslogPriority(priority), tag); err == nil {
		return &SyslogLogger{writer: wrt}, nil
	} else {
		return nil, err
	}
}

func DialSyslogLogger(network, address, tag string, priority Priority) (Logger, error) {
	if wrt, err := syslog.Dial(network, address, syslogPriority(priority), tag); err == nil {
		return &SyslogLogger{writer: wrt}, nil
	} else {
		return nil, err
	}
}
