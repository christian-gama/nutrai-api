package slice

type chain[V any] struct {
	slice []V
}

// Build returns the slice.
func (s *chain[V]) Build() []V {
	return s.slice
}

func (s *chain[V]) Len() int {
	return len(s.slice)
}

func (s *chain[V]) IsEmpty() bool {
	return len(s.slice) == 0
}

func (s *chain[V]) Push(v V) *chain[V] {
	return Push(s.slice, v)
}

func (s *chain[V]) Pop() (V, *chain[V]) {
	return Pop(s.slice)
}

func (s *chain[V]) Filter(f func(elem V) bool) *chain[V] {
	return Filter(s.slice, f)
}

func (s *chain[V]) Map(f func(elem V) V) *chain[V] {
	return Map(s.slice, f)
}

func (s *chain[V]) Shift() (V, *chain[V]) {
	return Shift(s.slice)
}

func (s *chain[V]) Unshift(v V) *chain[V] {
	return Unshift(s.slice, v)
}

func (s *chain[V]) Find(f func(elem V) bool) (V, bool) {
	return Find(s.slice, f)
}

func (s *chain[V]) FindIndex(f func(elem V) bool) int {
	return FindIndex(s.slice, f)
}

func (s *chain[V]) Some(f func(elem V) bool) bool {
	return Some(s.slice, f)
}

func (s *chain[V]) Every(f func(elem V) bool) bool {
	return Every(s.slice, f)
}

func (s *chain[V]) Clone() *chain[V] {
	return Clone(s.slice)
}
