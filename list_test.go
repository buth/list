package list

import (
	"testing"
)

func TestList(t *testing.T) {
	var l *List[int]
	for i := 0; i < 10; i++ {
		l = l.Push(i)
	}

	s := l.Slice(nil)
	if n := len(s); n != 10 {
		t.Errorf("expected %d, got %d", 10, n)
	}

	for i := 0; i < 10; i++ {
		e := 9 - i
		if s[i] != e {
			t.Errorf("expected %d, got %d", e, s[i])
		}
	}

	for i := 0; i < 10; i++ {
		e := 9 - i
		if l.Value != e {
			t.Errorf("expected %d, got %d", e, l.Value)
		}

		l = l.Pop()
	}

	if l != nil {
		t.Error("expected nil")
	}
}

func BenchmarkList(b *testing.B) {
	var l *List[int]
	for i := 0; i < 1024; i++ {
		l = l.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Slice(nil)
	}
}

func BenchmarkListWithSlice(b *testing.B) {
	var l *List[int]
	for i := 0; i < 1024; i++ {
		l = l.Push(i)
	}

	s := make([]int, 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Slice(s)
	}
}
