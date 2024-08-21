package printx

import "fmt"

func PrintMap[T1 comparable, T2 any](in map[T1]T2) {
	for key, value := range in {
		fmt.Printf("Key: %+v, Value: %+v\n", key, value)
	}
}
func PrintSlice[T any](in []T) {
	for key, value := range in {
		fmt.Printf("%v:%+v\n", key, value)
	}
}
