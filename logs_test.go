package logs_test

import (
	"github.com/moskvorechie/logs"
	"testing"
)

func TestInitLogsToFile(t *testing.T) {
	l, _ := logs.New("main.log")
	l.Info("123")
}
