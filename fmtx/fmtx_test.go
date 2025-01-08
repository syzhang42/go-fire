package fmtx

import (
	"fmt"
	"testing"
)

type B struct {
	Bb int
}
type A struct {
	Aa string
	Bb int
	Cc []string
	Dd *B
}

func TestPrint(t *testing.T) {
	PrintSlice([]*A{
		{
			Aa: "123",
			Bb: 123,
			Cc: []string{"123", "123"},
			Dd: &B{Bb: 123},
		},
		{
			Aa: "456",
			Bb: 456,
			Cc: []string{"456", "456"},
			Dd: &B{Bb: 456},
		},
	})

	PrintMap(map[string]*A{
		"123": {
			Aa: "123",
			Bb: 123,
			Cc: []string{"123", "123"},
			Dd: &B{Bb: 123},
		},
		"456": {
			Aa: "456",
			Bb: 456,
			Cc: []string{"456", "456"},
			Dd: &B{Bb: 456},
		},
	})
	PrintSliceJson([]A{
		{
			Aa: "123",
			Bb: 123,
			Cc: []string{"123", "123"},
			Dd: &B{Bb: 123},
		},
		{
			Aa: "456",
			Bb: 456,
			Cc: []string{"456", "456"},
			Dd: &B{Bb: 456},
		},
	})

	PrintMapJson(map[string]A{
		"123": {
			Aa: "123",
			Bb: 123,
			Cc: []string{"123", "123"},
			Dd: &B{Bb: 123},
		},
		"456": {
			Aa: "456",
			Bb: 456,
			Cc: []string{"456", "456"},
			Dd: &B{Bb: 456},
		},
	})

}
func TestFormat(t *testing.T) {
	fmt.Println(FormatSlice([]*A{
		{
			Aa: "123",
			Bb: 123,
			Cc: []string{"123", "123"},
		},
		{
			Aa: "456",
			Bb: 456,
			Cc: []string{"456", "456"},
		},
	}))

	fmt.Println(FormatMap(map[string]*A{
		"123": {
			Aa: "123",
			Bb: 123,
			Cc: []string{"123", "123"},
		},
		"456": {
			Aa: "456",
			Bb: 456,
			Cc: []string{"456", "456"},
		},
	}))
}
