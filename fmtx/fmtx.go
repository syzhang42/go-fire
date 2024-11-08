package fmtx

import (
	"fmt"
	"strings"
)

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

func FormatMap[T1 comparable, T2 any](in map[T1]T2) string {
	var temp []string
	for k, v := range in {
		temp = append(temp, fmt.Sprintf("%+v:%+v", k, v))
	}
	return strings.Join(temp, "\n")
}
func FormatSlice[T any](in []T) string {
	var temp []string
	for k, v := range in {
		temp = append(temp, fmt.Sprintf("%+v:%+v", k, v))
	}
	return strings.Join(temp, "\n")
}
