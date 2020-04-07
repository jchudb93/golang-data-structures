package linkedlists

import "fmt"

// Node linked list node
type Node struct {
	data int
	next *Node
}

// NewNode creates new node
func NewNode(data int) *Node {
	return &Node{data, nil}
}

// LinkedList structure
type LinkedList struct {
	Length int
	head   *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{0, nil}
}

// AppendToTail appends to tail
func (ll *LinkedList) AppendToTail(d int) {

	end := NewNode(d)

	if ll.Length == 0 {
		ll.head = end
	} else {
		head := ll.head

		for head.next != nil {
			head = head.next
		}

		head.next = end
	}

	ll.Length++
}

// DeleteNode deletes node with value d
func (ll *LinkedList) DeleteNode(d int) {
	if ll.Length == 0 {
		fmt.Println("Empty list")
		return
	} else {
		head := ll.head
		for head.next != nil {
			if head.next.data == d {
				head.next = head.next.next
			}
		}
	}
}
