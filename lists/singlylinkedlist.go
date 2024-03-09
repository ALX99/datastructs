package lists

type linkedList[T any] interface {
	Push(item ...T)
	Pop() (T, bool)
	Peek() (T, bool)
	Get(index uint) (T, bool)
	Delete(index uint) (T, bool)
	Len() uint
}

type comparableLinkedList[T comparable] interface {
	Locate(item T) int
	Has(item T) bool
	Remove(item T) bool
	linkedList[T]
}

// these are here to make sure the interface does not change
var (
	_ linkedList[int]           = &LinkedList[int]{}
	_ comparableLinkedList[int] = &ComparableLinkedList[int]{}
)

type LinkedList[T any] struct {
	head *node[T]
	len  uint
}

type ComparableLinkedList[T comparable] struct {
	LinkedList[T]
}

// NewLinkedList returns a new singly linked list.
func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T]{}
}

func NewComparableLinkedList[T comparable]() ComparableLinkedList[T] {
	return ComparableLinkedList[T]{}
}

// Push adds new items to the front of the list.
func (l *LinkedList[T]) Push(items ...T) {
	for _, item := range items {
		l.push(item)
	}
}

func (l *LinkedList[T]) push(item T) {
	l.len++
	l.head = &node[T]{value: item, next: l.head}
}

// Pop removes and return the first item from the list.
func (l *LinkedList[T]) Pop() (T, bool) {
	if l.len == 0 || l.head == nil {
		var t T
		return t, false
	}

	l.len--
	item := l.head.value
	l.head = l.head.next
	return item, true
}

// Peek returns the first item from the list without removing it.
func (l *LinkedList[T]) Peek() (T, bool) {
	if l.len == 0 || l.head == nil {
		var t T
		return t, false
	}
	return l.head.value, true
}

// Get returns the item at the given index.
func (l *LinkedList[T]) Get(index uint) (T, bool) {
	if index >= l.len {
		var t T
		return t, false
	}

	curr := l.head
	for range index {
		curr = curr.next
	}
	return curr.value, true
}

// Delete removes the item at the given index.
func (l *LinkedList[T]) Delete(index uint) (T, bool) {
	if index >= l.len {
		var t T
		return t, false
	}

	if index == 0 {
		return l.Pop()
	}

	curr := l.head
	for range index - 1 {
		curr = curr.next
	}

	l.len--
	deleted := curr.next.value
	curr.next = curr.next.next
	return deleted, true
}

// Len returns the number of items in the list.
func (l *LinkedList[T]) Len() uint {
	return l.len
}

// Locate returns the index of the given item.
// If the item is not found, it returns -1.
func (l *ComparableLinkedList[T]) Locate(item T) int {
	curr := l.head
	var i uint
	for ; i < l.len; i++ {
		if curr.value == item {
			return int(i)
		}
		curr = curr.next
	}
	return -1
}

// Remove removes the first occurrence of the given item.
func (l *ComparableLinkedList[T]) Remove(item T) bool {
	if l.len == 0 || l.head == nil {
		return false
	}

	if l.head.value == item {
		l.Pop()
		return true
	}

	curr := l.head
	for curr.next != nil {
		if curr.next.value == item {
			l.len--
			curr.next = curr.next.next
			return true
		}
		curr = curr.next
	}
	return false
}

// Has returns true if the item is in the list.
func (l *ComparableLinkedList[T]) Has(item T) bool {
	return l.Locate(item) != -1
}
