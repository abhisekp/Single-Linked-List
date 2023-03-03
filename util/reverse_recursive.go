package util

import "single_linked_list/linkedlist"

func ReverseRecursive[T any](ll *linkedlist.LinkedList[T]) {
	reverseRecur(ll, nil, ll.Head())
}

func reverseRecur[T any](ll *linkedlist.LinkedList[T], prev, curr *linkedlist.Node[T]) {
	if curr == nil {
		ll.SetHead(prev)
		return
	}

	next := curr.Next()
	curr.SetNext(prev)
	reverseRecur(ll, curr, next)
}
