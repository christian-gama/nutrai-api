package slice

import (
	"errors"
	"reflect"
)

// Contains returns true if the slice contains the element.
func Contains[T comparable](s []T, v T) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

// ContainsFunc returns true if the slice contains an element that satisfies the predicate.
func ContainsFunc[T any](s []T, f func(elem T) bool) bool {
	for _, x := range s {
		if f(x) {
			return true
		}
	}
	return false
}

// Filter returns a Slice chain containing only elements that satisfy the predicate.
func Filter[T any](s []T, f func(elem T) bool) *chain[T] {
	r := make([]T, 0, len(s))
	for _, x := range s {
		if f(x) {
			r = append(r, x)
		}
	}

	return &chain[T]{slice: r}
}

// Map returns a Slice chain containing the results of applying the function to each element.
func Map[T any, R any](s []T, f func(elem T) R) *chain[R] {
	r := make([]R, 0, len(s))
	for _, x := range s {
		r = append(r, f(x))
	}

	return &chain[R]{slice: r}
}

// Pop removes and returns the last element of the slice and also returns a Slice chain.
func Pop[T any](s []T) (T, *chain[T]) {
	if len(s) == 0 || s == nil {
		panic(errors.New("slice is empty"))
	}

	return s[len(s)-1], &chain[T]{slice: s[:len(s)-1]}
}

// Push adds an element to the end of the slice and also returns a Slice chain.
func Push[T any](s []T, v T) *chain[T] {
	return &chain[T]{slice: append(s, v)}
}

// Shift removes and returns the first element of the slice and also returns a Slice chain.
func Shift[T any](s []T) (T, *chain[T]) {
	if len(s) == 0 || s == nil {
		panic(errors.New("slice is empty"))
	}

	return s[0], &chain[T]{slice: s[1:]}
}

// Unshift adds an element to the beginning of the slice and also returns a Slice chain.
func Unshift[T any](s []T, v T) *chain[T] {
	return &chain[T]{slice: append([]T{v}, s...)}
}

// Find returns the first element that satisfies the predicate.
func Find[T any](s []T, f func(elem T) bool) (T, bool) {
	for _, x := range s {
		if f(x) {
			return x, true
		}
	}

	return reflect.Zero(reflect.TypeOf(s[0])).Interface().(T), false
}

// FindIndex returns the index of the first element that satisfies the predicate.
func FindIndex[T any](s []T, f func(elem T) bool) int {
	for i, x := range s {
		if f(x) {
			return i
		}
	}

	return -1
}

// Some returns true if at least one element satisfies the predicate.
func Some[T any](s []T, f func(elem T) bool) bool {
	for _, x := range s {
		if f(x) {
			return true
		}
	}

	return false
}

// Every returns true if all elements satisfy the predicate.
func Every[T any](s []T, f func(elem T) bool) bool {
	for _, x := range s {
		if !f(x) {
			return false
		}
	}

	return true
}

// IndexOf returns the index of the first occurrence of the element in the slice.
func IndexOf[T comparable](s []T, v T) int {
	for i, x := range s {
		if x == v {
			return i
		}
	}

	return -1
}

// Clone returns a shallow copy of the slice.
func Clone[S ~[]E, E any](s S) *chain[E] {
	if s == nil {
		return nil
	}

	return &chain[E]{slice: append(S([]E{}), s...)}
}

// Count returns the number of a given element in the slice.
func Count[T comparable](s []T, v T) int {
	count := 0
	for _, x := range s {
		if x == v {
			count++
		}
	}

	return count
}

// Chain converts a slice to a Slice chain.
func Chain[T any](s []T) *chain[T] {
	return &chain[T]{slice: s}
}

// RemoveDuplicates returns a slice with all duplicate elements removed.
func RemoveDuplicates[T comparable](s []T) []T {
	seen := make(map[T]struct{})
	r := make([]T, 0, len(s))

	for _, x := range s {
		if _, ok := seen[x]; !ok {
			r = append(r, x)
			seen[x] = struct{}{}
		}
	}

	return r
}

// FirstElementOrDefault returns the first element of the slice if it is not empty, otherwise it
// returns the default value.
func FirstElementOrDefault[V any](value []V, defaultValue ...V) V {
	if len(value) > 0 {
		return value[0]
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	zero := reflect.Zero(reflect.TypeOf(value).Elem()).Interface()
	return zero.(V)
}
