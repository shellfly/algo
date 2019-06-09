package algo

import (
	"github.com/shellfly/algo/linklist"
)

// Bag is a generic interface for bag
type Bag struct {
	First *linklist.Node
	N     int
}

// NewBag creates a new Bag object
func NewBag() *Bag {
	return &Bag{nil, 0}
}

// Add a new item to bag
func (b *Bag) Add(item interface{}) {
	oldFirst := b.First
	b.First = &linklist.Node{Item: item, Next: oldFirst}
	b.N++
}

// Size of the bag
func (b *Bag) Size() int {
	return b.N
}

// IsEmpty returns true if bag is empty
func (b *Bag) IsEmpty() bool {
	return b.First == nil
}
