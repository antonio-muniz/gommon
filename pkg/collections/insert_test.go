package collections_test

import (
	"testing"

	"github.com/antonio-muniz/gommon/pkg/collections"
	"github.com/stretchr/testify/assert"
)

type insertAtScenario[T any] struct {
	description    string
	elements       []T
	element        T
	index          uint
	expectedResult []T
}

func TestInsertAt(t *testing.T) {
	scenarios := []insertAtScenario[string]{
		{
			description:    "returns a new slice with the element added at the specified index",
			elements:       []string{"A", "B", "C"},
			element:        "X",
			index:          2,
			expectedResult: []string{"A", "B", "X", "C"},
		},
		{
			description:    "inserts the new element properly at the first position",
			elements:       []string{"A", "B", "C"},
			element:        "X",
			index:          0,
			expectedResult: []string{"X", "A", "B", "C"},
		},
		{
			description:    "inserts the new element properly at the last position",
			elements:       []string{"A", "B", "C"},
			element:        "X",
			index:          3,
			expectedResult: []string{"A", "B", "C", "X"},
		},
		{
			description:    "inserts the new element properly into an empty slice",
			elements:       []string{},
			element:        "X",
			index:          0,
			expectedResult: []string{"X"},
		},
		{
			description:    "inserts the new element properly into a nil slice",
			elements:       nil,
			element:        "X",
			index:          0,
			expectedResult: []string{"X"},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			result := collections.InsertAt(scenario.elements, scenario.element, scenario.index)
			assert.Equal(t, scenario.expectedResult, result)
		})
	}
}
