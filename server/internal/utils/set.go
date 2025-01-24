package utils

type Empty struct{}

type Set[T comparable] map[T]Empty

func (s Set[T]) Add(item T) {
	s[item] = Empty{}
}

func (s Set[T]) Contain(item T) bool {
	_, exists := s[item]
	return exists
}

func (s Set[T]) Range(fn func(T)) {
	for t := range s {
		fn(t)
	}
}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}
