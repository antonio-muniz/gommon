package collections

func Clone[T any](slice []T) []T {
	clone := make([]T, len(slice))
	copy(clone, slice)
	return clone
}
