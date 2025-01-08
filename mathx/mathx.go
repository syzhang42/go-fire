package mathx

import "sort"

type SortBy[T any] struct {
	t    []T
	less func(i, j T) bool
}

func (a SortBy[T]) Len() int           { return len(a.t) }
func (a SortBy[T]) Swap(i, j int)      { a.t[i], a.t[j] = a.t[j], a.t[i] }
func (a SortBy[T]) Less(i, j int) bool { return a.less(a.t[i], a.t[j]) }

func AnySort[T any](data []T, less func(i, j T) bool) []T {
	localSby := SortBy[T]{
		t:    data,
		less: less,
	}
	sort.Sort(localSby)
	return localSby.t
}
