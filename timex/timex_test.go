package timex

import (
	"fmt"
	"testing"
	"time"
)

func TestAnyToTimeDuration(t *testing.T) {
	fmt.Println(AnyToTimeDuration(int(1), time.Second))
	fmt.Println(AnyToTimeDuration(int8(1), time.Second))
	fmt.Println(AnyToTimeDuration(int16(1), time.Second))
	fmt.Println(AnyToTimeDuration(int32(1), time.Second))
	fmt.Println(AnyToTimeDuration(int64(1), time.Second))

	fmt.Println(AnyToTimeDuration(uint(1), time.Second))
	fmt.Println(AnyToTimeDuration(uint8(1), time.Second))
	fmt.Println(AnyToTimeDuration(uint16(1), time.Second))
	fmt.Println(AnyToTimeDuration(uint32(1), time.Second))
	fmt.Println(AnyToTimeDuration(uint64(1), time.Second))

	fmt.Println(AnyToTimeDuration(float32(1.1), time.Second))
	fmt.Println(AnyToTimeDuration(float64(1.2), time.Millisecond))

	fmt.Println(AnyToTimeDuration("1", time.Minute))
	fmt.Println(AnyToTimeDuration("1.93", time.Second))
	fmt.Println(AnyToTimeDuration("1.93s", time.Second))
	fmt.Println(AnyToTimeDuration("2m", time.Minute))
}
