package algs4

import (
	"github.com/shellfly/algo/linklist"
)

// Bag is a generic interface for bag
type Bag struct {
	*linklist.LinkList
	n int
}

// NewBag creates a new Bag object
func NewBag() *Bag {
	return &Bag{linklist.NewLinkList(), 0}
}

// Add a new item to bag
func (b *Bag) Add(item interface{}) {
	b.LinkList.Add(item)
	b.n++
}

// Size of the bag
func (b *Bag) Size() int {
	return b.n
}

// IsEmpty returns true if bag is empty
func (b *Bag) IsEmpty() bool {
	return b.n == 0
}
