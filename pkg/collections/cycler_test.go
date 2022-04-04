package collections_test

import (
	"testing"

	"github.com/antonio-muniz/gommon/pkg/collections"
	"github.com/stretchr/testify/require"
)

func TestNext(t *testing.T) {
	cycler := collections.NewCycler([]int{1, 2, 3, 4})

	require.Equal(t, 1, cycler.Next())
	require.Equal(t, 2, cycler.Next())
	require.Equal(t, 3, cycler.Next())
	require.Equal(t, 4, cycler.Next())
	require.Equal(t, 1, cycler.Next())
	require.Equal(t, 2, cycler.Next())
	require.Equal(t, 3, cycler.Next())
	require.Equal(t, 4, cycler.Next())
}

type dropCurrentScenario[T any] struct {
	description string
	elements    []T
	flow        func(*testing.T, *collections.Cycler[T])
}

func TestDropCurrent(t *testing.T) {
	t.Run("drops middle element", func(t *testing.T) {
		cycler := collections.NewCycler([]string{"A", "B", "C"})

		require.Equal(t, "A", cycler.Next())
		require.Equal(t, "B", cycler.Next())

		cycler.DropCurrent()

		require.Equal(t, "C", cycler.Next())
		require.Equal(t, "A", cycler.Next())
		require.Equal(t, "C", cycler.Next())
		require.Equal(t, "A", cycler.Next())
		require.Equal(t, "C", cycler.Next())
	})

	t.Run("drops first element", func(t *testing.T) {
		cycler := collections.NewCycler([]string{"A", "B", "C"})

		require.Equal(t, "A", cycler.Next())

		cycler.DropCurrent()

		require.Equal(t, "B", cycler.Next())
		require.Equal(t, "C", cycler.Next())
		require.Equal(t, "B", cycler.Next())
		require.Equal(t, "C", cycler.Next())
	})

	t.Run("drops last element", func(t *testing.T) {
		cycler := collections.NewCycler([]string{"A", "B", "C"})

		require.Equal(t, "A", cycler.Next())
		require.Equal(t, "B", cycler.Next())
		require.Equal(t, "C", cycler.Next())

		cycler.DropCurrent()

		require.Equal(t, "A", cycler.Next())
		require.Equal(t, "B", cycler.Next())
		require.Equal(t, "A", cycler.Next())
		require.Equal(t, "B", cycler.Next())
	})
}

func TestSize(t *testing.T) {
	cycler := collections.NewCycler([]int{1, 2, 3, 4})
	require.Equal(t, 4, cycler.Size())
}
