package collections_test

import (
	"testing"

	"github.com/antonio-muniz/gommon/pkg/collections"
	"github.com/stretchr/testify/require"
)

type generateScenario[T any] struct {
	description    string
	count          uint
	generator      func() T
	expectedResult []T
}

func TestGenerate(t *testing.T) {
	scenarios := []generateScenario[int]{
		{
			description:    "generates the specified number of elements using the generator",
			count:          5,
			generator:      func() int { return 42 },
			expectedResult: []int{42, 42, 42, 42, 42},
		},
		{
			description:    "returns an empty slice if count is zero",
			count:          0,
			generator:      func() int { return 42 },
			expectedResult: []int{},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			result := collections.Generate(scenario.count, scenario.generator)
			require.Equal(t, scenario.expectedResult, result)
		})
	}
}
