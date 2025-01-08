package mathx

import (
	"testing"

	"github.com/syzhang42/go-fire/fmtx"
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

func TestAnySort(t *testing.T) {
	fmtx.PrintSlice(AnySort([]string{"10", "2", "9.9"}, func(v1, v2 string) bool { return v1 < v2 }))
	fmtx.PrintSlice(AnySort([]string{"10", "2", "9.9"}, func(v1, v2 string) bool { return v1 > v2 }))
	fmtx.PrintSlice(AnySort([]float32{10, 2, 9.9}, func(v1, v2 float32) bool { return v1 > v2 }))
	fmtx.PrintSliceJson(AnySort([]A{
		{
			Aa: "1",
			Bb: 1,
			Cc: []string{"1"},
			Dd: &B{
				Bb: 1,
			},
		},
		{
			Aa: "2",
			Bb: 2,
			Cc: []string{"2"},
			Dd: &B{
				Bb: 2,
			},
		},
		{
			Aa: "3",
			Bb: 3,
			Cc: []string{"3"},
			Dd: &B{
				Bb: 3,
			},
		},
	}, func(v1, v2 A) bool { return v1.Dd.Bb < v2.Dd.Bb }))
}
