package log

import (
	"os"
	"testing"
)

func TestStdStreamLog(t *testing.T) {
	h, _ := NewDefaultStreamHandler(os.Stdout)
	s := NewDefault(h)
	s.Info("hello world")

	s.Close()

	Info("hello world")
}
