package algs4

// HashKey ...
type HashKey interface {
	hash() int
}

// StringHashKey ...
type StringHashKey string

// has to do this manually since go doesn't export the hash function for object
func (s StringHashKey) hash() int {
	R := 31
	hash := 0
	for _, runeValue := range s {
		hash = R*hash + int(runeValue)
	}
	return hash
}

// SeparateChainHashST is symbol table
type SeparateChainHashST struct {
	n  int // size of key value pairs
	m  int // hash table size
	st []*SequentialSearchST
}

// NewSeparateChainHashST ...
func NewSeparateChainHashST(capacity int) *SeparateChainHashST {
	initCapacity := 4
	if capacity == 0 {
		capacity = initCapacity
	}

	st := make([]*SequentialSearchST, capacity)
	for i := 0; i < capacity; i++ {
		st[i] = NewSequentialSearchST()
	}
	return &SeparateChainHashST{m: capacity, st: st}
}

func (st *SeparateChainHashST) resize(chains int) {
	tmpST := NewSeparateChainHashST(chains)
	for i := 0; i < st.m; i++ {
		for _, key := range st.st[i].Keys() {
			tmpST.Put(key.(HashKey), st.st[i].Get(key))
		}
	}
	st.m = tmpST.m
	st.n = tmpST.n
	st.st = tmpST.st
}
func (st *SeparateChainHashST) hash(key HashKey) int {
	return (key.hash() & 0x7fffffff) % st.m
}

// Put add new key value pair into the st
func (st *SeparateChainHashST) Put(key HashKey, val interface{}) {
	if key == nil {
		panic("key is nil")
	}
	if val == nil {
		st.Delete(key)
		return
	}

	// double table size if average length of list >= 10
	if st.n >= 10*st.m {
		st.resize(2 * st.m)
	}

	i := st.hash(key)
	if !st.st[i].Contains(key) {
		st.n++
	}
	st.st[i].Put(key, val)
}

// Get add new key value pair into the st
func (st *SeparateChainHashST) Get(key HashKey) interface{} {
	return st.st[st.hash(key)].Get(key)
}

// GetInt return the val as int( have to do this since GO doesn't have generics)
func (st *SeparateChainHashST) GetInt(key HashKey) int {
	return st.Get(key).(int)
}

// Delete ...
func (st *SeparateChainHashST) Delete(key HashKey) {
	if key == nil {
		panic("key is nil")
	}
	i := st.hash(key)
	if st.st[i].Contains(key) {
		st.n--
	}
	st.st[i].Delete(key)
	// halve table size if average length of list <= 2
	if st.m > 4 && st.n <= 2*st.m {
		st.resize(st.m / 2)
	}
}

// Contains ...
func (st *SeparateChainHashST) Contains(key HashKey) bool {
	i := st.hash(key)
	return st.st[i].Contains(key)
}

// Size ...
func (st *SeparateChainHashST) Size() int {
	return st.n
}

// IsEmpty add new key value pair into the st
func (st SeparateChainHashST) IsEmpty() bool {
	return st.Size() == 0
}

// Keys ...
func (st *SeparateChainHashST) Keys() []interface{} {
	queue := NewQueue()
	for i := 0; i < st.m; i++ {
		for _, key := range st.st[i].Keys() {
			queue.Enqueue(key)
		}
	}
	return queue.Slice()
}
