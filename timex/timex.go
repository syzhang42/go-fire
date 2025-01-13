package timex

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var (
	reInt   = regexp.MustCompile(`^[+-]?\d+$`)
	reFloat = regexp.MustCompile(`^[+-]?\d*\.\d+$`)
	reDur   = regexp.MustCompile(`^\d+(\.\d+)?(ns|µs|us|ms|s|m|h)$`)
)

func AnyToTimeDuration[T any](in T, level time.Duration) (time.Duration, error) {
	switch v := any(in).(type) {
	case int:
		return time.Duration(v) * level, nil
	case int8:
		return time.Duration(v) * level, nil
	case int16:
		return time.Duration(v) * level, nil
	case int32:
		return time.Duration(v) * level, nil
	case int64:
		return time.Duration(v) * level, nil

	case uint:
		return time.Duration(v) * level, nil
	case uint8:
		return time.Duration(v) * level, nil
	case uint16:
		return time.Duration(v) * level, nil
	case uint32:
		return time.Duration(v) * level, nil
	case uint64:
		return time.Duration(v) * level, nil

	case float64:
		return time.Duration(v * float64(level)), nil
	case float32:
		return time.Duration(v * float32(level)), nil

	case string:
		if reInt.MatchString(v) {
			if val, err := strconv.Atoi(v); err == nil {
				return time.Duration(val) * level, nil
			}
		}
		if reFloat.MatchString(v) {
			if val, err := strconv.ParseFloat(v, 32); err == nil {
				return time.Duration(float32(val) * float32(level)), nil
			}
		}
		if reDur.MatchString(v) {
			if dur, err := time.ParseDuration(v); err == nil {
				return dur, nil
			}
		}
	}
	return 0, fmt.Errorf("unsupported type:%v", in)
}
