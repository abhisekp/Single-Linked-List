package util

import "single_linked_list/linkedlist"

func Mid[T any](ll *linkedlist.LinkedList[T]) *linkedlist.Node[T] {
	slow, fast := ll.Head(), ll.Head()

	if slow == nil {
		return nil
	}

	for ; fast.Next() != nil && fast.Next().Next() != nil; slow, fast = slow.Next(), fast.Next().Next() {
	}

	return slow
}
