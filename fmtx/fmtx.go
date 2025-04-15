package fmtx

import (
	"encoding/json"
	"fmt"
	"strings"
)

// 适用于简单性map
func PrintMap[T1 comparable, T2 any](in map[T1]T2) {
	fmt.Printf("Total %v elements:\n", len(in))
	for key, value := range in {
		fmt.Printf("\tKey: %+v, Value: %+v\n", key, value)
	}
}

// 所有参数必须可见
func PrintMapJson[T1 comparable, T2 any](in map[T1]T2) {
	fmt.Printf("Total %v elements:\n", len(in))
	for key, value := range in {
		inJSON, err := json.Marshal(value)
		if err != nil {
			fmt.Printf("print error:%v,not use it\n", err)
			return
		}
		fmt.Printf("\tKey: %+v, Value: %+v\n", key, string(inJSON))
	}
}

// 适用于简单性slice
func PrintSlice[T any](in []T) {
	fmt.Printf("Total %v elements:\n", len(in))
	for key, value := range in {
		fmt.Printf("\t%v:%+v\n", key, value)
	}
}

// // 所有参数必须可见
func PrintSliceJson[T any](in []T) {
	fmt.Printf("Total %v elements:\n", len(in))
	for key, value := range in {
		valueJSON, err := json.Marshal(value)
		if err != nil {
			fmt.Printf("print error:%v,not use it\n", err)
			return
		}
		fmt.Printf("\t%v:%+v\n", key, string(valueJSON))
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

// 有风险，使用后test下是否是自己想要的
func FormatOtherJson[T any](in T) string {
	valueJSON, err := json.Marshal(in)
	if err != nil {
		fmt.Printf("error:%v,not use it\n", err)
		return fmt.Sprintf("error:%v,not use it\n", err)
	}
	return string(valueJSON)
}
func FormatOtherJsonIndent[T any](in T) string {

	valueJSON, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		fmt.Printf("error: %v, not use it\n", err)
		return fmt.Sprintf("error: %v, not use it\n", err)
	}
	return string(valueJSON)
}
