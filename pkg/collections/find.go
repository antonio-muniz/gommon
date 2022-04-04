package collections

func FindFirstIndex[T any](elements []T, predicate func(T) bool) int {
	for index, element := range elements {
		if predicate(element) {
			return index
		}
	}
	return -1
}
