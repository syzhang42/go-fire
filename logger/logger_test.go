package logger

import (
	"testing"
)

func TestXxx(t *testing.T) {
	dl := NewLoggerWithLumberjack(
		Config{
			Level:    "debug",
			Filename: "./go-fire.log",
			MaxSize:  100,
			MaxAge:   1,
			MaxCount: 1,
			Compress: false,
		},
	)
	dl.Debugw("test", "key:", 1)
	dl.Debugf("test,key:%v", 1)
}
