package algs4

const radix = 256

type trieNode struct {
	val  interface{}
	next []*trieNode
}

func newNode() *trieNode {
	return &trieNode{next: make([]*trieNode, radix)}
}

// TrieST ...
type TrieST struct {
	R    int
	n    int // size
	root *trieNode
}

// NewTrieST ...
func NewTrieST() *TrieST {
	return &TrieST{radix, 0, nil}
}

// Get ...
func (t *TrieST) Get(key string) interface{} {
	x := t.get(t.root, key, 0)
	if x == nil {
		return nil
	}
	return x.val
}

func (t *TrieST) get(x *trieNode, key string, d int) *trieNode {
	if x == nil {
		return nil
	}
	if d == len(key) {
		return x
	}
	c := key[d]
	return t.get(x.next[c], key, d+1)
}

// Put ...
func (t *TrieST) Put(key string, val interface{}) {
	t.root = t.put(t.root, key, val, 0)
}

func (t *TrieST) put(x *trieNode, key string, val interface{}, d int) *trieNode {
	if x == nil {
		x = newNode()
	}
	if d == len(key) {
		if x.val == nil {
			t.n++
		}
		x.val = val
		return x
	}
	c := key[d]
	x.next[c] = t.put(x.next[c], key, val, d+1)
	return x
}

// Keys ...
func (t *TrieST) Keys() []string {
	return t.KeysWithPrefix("")
}

// KeysWithPrefix ...
func (t *TrieST) KeysWithPrefix(pre string) (keys []string) {
	q := NewQueue()
	t.collect(t.get(t.root, pre, 0), pre, q)
	for _, item := range q.Slice() {
		keys = append(keys, item.(string))
	}
	return
}

func (t *TrieST) collect(x *trieNode, pre string, q *Queue) {
	if x == nil {
		return
	}
	if x.val != nil {
		q.Enqueue(pre)
	}
	for c := 0; c < t.R; c++ {
		t.collect(x.next[c], pre+string(c), q)
	}
}

// KeysThatMatch ...
func (t *TrieST) KeysThatMatch(pat string) (keys []string) {
	q := NewQueue()
	t.collectMatch(t.root, "", pat, q)
	for _, item := range q.Slice() {
		keys = append(keys, item.(string))
	}
	return
}

func (t *TrieST) collectMatch(x *trieNode, pre, pat string, q *Queue) {
	if x == nil {
		return
	}
	d := len(pre)
	if d == len(pat) && x.val != nil {
		q.Enqueue(pre)
	}
	if d == len(pat) {
		return
	}
	next := pat[d]
	for c := 0; c < t.R; c++ {
		if string(next) == "." || int(next) == c {
			t.collectMatch(x.next[c], pre+string(c), pat, q)
		}
	}
}

// LongPrefixOf ...
func (t *TrieST) LongPrefixOf(s string) string {
	length := t.search(t.root, s, 0, 0)
	return s[:length]
}
func (t *TrieST) search(x *trieNode, s string, d, length int) int {
	if x == nil {
		return length
	}
	if x.val != nil {
		length = d
	}
	if d == len(s) {
		return length
	}
	c := s[d]
	return t.search(x.next[c], s, d+1, length)
}

// Delete ...
func (t *TrieST) Delete(key string) {
	t.root = t.delete(t.root, key, 0)
}

func (t *TrieST) delete(x *trieNode, key string, d int) *trieNode {
	if x == nil {
		return nil
	}
	if d == len(key) {
		if x.val != nil {
			t.n--
		}
		x.val = nil
	} else {
		c := key[d]
		x.next[c] = t.delete(x.next[c], key, d+1)
	}
	if x.val == nil {
		return x
	}
	for c := 0; c < t.R; c++ {
		if x.next[c] != nil {
			return x
		}
	}
	return nil
}

// Size ...
func (t *TrieST) Size() int {
	return t.n
}
