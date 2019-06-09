package linklist

// Node is a generic interface for all nodes in a linklist
type Node struct {
	Item interface{}
	Next *Node
}
