package linkedlists

import "testing"

func TestAppendToTail(t *testing.T) {
	ll := NewLinkedList()
	ll.AppendToTail(3)
	ll.AppendToTail(4)
	if ll.Length < 2 {
		t.Error("No append")
	}
}
