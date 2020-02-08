package algs4

import (
	"github.com/shellfly/algo/linklist"
)

// Queue ...
type Queue struct {
	first, last *linklist.Node
	n           int
}

// NewQueue ...
func NewQueue() *Queue {
	return &Queue{}
}

// IsEmpty return a bool to indicate if the queue is empty
func (q Queue) IsEmpty() bool {
	return q.first == nil
}

// Enqueue add  a new item into the queue
func (q *Queue) Enqueue(item interface{}) {
	node := linklist.NewNode(item, nil)
	oldLast := q.last
	q.last = node
	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldLast.Next = q.last
	}
	q.n++
}

// Dequeue add a new item into the queue
func (q *Queue) Dequeue() (item interface{}) {
	if q.IsEmpty() {
		panic("Queue Empty")
	}

	item = q.first.Item
	q.first = q.first.Next
	if q.IsEmpty() {
		q.last = nil
	}
	q.n--
	return
}

// Size return the size of the queue
func (q Queue) Size() int {
	return q.n
}

// Slice ...
func (q Queue) Slice() (items []interface{}) {
	for x := q.first; x != nil; x = x.Next {
		items = append(items, x.Item)
	}
	return
}

// IntSlice returns a slice of int values for iterating
func (q *Queue) IntSlice() (items []int) {
	for curr := q.first; curr != nil; curr = curr.Next {
		items = append(items, curr.Item.(int))
	}
	return
}

// StringSlice returns a slice of int values for iterating
func (q *Queue) StringSlice() (items []string) {
	for curr := q.first; curr != nil; curr = curr.Next {
		items = append(items, curr.Item.(string))
	}
	return
}
