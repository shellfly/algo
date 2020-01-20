package algo

import "github.com/shellfly/algo/linklist"

// Stack ...
type Stack struct {
	first *linklist.Node
	n     int
}

// NewStack ...
func NewStack() *Stack {
	return &Stack{}
}

// IsEmpty ...
func (s Stack) IsEmpty() bool {
	return s.first == nil
}

// Push ...
func (s *Stack) Push(item interface{}) {
	s.first = linklist.NewNode(item, s.first)
	s.n++
}

// Pop ...
func (s *Stack) Pop() (item interface{}) {
	if s.IsEmpty() {
		panic("Stack Empty")
	}
	item = s.first.Item
	s.first = s.first.Next
	return
}

// Size ...
func (s Stack) Size() int {
	return s.n
}
