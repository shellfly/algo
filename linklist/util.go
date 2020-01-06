package linklist

// Node is a generic interface for all nodes in a linklist
type Node struct {
	item interface{}
	next *Node
}

// NewNode create a new Node
func NewNode(item interface{}, next *Node) *Node {
	return &Node{item, next}
}

// LinkList ...
type LinkList struct {
	first *Node
}

// Add adds a new node
func (l *LinkList) Add(item interface{}) {
	l.first = NewNode(item, l.first)
}

// Slice returns a slice of link list
func (l *LinkList) Slice() (items []interface{}) {
	for curr := l.first; curr != nil; curr = curr.next {
		items = append(items, curr.item)
	}
	return
}

// NewLinkList return a linked list
func NewLinkList() *LinkList {
	return &LinkList{nil}
}
