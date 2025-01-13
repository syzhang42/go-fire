package auth

import "testing"

type B struct {
	Bb int
}
type A struct {
	Aa string
	Bb int
	Cc []string
	Dd *B
}

func TestStdlog(t *testing.T) {
	StdDebugSlice("123", []*A{
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

	StdDebugMap("123", map[string]*A{
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
	StdDebugJson("123", []A{
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

	StdDebugJson("123", map[string]A{
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
