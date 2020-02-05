package algs4

type node struct {
	key, val interface{}
	next     *node
}

// SequentialSearchST is symbol table
type SequentialSearchST struct {
	first *node
	n     int
}

// NewSequentialSearchST ...
func NewSequentialSearchST() *SequentialSearchST {
	return &SequentialSearchST{}
}

// Put add new key value pair into the st
func (st *SequentialSearchST) Put(key, val interface{}) {
	for x := st.first; x != nil; x = x.next {
		if x.key == key {
			x.val = val
			return
		}
	}
	st.first = &node{key, val, st.first}
	st.n++
}

// Get add new key value pair into the st
func (st *SequentialSearchST) Get(key interface{}) interface{} {
	for x := st.first; x != nil; x = x.next {
		if x.key == key {
			return x.val
		}
	}
	return nil
}

// GetInt return the val as int( have to do this since GO doesn't have generics)
func (st *SequentialSearchST) GetInt(key interface{}) int {
	return st.Get(key).(int)
}

// Delete ...
func (st *SequentialSearchST) Delete(key interface{}) {
	st.Put(key, nil)
}

// Contains ...
func (st *SequentialSearchST) Contains(key interface{}) bool {
	return st.Get(key) != nil

}

// Size ...
func (st *SequentialSearchST) Size() int {
	return st.n
}

// IsEmpty add new key value pair into the st
func (st SequentialSearchST) IsEmpty() bool {
	return st.Size() == 0
}

// Keys ...
func (st *SequentialSearchST) Keys() []interface{} {
	var keys []interface{}
	for x := st.first; x != nil; x = x.next {
		keys = append(keys, x.key)
	}
	return keys
}
