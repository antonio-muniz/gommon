package collections

func Generate[T any](count uint, generator func() T) []T {
	result := make([]T, count, count)
	for index := uint(0); index < count; index++ {
		result[index] = generator()
	}
	return result
}
