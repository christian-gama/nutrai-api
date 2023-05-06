package slice_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/pkg/slice"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type SliceStructSuite struct {
	suite.Suite
}

func TestSliceStructSuite(t *testing.T) {
	suite.RunUnitTest(t, new(SliceStructSuite))
}

func (s *SliceStructSuite) TestChain() {
	ints := func() []int { return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} }

	s.Run("creates a chain from a slice", func() {
		chain := slice.
			Chain(ints()).
			Map(func(elem int) int { return elem * 2 }).
			Filter(func(elem int) bool { return elem%3 == 0 }).
			Unshift(0).
			Push(12)

		s.Equal([]int{0, 6, 12, 18, 12}, chain.Build())
	})

	s.Run("slice.Len returns the length of the slice", func() {
		chain := slice.Chain(ints())

		s.Equal(10, chain.Len())
	})

	s.Run("slice.IsEmpty returns true if the slice is empty", func() {
		chain := slice.Chain([]int{})

		s.True(chain.IsEmpty())
	})

	s.Run("slice.IsEmpty returns false if the slice is not empty", func() {
		chain := slice.Chain(ints())

		s.False(chain.IsEmpty())
	})

	s.Run("slice.Pop returns the last element of the slice", func() {
		chain := slice.Chain(ints())

		pop, chain := chain.Pop()

		s.Equal(10, pop)
		s.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, chain.Build())
	})

	s.Run("slice.Pop panics if the slice is empty", func() {
		chain := slice.Chain([]int{})

		s.Panics(func() { chain.Pop() })
	})

	s.Run("slice.Shift returns the first element of the slice", func() {
		chain := slice.Chain(ints())

		shift, chain := chain.Shift()

		s.Equal(1, shift)
		s.Equal([]int{2, 3, 4, 5, 6, 7, 8, 9, 10}, chain.Build())
	})

	s.Run("slice.Shift panics if the slice is empty", func() {
		chain := slice.Chain([]int{})

		s.Panics(func() { chain.Shift() })
	})

	s.Run("slice.Find returns the first element that matches the predicate", func() {
		chain := slice.Chain(ints())

		find, ok := chain.Find(func(elem int) bool { return elem == 10 })

		s.True(ok)
		s.Equal(10, find)
	})

	s.Run("slice.Find returns false if no element matches the predicate", func() {
		chain := slice.Chain(ints())

		find, ok := chain.Find(func(elem int) bool { return elem == 404 })

		s.False(ok)
		s.Equal(0, find)
	})

	s.Run(
		"slice.FindIndex returns the index of the first element that matches the predicate",
		func() {
			chain := slice.Chain(ints())

			findIndex := chain.FindIndex(func(elem int) bool { return elem == 10 })

			s.Equal(9, findIndex)
		},
	)

	s.Run("slice.FindIndex returns -1 if no element matches the predicate", func() {
		chain := slice.Chain(ints())

		findIndex := chain.FindIndex(func(elem int) bool { return elem == 404 })

		s.Equal(-1, findIndex)
	})

	s.Run("slice.Some returns true if at least one element matches the predicate", func() {
		chain := slice.Chain(ints())

		some := chain.Some(func(elem int) bool { return elem == 10 })

		s.True(some)
	})

	s.Run("slice.Some returns false if no element matches the predicate", func() {
		chain := slice.Chain(ints())

		some := chain.Some(func(elem int) bool { return elem == 404 })

		s.False(some)
	})

	s.Run("slice.Every returns true if all elements match the predicate", func() {
		chain := slice.Chain(ints())

		every := chain.Every(func(elem int) bool { return elem <= 10 })

		s.True(every)
	})

	s.Run("slice.Every returns false if at least one element does not match the predicate", func() {
		chain := slice.Chain(ints())

		every := chain.Every(func(elem int) bool { return elem%2 == 1 })

		s.False(every)
	})

	s.Run("slice.Contains returns true if the slice contains the element", func() {
		chain := slice.Chain(ints())

		contains := slice.Contains(chain.Build(), 10)

		s.True(contains)
	})

	s.Run("slice.Contains returns false if the slice does not contain the element", func() {
		chain := slice.Chain(ints())

		contains := slice.Contains(chain.Build(), 404)

		s.False(contains)
	})

	s.Run("slice.ContainsFunc returns true if the slice contains a matching element", func() {
		chain := slice.Chain(ints())

		containsFunc := slice.ContainsFunc(
			chain.Build(),
			func(elem int) bool { return elem == 10 },
		)

		s.True(containsFunc)
	},
	)

	s.Run(
		"slice.ContainsFunc returns false if the slice does not contain a matching element",
		func() {
			chain := slice.Chain(ints())

			containsFunc := slice.ContainsFunc(
				chain.Build(),
				func(elem int) bool { return elem == 404 },
			)

			s.False(containsFunc)
		},
	)

	s.Run("slice.Count returns the number of elements that match the predicate", func() {
		chain := slice.Chain(ints())

		count := slice.Count(chain.Build(), 10)

		s.Equal(1, count)
	})

	s.Run("slice.Count returns 0 if no element matches the predicate", func() {
		chain := slice.Chain(ints())

		count := slice.Count(chain.Build(), 404)

		s.Equal(0, count)
	})

	s.Run("slice.Clone returns a copy of the slice", func() {
		chain := slice.Chain(ints())

		clone := slice.Clone(chain.Build())

		s.Equal(chain.Build(), clone.Build())
	})

	s.Run("slice.IndexOf returns the index of the first element that matches the element", func() {
		chain := slice.Chain(ints())

		indexOf := slice.IndexOf(chain.Build(), 10)

		s.Equal(9, indexOf)
	})

	s.Run("slice.IndexOf returns -1 if no element matches the element", func() {
		chain := slice.Chain(ints())

		indexOf := slice.IndexOf(chain.Build(), 404)

		s.Equal(-1, indexOf)
	})

	s.Run("slice.RemoveDuplicates returns a slice without duplicate elements", func() {
		chain := slice.Chain(ints())

		withoutDuplicates := slice.RemoveDuplicates(chain.Build())

		s.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, withoutDuplicates)
	})
}
