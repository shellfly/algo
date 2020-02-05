package algs4

// MaxPQ is a priority queue
type MaxPQ struct {
	n  int
	pq []PQItem
}

// NewMaxPQ ...
func NewMaxPQ(n int) *MaxPQ {
	// index 0 is not used in PQ
	pq := make([]PQItem, n+1)
	return &MaxPQ{pq: pq}
}

func (mq MaxPQ) less(i, j int) bool {
	cmp := mq.pq[i].CompareTo(mq.pq[j])
	return cmp < 0
}

func (mq *MaxPQ) exch(i, j int) {
	mq.pq[i], mq.pq[j] = mq.pq[j], mq.pq[i]
}

func (mq *MaxPQ) swim(k int) {
	for k > 1 && mq.less(k/2, k) {
		mq.exch(k/2, k)
		k = k / 2
	}
}

func (mq *MaxPQ) sink(k int) {
	for 2*k <= mq.n {
		j := k * 2
		if j < mq.n && mq.less(j, j+1) {
			j++
		}
		if !mq.less(k, j) {
			break
		}
		mq.exch(k, j)
		k = j
	}
}

// Insert ...
func (mq *MaxPQ) Insert(t PQItem) {
	mq.pq = append(mq.pq, t)
	mq.n++
	mq.swim(mq.n)
}

// DelMax ...
func (mq *MaxPQ) DelMax() PQItem {
	m := mq.pq[1]
	mq.exch(1, mq.n)
	mq.n--
	mq.sink(1)
	return m
}

// IsEmpty ...
func (mq MaxPQ) IsEmpty() bool {
	return mq.n == 0
}

// Size ...
func (mq MaxPQ) Size() int {
	return mq.n
}
