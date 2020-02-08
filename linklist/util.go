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
	First *Node
	n     int // size
}

// NewLinkList return a linked list
func NewLinkList() *LinkList {
	return &LinkList{nil, 0}
}

// Add adds a new node to head, points to old head
func (l *LinkList) Add(item interface{}) (node *Node) {
	node = NewNode(item, l.First)
	l.First = node
	l.n++
	return
}

// Del removes head and returns its Item
func (l *LinkList) Del() (item interface{}) {
	item = l.First.Item
	l.First = l.First.Next
	l.n--
	return
}

// Size returns number of nodes in link list
func (l *LinkList) Size() int {
	return l.n
}

// IsEmpty reports if the link list is empty
func (l *LinkList) IsEmpty() bool {
	return l.First == nil
}

// Slice returns a slice of link list for iterating
func (l *LinkList) Slice() (items []interface{}) {
	for curr := l.First; curr != nil; curr = curr.Next {
		items = append(items, curr.Item)
	}
	return
}

// IntSlice returns a slice of int values for iterating
func (l *LinkList) IntSlice() (items []int) {
	for curr := l.First; curr != nil; curr = curr.Next {
		items = append(items, curr.Item.(int))
	}
	return
}

// StringSlice returns a slice of int values for iterating
func (l *LinkList) StringSlice() (items []string) {
	for curr := l.First; curr != nil; curr = curr.Next {
		items = append(items, curr.Item.(string))
	}
	return
}
