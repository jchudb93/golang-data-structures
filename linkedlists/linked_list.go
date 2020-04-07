package linkedlists

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
	length int
	head   *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{0, nil}
}

// AppendToTail appends to tail
func (ll *LinkedList) AppendToTail(d int) {
	end := NewNode(d)

	if ll.length == 0 {
		ll.head = end
	} else {
		head := ll.head

		for head.next != nil {
			head = head.next
		}

		head.next = end
	}

	ll.length++
}
