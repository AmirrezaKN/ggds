package hashset

import "fmt"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](vals ...T) Set[T] {
	s := Set[T]{}
	for _, v := range vals {
		s[v] = struct{}{}
	}
	return s
}

func (s *Set[T]) Add(vals ...T) {
	for _, v := range vals {
		(*s)[v] = struct{}{}
	}
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Members() []T {
	result := make([]T, 0, len(s))
	for v := range s {
		result = append(result, v)
	}
	return result
}

func (s Set[T]) String() string {
	return fmt.Sprintf("%v", s.Members())
}

func (s Set[T]) Union(s2 Set[T]) Set[T] {
	result := NewSet(s.Members()...)
	result.Add(s2.Members()...)
	return result
}

func (s Set[T]) Intersection(s2 Set[T]) Set[T] {
	result := NewSet[T]()
	for _, v := range s.Members() {
		if s2.Contains(v) {
			result.Add(v)
		}
	}
	return result
}
