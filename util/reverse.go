package util

import "single_linked_list/linkedlist"

func Reverse[T any](ll *linkedlist.LinkedList[T]) {
	if ll.Head() == nil || ll.Head().Next() == nil {
		return
	}

	var prev *linkedlist.Node[T]
	curr := ll.Head()

	for curr != nil {
		next := curr.Next()
		curr.SetNext(prev)
		prev = curr
		curr = next
	}

	ll.SetHead(prev)
}
