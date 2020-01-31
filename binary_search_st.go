package algo

const initCapacity = 2

// Key is an interface of the key in ST
type Key interface {
	compareTo(interface{}) int
}

// StringKey implements Key
type StringKey string

// CompareTo ...
func (k StringKey) compareTo(other interface{}) int {
	if k < other.(StringKey) {
		return -1
	} else if k > other.(StringKey) {
		return 1
	}
	return 0
}

// BinarySearchST is symbol table
type BinarySearchST struct {
	keys []Key
	vals []interface{}
	n    int
}

// NewBinarySearchST returns an bst with init capcity
func NewBinarySearchST() *BinarySearchST {
	return &BinarySearchST{
		keys: make([]Key, initCapacity),
		vals: make([]interface{}, initCapacity),
	}
}

func (st *BinarySearchST) resize(capacity int) {
	newKeys := make([]Key, capacity)
	copy(newKeys, st.keys)
	st.keys = newKeys

	newVals := make([]interface{}, capacity)
	copy(newVals, st.vals)
	st.vals = newVals
}

func (st BinarySearchST) rank(key Key) int {
	lo, hi := 0, st.n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		cmp := key.compareTo(st.keys[mid])

		if cmp < 0 {
			hi = mid - 1
		} else if cmp > 0 {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return lo
}

// Put add new key value pair into the st
func (st *BinarySearchST) Put(key Key, val interface{}) {
	i := st.rank(key)
	if i < st.n && st.keys[i] == key {
		st.vals[i] = val
		return
	}

	if st.n == len(st.keys) {
		st.resize(2 * len(st.keys))
	}

	for j := st.n; j > i; j-- {
		st.keys[j], st.vals[j] = st.keys[j-1], st.vals[j-1]
	}
	st.keys[i], st.vals[i] = key, val

	st.n++
}

// Get add new key value pair into the st
func (st *BinarySearchST) Get(key Key) interface{} {
	i := st.rank(key)
	if i < st.n && key.compareTo(st.keys[i]) == 0 {
		return st.vals[i]
	}
	return nil
}

// GetInt return the val as int( have to do this since GO doesn't have generics)
func (st *BinarySearchST) GetInt(key Key) int {
	return st.Get(key).(int)
}

// Delete ...
func (st *BinarySearchST) Delete(key Key) {
	st.Put(key, nil)
}

// Contains ...
func (st *BinarySearchST) Contains(key Key) bool {
	return st.Get(key) != nil

}

// Size ...
func (st *BinarySearchST) Size() int {
	return st.n
}

// IsEmpty add new key value pair into the st
func (st BinarySearchST) IsEmpty() bool {
	return st.Size() == 0
}

// Keys ...
func (st *BinarySearchST) Keys() []Key {
	return st.keys[:st.n]
}
