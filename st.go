package algo

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

// ST is symbol table implemented by go  map[string]int
type ST map[string]int

// NewST ...
func NewST() ST {
	return make(ST)
}

// Put add new key value pair into the st
func (st ST) Put(key string, val int) {
	st[key] = val
}

// Get add new key value pair into the st
func (st ST) Get(key string) int {
	return st[key]
}

// Delete ...
func (st ST) Delete(key string) {
	delete(st, key)
}

// Contains ...
func (st ST) Contains(key string) bool {
	_, ok := st[key]
	return ok
}

// Size ...
func (st ST) Size() int {
	return len(st)
}

// IsEmpty add new key value pair into the st
func (st ST) IsEmpty() bool {
	return st.Size() == 0
}

// Keys ...
func (st ST) Keys() []string {
	keys := make([]string, 0, len(st))
	for k := range st {
		keys = append(keys, k)
	}
	return keys
}
