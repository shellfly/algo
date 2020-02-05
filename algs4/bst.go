package algs4

// BSTNode ...
type BSTNode struct {
	key         Key // defined in st.go
	val         interface{}
	left, right *BSTNode
	n           int
}

// BST is symbol table implemented by a binary tree
type BST struct {
	root *BSTNode
}

// NewBST returns an bst with init capcity
func NewBST() *BST {
	return &BST{}
}

func (st *BST) put(x *BSTNode, key Key, val interface{}) *BSTNode {
	if x == nil {
		return &BSTNode{key: key, val: val, n: 1}
	}
	cmp := key.compareTo(x.key)
	if cmp > 0 {
		x.right = st.put(x.right, key, val)
	} else if cmp < 0 {
		x.left = st.put(x.left, key, val)
	} else {
		x.val = val
	}
	x.n = st.size(x.left) + st.size(x.right) + 1
	return x
}

// Put add new key value pair into the st
func (st *BST) Put(key Key, val interface{}) {
	st.root = st.put(st.root, key, val)
	return

}

func (st *BST) get(x *BSTNode, key Key) interface{} {
	if x == nil {
		return nil
	}
	cmp := key.compareTo(x.key)
	if cmp > 0 {
		return st.get(x.right, key)
	} else if cmp < 0 {
		return st.get(x.left, key)
	} else {
		return x.val
	}
}

// Get add new key value pair into the st
func (st *BST) Get(key Key) interface{} {
	return st.get(st.root, key)
}

// GetInt return the val as int( have to do this since GO doesn't have generics)
func (st *BST) GetInt(key Key) int {
	return st.Get(key).(int)
}

// MinNode returns the minimum node in the BST
func (st *BST) MinNode() *BSTNode {
	if st.IsEmpty() {
		panic("call Min on empty bst")
	}
	return st.min(st.root)
}

// Min returns the minimum key in the BST
func (st *BST) Min() Key {
	if st.IsEmpty() {
		panic("call Min on empty bst")
	}
	return st.min(st.root).key
}

// min returns the minium node in x
func (st *BST) min(x *BSTNode) *BSTNode {
	if x.left == nil {
		return x
	}
	return st.min(x.left)
}

// MaxNode returns the maximum node in the BST
func (st *BST) MaxNode() *BSTNode {
	if st.IsEmpty() {
		panic("call Max on empty bst")
	}
	return st.max(st.root)
}

// Max returns the maximum key in the BST
func (st *BST) Max() Key {
	if st.IsEmpty() {
		panic("call Max on empty bst")
	}
	return st.max(st.root).key
}

// max returns the maxium node in x
func (st *BST) max(x *BSTNode) *BSTNode {
	if x.right == nil {
		return x
	}
	return st.max(x.right)
}

// Floor ...
func (st *BST) Floor(key Key) Key {
	x := st.floor(st.root, key)
	if x == nil {
		return nil
	}
	return x.key
}
func (st *BST) floor(x *BSTNode, key Key) *BSTNode {
	if x == nil {
		return nil
	}
	cmp := key.compareTo(x.key)
	if cmp == 0 {
		return x
	} else if cmp < 0 {
		return st.floor(x.left, key)
	}
	t := st.floor(x.right, key)
	if t != nil {
		return t
	}
	return x
}

// Select ...
func (st *BST) Select(k int) Key {
	return st._select(st.root, k).key
}
func (st *BST) _select(x *BSTNode, k int) *BSTNode {
	if x == nil {
		return nil
	}
	t := st.size(x)
	if t > k {
		return st._select(x.left, k)
	} else if t < k {
		return st._select(x.right, k-t-1)
	} else {
		return x
	}
}

// Rank ...
func (st *BST) Rank(k Key) int {
	return st.rank(st.root, k)
}
func (st *BST) rank(x *BSTNode, key Key) int {
	if x == nil {
		return 0
	}
	cmp := key.compareTo(x.key)
	if cmp < 0 {
		return st.rank(x.left, key)
	} else if cmp > 0 {
		return 1 + st.size(x.left) + st.rank(x.right, key)
	} else {
		return st.size(x.left)
	}
}

// DeleteMin ...
func (st *BST) DeleteMin() {
	st.root = st.deleteMin(st.root)
}

func (st *BST) deleteMin(x *BSTNode) *BSTNode {
	if x.left == nil {
		return x.right
	}
	x.left = st.deleteMin(x.left)
	x.n = st.size(x.left) + st.size(x.right) + 1
	return x
}

// Delete ...
func (st *BST) Delete(key Key) {
	st.root = st.delete(st.root, key)
}

func (st *BST) delete(x *BSTNode, key Key) *BSTNode {
	if x == nil {
		return nil
	}
	cmp := key.compareTo(x.key)
	if cmp < 0 {
		x.left = st.delete(x.left, key)
	} else if cmp > 0 {
		x.right = st.delete(x.right, key)
	} else {
		if x.right == nil {
			return x.left
		} else if x.left == nil {
			return x.right
		}
		t := x
		x = st.min(t.right)
		x.right = st.deleteMin(t.right)
		x.left = t.left
	}
	x.n = st.size(x.left) + st.size(x.right) + 1
	return x
}

// Contains ...
func (st *BST) Contains(key Key) bool {
	return st.Get(key) != nil

}

func (st *BST) size(x *BSTNode) int {
	if x == nil {
		return 0
	}
	return x.n
}

// Size ...
func (st *BST) Size() int {
	return st.size(st.root)
}

// IsEmpty add new key value pair into the st
func (st BST) IsEmpty() bool {
	return st.Size() == 0
}

// Keys ...
func (st *BST) Keys() []Key {
	return st.keys(st.Min(), st.Max())
}
func (st *BST) keys(lo, hi Key) (keys []Key) {
	queue := NewQueue()
	st.nodeKeys(st.root, queue, lo, hi)

	for _, item := range queue.Slice() {
		keys = append(keys, item.(Key))
	}
	return
}

func (st *BST) nodeKeys(x *BSTNode, queue *Queue, lo, hi Key) {
	if x == nil {
		return
	}
	cmplo := lo.compareTo(x.key)
	cmphi := hi.compareTo(x.key)
	if cmplo < 0 {
		st.nodeKeys(x.left, queue, lo, hi)
	}
	if cmplo <= 0 && cmphi >= 0 {
		queue.Enqueue(x.key)
	}
	if cmphi > 0 {
		st.nodeKeys(x.right, queue, lo, hi)
	}
}
