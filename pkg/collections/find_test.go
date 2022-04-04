package collections_test

import (
	"testing"

	"github.com/antonio-muniz/gommon/pkg/collections"
	"github.com/stretchr/testify/assert"
)

type findFirstIndexScenario[T any] struct {
	description   string
	elements      []T
	predicate     func(T) bool
	expectedIndex int
}

func TestFindFirstIndex(t *testing.T) {
	scenarios := []findFirstIndexScenario[string]{
		{
			description:   "returns the index of the matching element",
			elements:      []string{"A", "B", "C", "D"},
			predicate:     func(element string) bool { return element == "C" },
			expectedIndex: 2,
		},
		{
			description:   "returns the index of the first matching element",
			elements:      []string{"A", "B", "C", "D"},
			predicate:     func(element string) bool { return true },
			expectedIndex: 0,
		},
		{
			description:   "returns -1 if there are no matching elements",
			elements:      []string{"A", "B", "C", "D"},
			predicate:     func(element string) bool { return false },
			expectedIndex: -1,
		},
		{
			description:   "returns -1 if no elements are provided",
			elements:      []string{},
			predicate:     func(element string) bool { return true },
			expectedIndex: -1,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			index := collections.FindFirstIndex(scenario.elements, scenario.predicate)
			assert.Equal(t, scenario.expectedIndex, index)
		})
	}
}
