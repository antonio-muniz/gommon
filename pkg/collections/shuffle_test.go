package collections_test

import (
	"testing"

	"github.com/antonio-muniz/gommon/pkg/collections"
	"github.com/stretchr/testify/require"
)

type shuffleScenario[T any] struct {
	description   string
	elements      []T
	expectedEqual bool
}

func TestShuffle(t *testing.T) {
	scenarios := []shuffleScenario[int]{
		{
			description:   "shuffles the elements in random",
			elements:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expectedEqual: false,
		},
		{
			description:   "works with a single element slice",
			elements:      []int{42},
			expectedEqual: true,
		},
		{
			description:   "works with an empty slice",
			elements:      []int{},
			expectedEqual: true,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			shuffledElements := collections.Shuffle(scenario.elements)
			if scenario.expectedEqual {
				require.Equal(t, scenario.elements, shuffledElements)
			} else {
				require.ElementsMatch(t, scenario.elements, shuffledElements)
				require.NotEqual(t, scenario.elements, shuffledElements)
			}
		})
	}
}
