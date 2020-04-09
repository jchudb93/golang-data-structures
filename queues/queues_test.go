package queues

import "testing"

func TestNewNode(t *testing.T) {
	n := NewNode(1)
	if n.data != 1 {
		t.Error("No node")
	}
}

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	if q.size != 0 {
		t.Error("No new queue")
	}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(4)
	if q.tail.data != 4 {
		t.Error("Did not enqueue")
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue()
	q.Dequeue()
	q.Enqueue(1)
	q.Enqueue(4)
	res := q.Dequeue()
	if res != 1 && q.size != 1 {
		t.Error("Did not dequeue")
	}
	res = q.Dequeue()
	if res != 4 && q.size != 0 {
		t.Error("Did not dequeue")
	}

}
