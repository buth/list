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

func (l *List[T]) slice(s []T, n int) []T {
	if l.next != nil {
		s = l.next.slice(s, n+1)
	} else {
		s = size(s, n+1)
	}

	s[n] = l.Value
	return s
}

func (l *List[T]) Slice(dst []T) []T {
	if l == nil {
		return nil
	}

	return l.slice(dst, 0)
}

func (l *List[T]) reverseSlice(s []T, n int) []T {
	m := n + 1
	if l.next != nil {
		s = l.next.reverseSlice(s, m)
		s[len(s)-m] = l.Value
	} else {
		s = size(s, m)
		s[0] = l.Value
	}

	return s
}

func (l *List[T]) ReverseSlice(dst []T) []T {
	if l == nil {
		return nil
	}

	return l.reverseSlice(dst, 0)
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
