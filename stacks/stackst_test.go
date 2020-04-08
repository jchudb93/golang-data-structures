package stacks

import (
	"testing"
)

func TestPush(t *testing.T) {
	stack := NewStack()
	stack.Push(3)
	if stack.top.data != 3 && stack.size != 1 {
		t.Error("Error no push")
	}
}

func TestPop(t *testing.T) {
	stack := NewStack()
	stack.Push(3)
	stack.Push(4)
	stack.Push(2)
	stack.Pop()
	if stack.top.data != 4 && stack.size != 2 {
		t.Error("Error no pop")
	}
}

func TestPeek(t *testing.T) {
	stack := NewStack()
	stack.Push(3)
	stack.Push(4)
	stack.Push(2)
	peek := stack.Peek()
	if peek != 2 {
		t.Error("No peek")
	}
}
