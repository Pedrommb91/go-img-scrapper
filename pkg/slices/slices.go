package slices

func Contains[T any](ss []T, exists func(T) bool) (int, bool) {
	for i, s := range ss {
		if exists(s) {
			return i, true
		}
	}
	return -1, false
}
