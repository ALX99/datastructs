package lists_test

import (
	"testing"

	"github.com/alx99/datastructs/lists"
)

func TestPush(t *testing.T) {
	t.Parallel()
	tests := []struct {
		values      []int
		expectedLen uint
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{}, 0},
	}

	for _, test := range tests {
		l := lists.NewLinkedList[int]()
		l.Push(test.values...)
		if l.Len() != test.expectedLen {
			t.Fatalf("expected %d but got %d", test.expectedLen, l.Len())
		}
	}
}

func TestPop(t *testing.T) {
	t.Parallel()
	tests := []struct {
		values           []int
		expectedPopOrder []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		l := lists.NewLinkedList[int]()
		l.Push(test.values...)
		for _, expected := range test.expectedPopOrder {
			actual, ok := l.Pop()
			if !ok {
				t.Fatalf("expected true but got %t", ok)
			}
			if actual != expected {
				t.Fatalf("expected %d but got %d", expected, actual)
			}
		}
	}
}

func TestPeek(t *testing.T) {
	t.Parallel()
	tests := []struct {
		values     []int
		expected   int
		expectedOk bool
	}{
		{[]int{99}, 99, true},
		{[]int{99, -1}, -1, true},
		{[]int{}, 0, false},
	}

	for _, test := range tests {
		l := lists.NewLinkedList[int]()
		l.Push(test.values...)
		actual, ok := l.Peek()
		if ok != test.expectedOk {
			t.Fatalf("expected %t but got %t", test.expectedOk, ok)
		}
		if actual != test.expected {
			t.Fatalf("expected %d but got %d", test.expected, actual)
		}
	}
}

func TestLocate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		values   []int
		item     int
		expected int
	}{
		{[]int{1, 2, 3}, 2, 1},
		{[]int{1, 2, 3, 4, 5}, 5, 0},
		{[]int{5, 4, 3, 2, 1}, 5, 4},
		{[]int{1, 2, 3, 4, 5}, 99, -1},
		{[]int{}, 99, -1},
	}

	for _, test := range tests {
		l := lists.NewComparableLinkedList[int]()
		l.Push(test.values...)
		actual := l.Locate(test.item)
		if actual != test.expected {
			t.Fatalf("expected %d but got %d", test.expected, actual)
		}
	}
}

func TestHas(t *testing.T) {
	t.Parallel()
	tests := []struct {
		values   []int
		item     int
		expected bool
	}{
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3, 4, 5}, 5, true},
		{[]int{5, 4, 3, 2, 1}, 5, true},
		{[]int{1, 2, 3, 4, 5}, 99, false},
		{[]int{}, 99, false},
	}

	for _, test := range tests {
		l := lists.NewComparableLinkedList[int]()
		l.Push(test.values...)
		actual := l.Has(test.item)
		if actual != test.expected {
			t.Fatalf("expected %t but got %t", test.expected, actual)
		}
	}
}

func TestRemove(t *testing.T) {
	t.Parallel()
	tests := []struct {
		values   []int
		item     int
		expected bool
	}{
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3, 4, 5}, 5, true},
		{[]int{5, 4, 3, 2, 1}, 5, true},
		{[]int{1, 2, 3, 4, 5}, 99, false},
		{[]int{}, 99, false},
	}

	for _, test := range tests {
		l := lists.NewComparableLinkedList[int]()
		l.Push(test.values...)
		actual := l.Remove(test.item)
		if actual != test.expected {
			t.Fatalf("expected %t but got %t", test.expected, actual)
		}
	}
}

func BenchmarkPush(b *testing.B) {
	l := lists.NewLinkedList[int]()
	for i := 0; i < b.N; i++ {
		l.Push(i)
	}
}
