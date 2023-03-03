package linkedlist

type Noder[T any] interface {
	SetNext(newNode *Node[T])
	Next() *Node[T]
	Value() T
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

var _ Noder[any] = (*Node[any])(nil)

func NewNode[T any](value T) Noder[T] {
	return &Node[T]{value: value}
}

func (n *Node[T]) SetNext(newNode *Node[T]) {
	n.next = newNode
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) Value() T {
	return n.value
}
