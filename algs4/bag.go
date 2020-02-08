package algs4

import (
	"github.com/shellfly/algo/linklist"
)

// Bag implements a bag data type by linklist
type Bag struct {
	*linklist.LinkList
}

// NewBag ...
func NewBag() *Bag {
	return &Bag{linklist.NewLinkList()}
}
