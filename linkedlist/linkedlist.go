package linkedlist

type LinkedLister[T any] interface {
	SetHead(n *Node[T])
	Head() *Node[T]
	Add(value T) *Node[T]
	Remove(n *Node[T])
	InsertAfter(ref *Node[T], ins *Node[T])
	DeleteAfter(ref *Node[T])
}

type LinkedList[T any] struct {
	head *Node[T]
}

func NewLinkedList[T any]() LinkedLister[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) SetHead(n *Node[T]) {
	ll.head = n
}

func (ll *LinkedList[T]) Head() *Node[T] {
	return ll.head
}

func (ll *LinkedList[T]) Add(value T) *Node[T] {
	newNode := NewNode(value)

	if ll.Head() == nil {
		ll.SetHead(newNode.(*Node[T]))
		return newNode.(*Node[T])
	}

	currNode := ll.Head()
	for ; currNode.Next() != nil; currNode = currNode.Next() {
	}

	currNode.SetNext(newNode.(*Node[T]))

	return newNode.(*Node[T])
}

func (ll *LinkedList[T]) Remove(ref *Node[T]) {
	prevNode := ll.Head()

	if prevNode == nil {
		return
	}

	if prevNode == ref {
		ll.SetHead(ref.Next())
		ref.SetNext(nil)
		return
	}

	for ; prevNode != nil && prevNode.Next() != ref; prevNode = prevNode.Next() {
	}

	if prevNode == nil {
		return
	}
	prevNode.SetNext(ref.Next())
	ref.SetNext(nil)
}

func (ll *LinkedList[T]) InsertAfter(ref *Node[T], ins *Node[T]) {
	if ref == nil || ins == nil {
		return
	}

	refNext := ref.Next()
	ref.SetNext(ins)
	ins.SetNext(refNext)
}

func (ll *LinkedList[T]) DeleteAfter(ref *Node[T]) {
	if ref == nil {
		return
	}

	refNext := ref.Next()
	if refNext == nil {
		return
	}

	ref.SetNext(refNext.Next())
	refNext.SetNext(nil)
}
