package collections_test

import (
	"testing"

	"github.com/antonio-muniz/gommon/pkg/collections"
	"github.com/stretchr/testify/require"
)

type concatScenario[T any] struct {
	description    string
	inputs         [][]T
	expectedResult []T
}

func TestConcat(t *testing.T) {
	scenarios := []concatScenario[int]{
		{
			description: "combines all slices into one, retaining their sequence",
			inputs: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expectedResult: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			description: "ignores empty slices",
			inputs: [][]int{
				{1, 2, 3},
				{},
				{7, 8, 9},
			},
			expectedResult: []int{1, 2, 3, 7, 8, 9},
		},
		{
			description: "ignores nil slices",
			inputs: [][]int{
				{1, 2, 3},
				nil,
				{7, 8, 9},
			},
			expectedResult: []int{1, 2, 3, 7, 8, 9},
		},
		{
			description: "returns an empty slice given multiple empty slices",
			inputs: [][]int{
				{},
				nil,
				{},
			},
			expectedResult: []int{},
		},
		{
			description:    "returns an empty slice given zero slices",
			inputs:         [][]int{},
			expectedResult: []int{},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			result := collections.Concat(scenario.inputs...)
			require.Equal(t, scenario.expectedResult, result)
		})
	}
}
