package util

import (
	"testing"

	"single_linked_list/linkedlist"
)

func TestReverseRecursiveRecursive(t *testing.T) {
	t.Parallel()

	t.Run("ReverseRecursive empty", func(t *testing.T) {
		ll := linkedlist.NewLinkedList[int]()
		ReverseRecursive(ll.(*linkedlist.LinkedList[int]))

		if ll.Head() != nil {
			t.Errorf("Expected head to be nil, got %v", ll.Head())
		}
	})

	t.Run("ReverseRecursive 1 node", func(t *testing.T) {
		ll := linkedlist.NewLinkedList[int]()
		ll.Add(1)
		ReverseRecursive(ll.(*linkedlist.LinkedList[int]))

		if ll.Head().Value() != 1 {
			t.Errorf("Expected value to be 1, got %d", ll.Head().Value())
		}

		if ll.Head().Next() != nil {
			t.Errorf("Expected next to be nil, got %v", ll.Head().Next())
		}
	})

	t.Run("ReverseRecursive 2 nodes", func(t *testing.T) {
		ll := linkedlist.NewLinkedList[int]()
		ll.Add(1)
		ll.Add(2)
		ReverseRecursive(ll.(*linkedlist.LinkedList[int]))

		if ll.Head().Value() != 2 {
			t.Errorf("Expected value to be 2, got %d", ll.Head().Value())
		}

		if ll.Head().Next().Value() != 1 {
			t.Errorf("Expected value to be 1, got %d", ll.Head().Next().Value())
		}
	})

	t.Run("ReverseRecursive 3 nodes", func(t *testing.T) {
		ll := linkedlist.NewLinkedList[int]()
		ll.Add(1)
		ll.Add(2)
		ll.Add(3)
		ReverseRecursive(ll.(*linkedlist.LinkedList[int]))

		if ll.Head().Value() != 3 {
			t.Errorf("Expected value to be 3, got %d", ll.Head().Value())
		}

		if ll.Head().Next().Value() != 2 {
			t.Errorf("Expected value to be 2, got %d", ll.Head().Next().Value())
		}
	})

	t.Run("ReverseRecursive 4 nodes", func(t *testing.T) {
		ll := linkedlist.NewLinkedList[int]()
		ll.Add(1)
		ll.Add(2)
		ll.Add(3)
		ll.Add(4)
		ReverseRecursive(ll.(*linkedlist.LinkedList[int]))

		if ll.Head().Value() != 4 {
			t.Errorf("Expected value to be 4, got %d", ll.Head().Value())
		}

		if ll.Head().Next().Value() != 3 {
			t.Errorf("Expected value to be 3, got %d", ll.Head().Next().Value())
		}
	})

	t.Run("ReverseRecursive 5 nodes", func(t *testing.T) {
		ll := linkedlist.NewLinkedList[int]()
		ll.Add(1)
		ll.Add(2)
		ll.Add(3)
		ll.Add(4)
		ll.Add(5)
		ReverseRecursive(ll.(*linkedlist.LinkedList[int]))

		if ll.Head().Value() != 5 {
			t.Errorf("Expected value to be 5, got %d", ll.Head().Value())
		}

		if ll.Head().Next().Value() != 4 {
			t.Errorf("Expected value to be 4, got %d", ll.Head().Next().Value())
		}
	})
}
