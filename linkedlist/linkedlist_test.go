package linkedlist_test

import (
	"math/rand"
	"testing"
	"time"

	"single_linked_list/linkedlist"
)

func TestLinkedList(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		t.Run("Add 1 node", func(t *testing.T) {
			ll := linkedlist.NewLinkedList[int]()
			ll.Add(1)

			if ll.Head().Value() != 1 {
				t.Errorf("Expected value to be 1, got %d", ll.Head().Value())
			}
		})

		t.Run("Add 2 nodes", func(t *testing.T) {
			ll := linkedlist.NewLinkedList[int]()
			ll.Add(1)
			ll.Add(2)

			if ll.Head().Value() != 1 {
				t.Errorf("Expected value to be 1, got %d", ll.Head().Value())
			}

			if ll.Head().Next().Value() != 2 {
				t.Errorf("Expected value to be 2, got %d", ll.Head().Next().Value())
			}
		})

		t.Run("Add 3 nodes", func(t *testing.T) {
			ll := linkedlist.NewLinkedList[int]()
			ll.Add(1)
			ll.Add(2)
			ll.Add(3)

			if ll.Head().Value() != 1 {
				t.Errorf("Expected value to be 1, got %d", ll.Head().Value())
			}

			if ll.Head().Next().Value() != 2 {
				t.Errorf("Expected value to be 2, got %d", ll.Head().Next().Value())
			}

			if ll.Head().Next().Next().Value() != 3 {
				t.Errorf("Expected value to be 3, got %d", ll.Head().Next().Next().Value())
			}
		})
	})

	t.Run("Remove", func(t *testing.T) {
		t.Run("Remove 1 node", func(t *testing.T) {
			ll := linkedlist.NewLinkedList[int]()
			ll.Add(1)
			delNode := ll.Head()
			ll.Remove(delNode)

			if ll.Head() != nil {
				t.Errorf("Expected head to be nil, got %v", ll.Head())
			}

			if delNode.Next() != nil {
				t.Errorf("Expected next to be nil, got %v", delNode.Next())
			}
		})

		t.Run("Remove 2 nodes", func(t *testing.T) {
			ll := linkedlist.NewLinkedList[int]()
			ll.Add(1)
			ll.Add(2)
			delNode := ll.Head()
			ll.Remove(delNode)

			if ll.Head().Value() != 2 {
				t.Errorf("Expected value to be 2, got %d", ll.Head().Value())
			}

			if ll.Head().Next() != nil {
				t.Errorf("Expected next to be nil, got %v", ll.Head().Next())
			}

			if delNode.Next() != nil {
				t.Errorf("Expected next to be nil, got %v", delNode.Next())
			}
		})

		t.Run("Remove 3 nodes", func(t *testing.T) {
			ll := linkedlist.NewLinkedList[int]()
			ll.Add(1)
			ll.Add(2)
			ll.Add(3)

			delNode := ll.Head().Next()
			ll.Remove(delNode)

			if ll.Head().Value() != 1 {
				t.Errorf("Expected value to be 1, got %d", ll.Head().Value())
			}

			if ll.Head().Next().Value() != 3 {
				t.Errorf("Expected value to be 3, got %d", ll.Head().Next().Value())
			}

			if ll.Head().Next().Next() != nil {
				t.Errorf("Expected next to be nil, got %v", ll.Head().Next().Next())
			}

			if delNode.Next() != nil {
				t.Errorf("Expected next to be nil, got %v", delNode.Next())
			}
		})

		t.Run("Remove random nodes", func(t *testing.T) {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			ll := linkedlist.NewLinkedList[int]()
			totalNodeCount := r.Int() % 100
			totalNodesToDelete := r.Int() % totalNodeCount

			deleteNodeIdxs := make(map[int]struct{}, totalNodesToDelete)
			for i := 0; i < totalNodesToDelete; i++ {
				deleteNodeIdxs[r.Intn(totalNodeCount)] = struct{}{}
			}
			totalNodesToDelete = len(deleteNodeIdxs)

			delNodes := make([]*linkedlist.Node[int], 0, totalNodesToDelete)
			for i := 0; i < totalNodeCount; i++ {
				newNode := ll.Add(r.Int())
				if _, ok := deleteNodeIdxs[i]; ok {
					delNodes = append(delNodes, newNode)
				}
			}

			for _, delNode := range delNodes {
				ll.Remove(delNode)
			}

			count := 0
			for currNode := ll.Head(); currNode != nil; currNode, count = currNode.Next(), count+1 {
			}
			if count != totalNodeCount-totalNodesToDelete {
				t.Errorf("Expected count to be %d, got %d", totalNodeCount-totalNodesToDelete, count)
			}
		})
	})

	t.Run("InsertAfter", func(t *testing.T) {
		t.Run("InsertAfter 2 nodes", func(t *testing.T) {
			ll := linkedlist.NewLinkedList[int]()
			ll.Add(1)
			ll.Add(2)
			ll.Add(4)
			newNode := linkedlist.NewNode(3)
			nodeVal2 := ll.Head().Next()
			ll.InsertAfter(nodeVal2, newNode.(*linkedlist.Node[int]))

			if ll.Head().Value() != 1 {
				t.Errorf("Expected value to be 1, got %d", ll.Head().Value())
			}

			if ll.Head().Next().Next().Value() != 3 {
				t.Errorf("Expected value to be 3, got %d", ll.Head().Next().Next().Value())
			}

			if nodeVal2.Value() != 2 {
				t.Errorf("Expected value to be 2, got %d", nodeVal2.Value())
			}

			if newNode.Value() != 3 {
				t.Errorf("Expected value to be 3, got %d", newNode.Value())
			}

			if newNode.Next().Value() != 4 {
				t.Errorf("Expected value to be 4, got %d", newNode.Next().Value())
			}

			if newNode.Next().Next() != nil {
				t.Errorf("Expected next to be nil, got %v", newNode.Next().Next())
			}
		})
	})

	t.Run("DeleteAfter", func(t *testing.T) {
		t.Run("DeleteAfter 2 nodes", func(t *testing.T) {
			ll := linkedlist.NewLinkedList[int]()
			ll.Add(1)
			ll.Add(2)
			ll.Add(3)
			ll.Add(4)
			nodeVal2 := ll.Head().Next()
			ll.DeleteAfter(nodeVal2)

			if ll.Head().Value() != 1 {
				t.Errorf("Expected value to be 1, got %d", ll.Head().Value())
			}

			if ll.Head().Next().Value() != 2 {
				t.Errorf("Expected value to be 2, got %d", ll.Head().Next().Value())
			}

			if nodeVal2.Next().Value() != 4 {
				t.Errorf("Expected value to be 4, got %d", ll.Head().Next().Next().Value())
			}

			if ll.Head().Next().Next().Next() != nil {
				t.Errorf("Expected next to be nil, got %v", ll.Head().Next().Next().Next())
			}
		})
	})
}
