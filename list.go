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
		m := n + 1
		switch {
		case len(s) == m:
		case cap(s) >= m:
			s = s[:m]
		default:
			s = make([]T, m)
		}
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
