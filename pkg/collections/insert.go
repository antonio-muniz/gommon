package collections

func InsertAt[T any](elements []T, element T, index uint) []T {
	result := make([]T, 0)
	result = append(result, elements[:index]...)
	result = append(result, element)
	result = append(result, elements[index:]...)
	return result
}
