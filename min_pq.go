package algo

// PQItem is an interface of the item in PQ
type PQItem interface {
	CompareTo(interface{}) int
}

// MinPQ is a priority queue
type MinPQ struct {
	n  int
	pq []PQItem
}

// NewMinPQ ...
func NewMinPQ() *MinPQ {
	pq := make([]PQItem, 1)
	return &MinPQ{pq: pq}
}

// NewMinPQN ...
func NewMinPQN(n int) *MinPQ {
	pq := make([]PQItem, n)
	return &MinPQ{pq: pq}
}


func (mq *MinPQ) resize(capacity int) {
	pq := make([]PQItem, capacity)
	copy(pq, mq.pq)
	mq.pq = pq
}

func (mq MinPQ) less(i, j int) bool {
	cmp := mq.pq[i].CompareTo(mq.pq[j])
	return cmp < 0
}

func (mq *MinPQ) exch(i, j int) {
	mq.pq[i], mq.pq[j] = mq.pq[j], mq.pq[i]
}

func (mq *MinPQ) swim(k int) {
	for k > 1 && mq.less(k, k/2) {
		mq.exch(k/2, k)
		k = k / 2
	}
}

func (mq *MinPQ) sink(k int) {
	for 2*k <= mq.n {
		j := k * 2
		if j < mq.n && mq.less(j+1, j) {
			j++
		}
		if !mq.less(j, k) {
			break
		}
		mq.exch(k, j)
		k = j
	}
}

// Insert ...
func (mq *MinPQ) Insert(t PQItem) {
	if (mq.n == len(mq.pq) - 1){
		 mq.resize(2 * len(mq.pq))
	}

	mq.n++
	mq.pq[mq.n] = t
	mq.swim(mq.n)
}

// DelMin ...
func (mq *MinPQ) DelMin() PQItem {
	m := mq.pq[1]
	mq.exch(1, mq.n)
	mq.n--
	mq.sink(1)
	return m
}

// IsEmpty ...
func (mq MinPQ) IsEmpty() bool {
	return mq.n == 0
}

// Size ...
func (mq MinPQ) Size() int {
	return mq.n
}
