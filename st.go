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

// ST is symbol table
type ST struct{}

// Put add new key value pair into the st
func (st ST) Put(key, value interface{}) {

}

// Get add new key value pair into the st
func (st ST) Get(key interface{}) (value interface{}) {
	return nil
}

// Delete ...
func (st ST) Delete(key interface{}) {
	return
}

// Contains ...
func (st ST) Contains(key interface{}) bool {
	return false

}

// Size ...
func (st ST) Size() int {
	return 0
}

// IsEmpty add new key value pair into the st
func (st ST) IsEmpty() bool {
	return st.Size() == 0
}

// Keys ...
func (st ST) Keys() []interface{} {
	return nil
}
