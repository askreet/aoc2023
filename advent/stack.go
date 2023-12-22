package advent

type Stack[T any] struct {
	data []T
	def  T
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Pop() T {
	r := s.data[0]
	s.data = s.data[1:]
	return r
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}
