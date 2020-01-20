package linklist

// Node is a generic interface for all nodes in a linklist
type Node struct {
	Item interface{}
	Next *Node
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
func (l *LinkList) Add(item interface{}) (n *Node) {
	n = NewNode(item, l.first)
	l.first = n
	return
}

// Slice returns a slice of link list for iterating
func (l *LinkList) Slice() (items []interface{}) {
	for curr := l.first; curr != nil; curr = curr.Next {
		items = append(items, curr.Item)
	}
	return
}

// NewLinkList return a linked list
func NewLinkList() *LinkList {
	return &LinkList{nil}
}
