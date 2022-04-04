package collections_test

import (
	"testing"

	"github.com/antonio-muniz/gommon/pkg/collections"
	"github.com/stretchr/testify/assert"
)

type cloneScenario[T any] struct {
	description   string
	slice         []T
	expectedClone []T
}

func TestClone(t *testing.T) {
	scenarios := []cloneScenario[int]{
		{
			description:   "returns a new slice equal to the provided slice",
			slice:         []int{1, 2, 3, 4},
			expectedClone: []int{1, 2, 3, 4},
		},
		{
			description:   "returns a new empty slice if given an empty slice",
			slice:         []int{},
			expectedClone: []int{},
		},
		{
			description:   "returns a new empty slice if given a nil slice",
			slice:         nil,
			expectedClone: []int{},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			clone := collections.Clone(scenario.slice)
			assert.Equal(t, scenario.expectedClone, clone)
		})
	}
}
