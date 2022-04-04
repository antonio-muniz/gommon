package collections

func Map[TIn any, TOut any](elements []TIn, mapper func(TIn) TOut) []TOut {
	mappedElements := make([]TOut, len(elements))
	for index, element := range elements {
		mappedElements[index] = mapper(element)
	}
	return mappedElements
}
