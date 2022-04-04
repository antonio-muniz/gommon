package collections_test

import (
	"testing"

	"github.com/antonio-muniz/gommon/pkg/collections"
	"github.com/stretchr/testify/require"
)

type removeScenario[T any] struct {
	description    string
	elements       []T
	index          int
	expectedResult []T
}

func TestRemove(t *testing.T) {
	scenarios := []removeScenario[string]{
		{
			description:    "returns a new slice without the element at the specified index",
			elements:       []string{"A", "B", "C", "D"},
			index:          1,
			expectedResult: []string{"A", "C", "D"},
		},
		{
			description:    "returns an empty slice when given a single element slice and index 0",
			elements:       []string{"A"},
			index:          0,
			expectedResult: []string{},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			result := collections.Remove(scenario.elements, scenario.index)
			require.Equal(t, scenario.expectedResult, result)
		})
	}
}
