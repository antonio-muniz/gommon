package collections

func Filter[T any](elements []T, filterFn func(T) bool) []T {
	filteredElements := make([]T, 0, len(elements))
	for _, element := range elements {
		if filterFn(element) {
			filteredElements = append(filteredElements, element)
		}
	}
	return filteredElements
}
