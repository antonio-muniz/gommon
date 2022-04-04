package collections_test

import (
	"testing"

	"github.com/antonio-muniz/gommon/pkg/collections"
	"github.com/stretchr/testify/require"
)

type filterScenario[T any] struct {
	description              string
	elements                 []T
	predicate                func(T) bool
	expectedFilteredElements []T
}

func TestFilter(t *testing.T) {
	scenarios := []filterScenario[int]{
		{
			description:              "filters elements according to the predicate",
			elements:                 []int{1, 2, 3, 4, 5, 6},
			predicate:                func(number int) bool { return number%2 == 0 },
			expectedFilteredElements: []int{2, 4, 6},
		},
		{
			description:              "returns an empty slice if no elements match the predicate",
			elements:                 []int{1, 2, 3, 4, 5, 6},
			predicate:                func(number int) bool { return false },
			expectedFilteredElements: []int{},
		},
		{
			description:              "returns an identical slice if all elements match the predicate",
			elements:                 []int{1, 2, 3, 4, 5, 6},
			predicate:                func(number int) bool { return true },
			expectedFilteredElements: []int{1, 2, 3, 4, 5, 6},
		},
		{
			description:              "returns an empty slice if given an empty slice",
			elements:                 []int{},
			predicate:                func(number int) bool { return number%2 == 0 },
			expectedFilteredElements: []int{},
		},
		{
			description:              "returns an empty slice if given a nil slice",
			elements:                 nil,
			predicate:                func(number int) bool { return number%2 == 0 },
			expectedFilteredElements: []int{},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			filteredElements := collections.Filter(scenario.elements, scenario.predicate)
			require.Equal(t, scenario.expectedFilteredElements, filteredElements)
		})
	}
}
