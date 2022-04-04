package collections

func Remove[T any](elements []T, index int) []T {
	return Concat(elements[:index], elements[index+1:])
}
