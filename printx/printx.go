package printx

//已弃用
import "fmt"

// 已弃用
func PrintMap[T1 comparable, T2 any](in map[T1]T2) {
	for key, value := range in {
		fmt.Printf("Key: %+v, Value: %+v\n", key, value)
	}
}

// 已弃用
func PrintSlice[T any](in []T) {
	for key, value := range in {
		fmt.Printf("%v:%+v\n", key, value)
	}
}
