package stacks

import (
	"errors"
	"fmt"
)

// StackNode node
type StackNode struct {
	data int
	next *StackNode
}

// NewNode new stack node
func NewNode(data int) *StackNode {
	return &StackNode{data, nil}
}

// Stack structure
type Stack struct {
	size int
	top  *StackNode
}

// NewStack creates stack
func NewStack() *Stack {
	return &Stack{0, nil}
}

// Pop from stsack
func (s *Stack) Pop() int {
	if s.top == nil {
		fmt.Println("", errors.New("Empty stack"))
		return -1
	} else {
		item := s.top.data
		top := s.top
		top = top.next
		s.size--
		return item
	}
}

// Push to stack
func (s *Stack) Push(data int) {
	s.top = NewNode(data)
	s.size++
}

// Peek of stack
func (s *Stack) Peek() int {
	if s.top == nil {
		fmt.Println("", errors.New("Empty stack"))
	} else {
		return s.top.data
	}
	return -1
}

// IsEmpty ?
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}
