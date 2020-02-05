package algs4

const (
	red   = true
	black = false
)

// RedBlackBSTNode ...
type RedBlackBSTNode struct {
	key         Key // defined in st.go
	val         interface{}
	left, right *RedBlackBSTNode
	color       bool
	n           int
}

// RedBlackBST is symbol table implemented by a binary tree
type RedBlackBST struct {
	root *RedBlackBSTNode
}

// NewRedBlackBST returns an RedBlackbst with init capcity
func NewRedBlackBST() *RedBlackBST {
	return &RedBlackBST{}
}

func (st *RedBlackBST) put(x *RedBlackBSTNode, key Key, val interface{}) *RedBlackBSTNode {
	if x == nil {
		return &RedBlackBSTNode{key: key, val: val, n: 1, color: red}
	}
	cmp := key.compareTo(x.key)
	if cmp > 0 {
		x.right = st.put(x.right, key, val)
	} else if cmp < 0 {
		x.left = st.put(x.left, key, val)
	} else {
		x.val = val
	}

	// fix-up any right-leaning links
	if st.isRed(x.right) && !st.isRed(x.left) {
		x = st.rotateLeft(x)
	}
	if st.isRed(x.left) && st.isRed(x.left.left) {
		x = st.rotateRight(x)
	}
	if st.isRed(x.left) && st.isRed(x.right) {
		st.flipColors(x)
	}
	x.n = st.size(x.left) + st.size(x.right) + 1
	return x
}

// Put add new key value pair into the st
func (st *RedBlackBST) Put(key Key, val interface{}) {
	st.root = st.put(st.root, key, val)
	st.root.color = black
}

func (st *RedBlackBST) get(x *RedBlackBSTNode, key Key) interface{} {
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
func (st *RedBlackBST) Get(key Key) interface{} {
	return st.get(st.root, key)
}

// GetInt return the val as int( have to do this since GO doesn't have generics)
func (st *RedBlackBST) GetInt(key Key) int {
	return st.Get(key).(int)
}

// MinNode returns the minimum node in the RedBlackBST
func (st *RedBlackBST) MinNode() *RedBlackBSTNode {
	if st.IsEmpty() {
		panic("call Min on empty RedBlackbst")
	}
	return st.min(st.root)
}

// Min returns the minimum key in the RedBlackBST
func (st *RedBlackBST) Min() Key {
	if st.IsEmpty() {
		panic("call Min on empty RedBlackbst")
	}
	return st.min(st.root).key
}

// min returns the minium node in x
func (st *RedBlackBST) min(x *RedBlackBSTNode) *RedBlackBSTNode {
	if x.left == nil {
		return x
	}
	return st.min(x.left)
}

// MaxNode returns the maximum node in the RedBlackBST
func (st *RedBlackBST) MaxNode() *RedBlackBSTNode {
	if st.IsEmpty() {
		panic("call Max on empty RedBlackbst")
	}
	return st.max(st.root)
}

// Max returns the maximum key in the RedBlackBST
func (st *RedBlackBST) Max() Key {
	if st.IsEmpty() {
		panic("call Max on empty RedBlackbst")
	}
	return st.max(st.root).key
}

// max returns the maxium node in x
func (st *RedBlackBST) max(x *RedBlackBSTNode) *RedBlackBSTNode {
	if x.right == nil {
		return x
	}
	return st.max(x.right)
}

// Floor ...
func (st *RedBlackBST) Floor(key Key) Key {
	x := st.floor(st.root, key)
	if x == nil {
		return nil
	}
	return x.key
}
func (st *RedBlackBST) floor(x *RedBlackBSTNode, key Key) *RedBlackBSTNode {
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
func (st *RedBlackBST) Select(k int) Key {
	return st._select(st.root, k).key
}
func (st *RedBlackBST) _select(x *RedBlackBSTNode, k int) *RedBlackBSTNode {
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
func (st *RedBlackBST) Rank(k Key) int {
	return st.rank(st.root, k)
}
func (st *RedBlackBST) rank(x *RedBlackBSTNode, key Key) int {
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
func (st *RedBlackBST) DeleteMin() {
	if st.IsEmpty() {
		panic("BST is empty")
	}
	if !st.isRed(st.root.left) && !st.isRed(st.root.right) {
		st.root.color = red
	}
	st.root = st.deleteMin(st.root)
	if !st.IsEmpty() {
		st.root.color = black
	}
}

func (st *RedBlackBST) moveRedLeft(x *RedBlackBSTNode) *RedBlackBSTNode {
	st.flipColors(x)
	if st.isRed(x.right.left) {
		x.right = st.rotateRight(x.right)
		x = st.rotateLeft(x)
	}
	return x
}
func (st *RedBlackBST) moveRedRight(x *RedBlackBSTNode) *RedBlackBSTNode {
	st.flipColors(x)
	if st.isRed(x.right.left) {
		x = st.rotateRight(x.right)
		st.flipColors(x)
	}
	return x
}
func (st *RedBlackBST) balance(x *RedBlackBSTNode) *RedBlackBSTNode {
	if st.isRed(x.right) {
		x = st.rotateLeft(x)
	}
	if st.isRed(x.left) && st.isRed(x.left.left) {
		x = st.rotateRight(x)
	}
	if st.isRed(x.left) && st.isRed(x.right) {
		st.flipColors(x)
	}
	x.n = st.size(x.left) + st.size(x.right) + 1
	return x
}

func (st *RedBlackBST) deleteMin(x *RedBlackBSTNode) *RedBlackBSTNode {
	if x.left == nil {
		return nil
	}
	if !st.isRed(x.left) && !st.isRed(x.left.left) {
		x = st.moveRedLeft(x)
	}
	x.left = st.deleteMin(x.left)
	return st.balance(x)
}

// Delete ...
func (st *RedBlackBST) Delete(key Key) {
	if !st.isRed(st.root.left) && !st.isRed(st.root.right) {
		st.root.color = red
	}
	st.root = st.delete(st.root, key)
	if !st.IsEmpty() {
		st.root.color = black
	}
}

func (st *RedBlackBST) delete(x *RedBlackBSTNode, key Key) *RedBlackBSTNode {
	if key.compareTo(x.key) < 0 {
		if !st.isRed(x.left) && !st.isRed(x.left.left) {
			x = st.moveRedLeft(x)
		}
		x.left = st.delete(x.left, key)
	} else {
		if st.isRed(x.left) {
			x = st.rotateRight(x)
		}
		if key.compareTo(x.key) == 0 && x.right == nil {
			return nil
		}
		if !st.isRed(x.right) && !st.isRed(x.right.left) {
			x = st.moveRedRight(x)
		}
		if key.compareTo(x.key) == 0 {
			x.val = st.get(x.right, st.min(x.right).key)
			x.key = st.min(x.right).key
			x.right = st.deleteMin(x.right)
		} else {
			x.right = st.delete(x.right, key)
		}
	}
	return st.balance(x)
}

// Contains ...
func (st *RedBlackBST) Contains(key Key) bool {
	return st.Get(key) != nil

}

func (st *RedBlackBST) size(x *RedBlackBSTNode) int {
	if x == nil {
		return 0
	}
	return x.n
}

// Size ...
func (st *RedBlackBST) Size() int {
	return st.size(st.root)
}

// IsEmpty add new key value pair into the st
func (st RedBlackBST) IsEmpty() bool {
	return st.Size() == 0
}

// Keys ...
func (st *RedBlackBST) Keys() []Key {
	return st.keys(st.Min(), st.Max())
}
func (st *RedBlackBST) keys(lo, hi Key) (keys []Key) {
	queue := NewQueue()
	st.nodeKeys(st.root, queue, lo, hi)

	for _, item := range queue.Slice() {
		keys = append(keys, item.(Key))
	}
	return
}

func (st *RedBlackBST) nodeKeys(x *RedBlackBSTNode, queue *Queue, lo, hi Key) {
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

func (st *RedBlackBST) isRed(x *RedBlackBSTNode) bool {
	if x == nil {
		return false
	}
	return x.color
}

func (st *RedBlackBST) rotateLeft(h *RedBlackBSTNode) *RedBlackBSTNode {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = red
	x.n = h.n
	h.n = st.size(h.left) + st.size(h.right) + 1
	return x
}

func (st *RedBlackBST) rotateRight(h *RedBlackBSTNode) *RedBlackBSTNode {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = red
	x.n = h.n
	h.n = st.size(h.left) + st.size(h.right) + 1
	return x
}

func (st *RedBlackBST) flipColors(h *RedBlackBSTNode) {
	h.color = red
	h.left.color = black
	h.right.color = black
}
