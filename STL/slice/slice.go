package slice

func Slice2Map[T comparable](in []T) map[T]struct{} {
	out := make(map[T]struct{}, len(in))
	for _, v := range in {
		out[v] = struct{}{}
	}
	return out
}

func Map2Slice[T comparable](in map[T]struct{}) []T {
	out := make([]T, len(in))
	for k := range in {
		out = append(out, k)
	}
	return out
}
