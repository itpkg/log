package log_test

import (
	"testing"
	"time"

	"github.com/itpkg/log"
)

func TestSyslog(t *testing.T) {
	if logger, err := log.NewSyslogLogger("test", log.DEBUG); err == nil {
		logger.Debug("Now: %v", time.Now())
	} else {
		t.Errorf("error on syslog: %v", err)
	}
}

func TestStdout(t *testing.T) {
	logger := log.NewStdoutLogger("test", log.INFO)
	logger.Debug("Should not display: %v", time.Now())
	logger.Info("Hello: %v", time.Now())
	logger.Info("Hello: %v", time.Now())
}
