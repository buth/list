// Package list implements a structurally-immutable singly linked list.
package list

type List[T any] struct {
	Value T
	next  *List[T]
}

func (l *List[T]) Push(x T) *List[T] {
	return &List[T]{
		Value: x,
		next:  l,
	}
}

func (l *List[T]) Pop() *List[T] {
	if l == nil {
		return nil
	}

	return l.next
}

func (l *List[T]) Slice(dst []T) []T {
	if l == nil {
		return nil
	}

	dst = size(dst, l.len())
	for i := 0; l != nil; i++ {
		dst[i] = l.Value
		l = l.next
	}

	return dst
}

func (l *List[T]) ReverseSlice(dst []T) []T {
	if l == nil {
		return nil
	}

	dst = size(dst, l.len())
	for i := len(dst) - 1; i >= 0; i-- {
		dst[i] = l.Value
		l = l.next
	}

	return dst
}

func (l *List[T]) len() int {
	i := 0
	for ; l != nil; i++ {
		l = l.next
	}

	return i
}

func size[T any](s []T, n int) []T {
	switch {
	case len(s) == n:
	case cap(s) >= n:
		s = s[:n]
	default:
		s = make([]T, n)
	}

	return s
}
