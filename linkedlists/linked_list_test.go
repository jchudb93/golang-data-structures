package linkedlists

import "testing"

func TestAppendToTail(t *testing.T) {
	ll := NewLinkedList()
	ll.AppendToTail(3)
	if ll.length < 1 {
		t.Error("No append")
	}
}
