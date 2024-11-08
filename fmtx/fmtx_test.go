package fmtx

import (
	"fmt"
	"testing"
)

type A struct {
	a string
	b int
	c []string
}

func TestXxx(t *testing.T) {
	fmt.Println(FormatSlice([]*A{
		{
			a: "123",
			b: 123,
			c: []string{"123", "123"},
		},
		{
			a: "456",
			b: 456,
			c: []string{"456", "456"},
		},
	}))

	fmt.Println(FormatMap(map[string]*A{
		"123": {
			a: "123",
			b: 123,
			c: []string{"123", "123"},
		},
		"456": {
			a: "456",
			b: 456,
			c: []string{"456", "456"},
		},
	}))
}
