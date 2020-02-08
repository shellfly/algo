package algs4

import "github.com/shellfly/algo/linklist"

// Stack implements a stack by linklist
type Stack struct {
	*linklist.LinkList
}

// NewStack ...
func NewStack() *Stack {
	return &Stack{linklist.NewLinkList()}
}

// Push ...
func (s *Stack) Push(item interface{}) {
	s.Add(item)
}

// Pop ...
func (s *Stack) Pop() (item interface{}) {
	if s.IsEmpty() {
		panic("Stack Empty")
	}
	item = s.Del()
	return
}
