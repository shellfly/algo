package algs4

type tstNode struct {
	val              interface{}
	c                byte
	left, mid, right *tstNode
}

func newTSTNode() *tstNode {
	return &tstNode{}
}

// TST ...
type TST struct {
	n    int // size
	root *tstNode
}

// NewTST ...
func NewTST() *TST {
	return &TST{0, nil}
}

// Get ...
func (t *TST) Get(key string) interface{} {
	if len(key) == 0 {
		panic("key must have length >=1")
	}
	x := t.get(t.root, key, 0)
	if x == nil {
		return nil
	}
	return x.val
}

func (t *TST) get(x *tstNode, key string, d int) *tstNode {
	if x == nil {
		return nil
	}
	c := key[d]
	if c < x.c {
		return t.get(x.left, key, d)
	} else if c > x.c {
		return t.get(x.right, key, d)
	} else if d < len(key)-1 {
		return t.get(x.mid, key, d+1)
	} else {
		return x
	}
}

// Contains ...
func (t *TST) Contains(key string) bool {
	return t.Get(key) != nil
}

// Put ...
func (t *TST) Put(key string, val interface{}) {
	if !t.Contains(key) {
		t.n++
	} else if val == nil {
		t.n--
	}
	t.root = t.put(t.root, key, val, 0)
}

func (t *TST) put(x *tstNode, key string, val interface{}, d int) *tstNode {
	c := key[d]
	if x == nil {
		x = newTSTNode()
		x.c = c
	}
	if c < x.c {
		x.left = t.put(x.left, key, val, d)
	} else if c > x.c {
		x.right = t.put(x.right, key, val, d)
	} else if d < len(key)-1 {
		x.mid = t.put(x.mid, key, val, d+1)
	} else {
		x.val = val
	}
	return x
}

// Keys ...
func (t *TST) Keys() (keys []string) {
	queue := NewQueue()
	t.collect(t.root, "", queue)
	for _, item := range queue.Slice() {
		keys = append(keys, item.(string))
	}
	return
}

// KeysWithPrefix ...
func (t *TST) KeysWithPrefix(prefix string) (keys []string) {
	queue := NewQueue()
	x := t.get(t.root, prefix, 0)
	if x == nil {
		return
	}
	if x.val != nil {
		queue.Enqueue(prefix)
	}
	t.collect(x, prefix, queue)
	for _, item := range queue.Slice() {
		keys = append(keys, item.(string))
	}
	return
}

func (t *TST) collect(x *tstNode, prefix string, queue *Queue) {
	if x == nil {
		return
	}
	t.collect(x.left, prefix, queue)
	if x.val != nil {
		queue.Enqueue(prefix + string(x.c))
	}
	t.collect(x.mid, prefix+string(x.c), queue)
	// prefix = prefix[:len(prefix)-1]
	t.collect(x.right, prefix, queue)
}

// KeysThatMatch ...
func (t *TST) KeysThatMatch(pat string) (keys []string) {
	q := NewQueue()
	t.collectMatch(t.root, "", 0, pat, q)
	for _, item := range q.Slice() {
		keys = append(keys, item.(string))
	}
	return
}

func (t *TST) collectMatch(x *tstNode, prefix string, i int, pattern string, queue *Queue) {
	if x == nil {
		return
	}
	c := pattern[i]
	if string(c) == "." || c < x.c {
		t.collectMatch(x.left, prefix, i, pattern, queue)
	}
	if string(c) == "." || c == x.c {
		if i == len(pattern)-1 && x.val != nil {
			queue.Enqueue(prefix + string(x.c))
		}
		if i < len(pattern)-1 {
			t.collectMatch(x.mid, prefix+string(x.c), i+1, pattern, queue)
			// prefix = prefix[:len(prefix)-1]
		}
	}
	if string(c) == "." || c > x.c {
		t.collectMatch(x.right, prefix, i, pattern, queue)
	}
}

// LongPrefixOf ...
func (t *TST) LongPrefixOf(query string) string {
	if len(query) == 0 {
		return ""
	}
	length := 0
	x := t.root
	i := 0
	for x != nil && i < len(query) {
		c := query[i]
		if c < x.c {
			x = x.left
		} else if c > x.c {
			x = x.right
		} else {
			i++
			if x.val != nil {
				length = i
			}
			x = x.mid
		}
	}
	return query[:length]
}

// Size ...
func (t *TST) Size() int {
	return t.n
}
