package algs4

// LinearProbingHashST is symbol table
type LinearProbingHashST struct {
	n    int       // size of key value pairs
	m    int       // hash table size
	keys []HashKey // defined in separate_chaining_hash_st.go
	vals []interface{}
}

// NewLinearProbingHashST ...
func NewLinearProbingHashST(capacity int) *LinearProbingHashST {
	initCapacity := 4
	if capacity == 0 {
		capacity = initCapacity
	}

	keys := make([]HashKey, capacity)
	vals := make([]interface{}, capacity)
	return &LinearProbingHashST{m: capacity, keys: keys, vals: vals}
}

func (st *LinearProbingHashST) resize(capacity int) {
	tmpST := NewLinearProbingHashST(capacity)
	for i := 0; i < st.m; i++ {
		if st.keys[i] != nil {
			tmpST.Put(st.keys[i], st.vals[i])
		}
	}
	st.keys = tmpST.keys
	st.vals = tmpST.vals
	st.m = tmpST.m
}

func (st *LinearProbingHashST) hash(key HashKey) int {
	return (key.hash() & 0x7fffffff) % st.m
}

// Put add new key value pair into the st
func (st *LinearProbingHashST) Put(key HashKey, val interface{}) {
	if key == nil {
		panic("key is nil")
	}
	if val == nil {
		st.Delete(key)
		return
	}

	// // double table size if 50% full
	if st.n >= st.m/2 {
		st.resize(2 * st.m)
	}

	var i int
	for i = st.hash(key); st.keys[i] != nil; i = (i + 1) % st.m {
		if st.keys[i] == key {
			st.vals[i] = val
			return
		}
	}
	st.keys[i] = key
	st.vals[i] = val
	st.n++
}

// Get add new key value pair into the st
func (st *LinearProbingHashST) Get(key HashKey) interface{} {
	if key == nil {
		panic("key is nil")
	}
	for i := st.hash(key); st.keys[i] != nil; i = (i + 1) % st.m {
		if st.keys[i] == key {
			return st.vals[i]
		}
	}
	return nil
}

// GetInt return the val as int( have to do this since GO doesn't have generics)
func (st *LinearProbingHashST) GetInt(key HashKey) int {
	return st.Get(key).(int)
}

// Delete ...
func (st *LinearProbingHashST) Delete(key HashKey) {
	if key == nil {
		panic("key is nil")
	}
	if !st.Contains(key) {
		return
	}
	i := st.hash(key)
	for ; key != st.keys[i]; i = (i + 1) % st.m {
	}
	st.keys[i] = nil
	st.vals[i] = nil

	// rehash all keys in same cluster
	for j := (i + 1) % st.m; st.keys[i] != nil; j = (j + 1) % st.m {
		key, val := st.keys[j], st.vals[j]
		st.keys[i] = nil
		st.vals[i] = nil
		st.n--
		st.Put(key, val)
	}
	st.n--

	// halves size of array if it's 12.5% full or less
	if st.n > 0 && st.n <= st.m/8 {
		st.resize(st.m / 2)
	}

}

// Contains ...
func (st *LinearProbingHashST) Contains(key HashKey) bool {
	return st.Get(key) != nil
}

// Size ...
func (st *LinearProbingHashST) Size() int {
	return st.n
}

// IsEmpty add new key value pair into the st
func (st LinearProbingHashST) IsEmpty() bool {
	return st.Size() == 0
}

// Keys ...
func (st *LinearProbingHashST) Keys() []interface{} {
	queue := NewQueue()
	for i := 0; i < st.m; i++ {
		if st.keys[i] != nil {
			queue.Enqueue(st.keys[i])
		}
	}
	return queue.Slice()
}
