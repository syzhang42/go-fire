package print

import "testing"

func TestXxx(t *testing.T) {
	PrintMap(map[string]string{"2k": "2v", "3k": "3v"})
	PrintMap(map[int]string{2: "2v", 3: "3v"})
	PrintSlice([]string{"2v", "3v"})
	PrintSlice([]int64{2, 3})
}
