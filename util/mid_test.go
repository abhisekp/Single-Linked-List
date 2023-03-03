package util

import (
	"testing"

	"single_linked_list/linkedlist"
)

func TestMid(t *testing.T) {
	t.Run("Mid", func(t *testing.T) {
		t.Run("Mid empty", func(t *testing.T) {
			ll := linkedlist.NewLinkedList[int]()
			midNode := Mid(ll.(*linkedlist.LinkedList[int]))

			if midNode != nil {
				t.Errorf("Expected node to be nil, got %v", midNode)
			}
		})

		t.Run("Mid 1 node", func(t *testing.T) {
			ll := linkedlist.NewLinkedList[int]()
			ll.Add(1)
			midNode := Mid(ll.(*linkedlist.LinkedList[int]))

			if midNode.Value() != 1 {
				t.Errorf("Expected value to be 1, got %d", midNode.Value())
			}
		})

		t.Run("Mid 2 nodes", func(t *testing.T) {
			ll := linkedlist.NewLinkedList[int]()
			ll.Add(1)
			ll.Add(2)
			midNode := Mid(ll.(*linkedlist.LinkedList[int]))

			if midNode.Value() != 1 {
				t.Errorf("Expected value to be 1, got %d", midNode.Value())
			}
		})

		t.Run("Mid 3 nodes", func(t *testing.T) {
			ll := linkedlist.NewLinkedList[int]()
			ll.Add(1)
			ll.Add(2)
			ll.Add(3)
			midNode := Mid(ll.(*linkedlist.LinkedList[int]))

			if midNode.Value() != 2 {
				t.Errorf("Expected value to be 2, got %d", midNode.Value())
			}
		})

		t.Run("Mid 4 nodes", func(t *testing.T) {
			ll := linkedlist.NewLinkedList[int]()
			ll.Add(1)
			ll.Add(2)
			ll.Add(3)
			ll.Add(4)
			midNode := Mid(ll.(*linkedlist.LinkedList[int]))

			if midNode.Value() != 2 {
				t.Errorf("Expected value to be 2, got %d", midNode.Value())
			}
		})

		t.Run("Mid 5 nodes", func(t *testing.T) {
			ll := linkedlist.NewLinkedList[int]()
			ll.Add(1)
			ll.Add(2)
			ll.Add(3)
			ll.Add(4)
			ll.Add(5)
			midNode := Mid(ll.(*linkedlist.LinkedList[int]))

			if midNode.Value() != 3 {
				t.Errorf("Expected value to be 3, got %d", midNode.Value())
			}
		})

		t.Run("Mid 6 nodes", func(t *testing.T) {
			ll := linkedlist.NewLinkedList[int]()
			ll.Add(1)
			ll.Add(2)
			ll.Add(3)
			ll.Add(4)
			ll.Add(5)
			ll.Add(6)
			midNode := Mid(ll.(*linkedlist.LinkedList[int]))

			if midNode.Value() != 3 {
				t.Errorf("Expected value to be 3, got %d", midNode.Value())
			}
		})
	})
}
