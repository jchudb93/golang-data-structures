package linkedlists

import (
	"fmt"
	"testing"
)

func TestAppendToTail(t *testing.T) {
	ll := NewLinkedList()
	ll.AppendToTail(3)
	ll.AppendToTail(4)
	if ll.Length < 2 {
		t.Error("No append")
	}
}

func TestDeleteNode(t *testing.T) {
	ll := NewLinkedList()
	ll.AppendToTail(3)
	ll.AppendToTail(4)
	ll.AppendToTail(2)
	ll.DeleteNode(3)
	fmt.Printf("%d", ll.head.data)
	if ll.Length > 2 {
		t.Error("No delete")
	}
}
