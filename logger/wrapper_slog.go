// 搬运自github.com/cocktail828/go-tools
package logger

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/syzhang42/go-fire/pkg/slog"
	"github.com/syzhang42/go-fire/stringsx"
)

// slogWrapper implements Logger interface
type slogWrapper struct {
	l *slog.Logger
}

// NewLoggerWithSlog creates a new logger which wraps
// the given logrus.Logger
func NewLoggerWithSlog(logger *slog.Logger) Logger {
	logger.Error("slog begin......................................................")
	return slogWrapper{
		l: logger.With(
			slog.String("caller", getFileAndLine()),
		),
	}
}

func (c slogWrapper) Debugw(msg string, args ...any)    { c.l.Debug(msg, args...) }
func (c slogWrapper) Debugf(format string, args ...any) { c.l.Debug(fmt.Sprintf(format, args...)) }

func (c slogWrapper) Infow(msg string, args ...any)    { c.l.Info(msg, args...) }
func (c slogWrapper) Infof(format string, args ...any) { c.l.Info(fmt.Sprintf(format, args...)) }

func (c slogWrapper) Warnw(msg string, args ...any)    { c.l.Warn(msg, args...) }
func (c slogWrapper) Warnf(format string, args ...any) { c.l.Warn(fmt.Sprintf(format, args...)) }

func (c slogWrapper) Errorw(msg string, args ...any)    { c.l.Error(msg, args...) }
func (c slogWrapper) Errorf(format string, args ...any) { c.l.Error(fmt.Sprintf(format, args...)) }

func (c slogWrapper) With(args ...any) Logger {
	return NewLoggerWithSlog(c.l.With(args...))
}

func (c slogWrapper) WithGroup(name string) Logger {
	return NewLoggerWithSlog(c.l.WithGroup(name))
}

func getFileAndLine() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "unknown"
	}
	files := stringsx.Split(file, "/")
	if len(file) < 2 {
		return file + ":" + strconv.Itoa(line)
	} else {
		return fmt.Sprintf("%v/%v:%v", files[len(files)-2], files[len(files)-1], strconv.Itoa(line))
	}
}
