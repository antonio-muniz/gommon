package collections

type Cycler[T any] struct {
	elements  []T
	nextIndex int
}

func NewCycler[T any](elements []T) *Cycler[T] {
	return &Cycler[T]{elements: elements}
}

func (cycler *Cycler[T]) Next() T {
	element := cycler.elements[cycler.nextIndex]
	cycler.nextIndex = (cycler.nextIndex + cycler.Size() + 1) % cycler.Size()
	return element
}

func (cycler *Cycler[T]) DropCurrent() {
	dropIndex := (cycler.nextIndex + cycler.Size() - 1) % cycler.Size()
	cycler.elements = Remove(cycler.elements, dropIndex)
	cycler.nextIndex = dropIndex % cycler.Size()
}

func (cycler *Cycler[T]) Size() int {
	return len(cycler.elements)
}
