package collections

import (
	"math/rand"
	"time"
)

func Shuffle[T any](elements []T) []T {
	shuffledElements := Clone(elements)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	random.Shuffle(len(elements), func(i, j int) {
		shuffledElements[i], shuffledElements[j] = shuffledElements[j], shuffledElements[i]
	})
	return shuffledElements
}
