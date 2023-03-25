package stack

type Stack[T comparable] struct {
	data []T
}

func New[T comparable](vals ...T) Stack[T] {
	return Stack[T]{
		data: vals,
	}
}

func (s Stack[T]) Len() int {
	return len(s.data)
}

func (s *Stack[T]) Push(vals ...T) {
	s.data = append(s.data, vals...)
}

func (s *Stack[T]) Pop() (v T, ok bool) {
	if len(s.data) == 0 {
		return v, false
	}
	v = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, true
}
