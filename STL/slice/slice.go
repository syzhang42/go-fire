package slice

func Slice2Map[T comparable](in []T) map[T]struct{} {
	out := make(map[T]struct{}, len(in))
	for _, v := range in {
		out[v] = struct{}{}
	}
	return out
}

func MapVal2Slice[T1 comparable, T2 any](in map[T1]T2) []T2 {
	out := make([]T2, len(in))
	for _, v := range in {
		out = append(out, v)
	}
	return out
}

func MapKey2Slice[T1 comparable, T2 any](in map[T1]T2) []T1 {
	out := make([]T1, len(in))
	for k := range in {
		out = append(out, k)
	}
	return out
}
