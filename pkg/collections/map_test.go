package collections_test

import (
	"fmt"
	"testing"

	"github.com/antonio-muniz/gommon/pkg/collections"
	"github.com/stretchr/testify/require"
)

type mapScenario[TIn any, TOut any] struct {
	description            string
	elements               []TIn
	mapper                 func(TIn) TOut
	expectedMappedElements []TOut
}

func TestMap(t *testing.T) {
	scenarios := []mapScenario[int, string]{
		{
			description:            "transforms all elements using the mapper function",
			elements:               []int{1, 2, 3, 4, 5},
			mapper:                 func(number int) string { return fmt.Sprintf("Mapped #%d", number) },
			expectedMappedElements: []string{"Mapped #1", "Mapped #2", "Mapped #3", "Mapped #4", "Mapped #5"},
		},
		{
			description:            "returns an empty slice for an empty slice",
			elements:               []int{},
			mapper:                 func(number int) string { return fmt.Sprintf("Mapped #%d", number) },
			expectedMappedElements: []string{},
		},
		{
			description:            "returns an empty slice for nil",
			elements:               nil,
			mapper:                 func(number int) string { return fmt.Sprintf("Mapped #%d", number) },
			expectedMappedElements: []string{},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			mappedElements := collections.Map(scenario.elements, scenario.mapper)
			require.Equal(t, scenario.expectedMappedElements, mappedElements)
		})
	}
}
