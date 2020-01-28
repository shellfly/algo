package algo

// MinPQ is a priority queue
type MinPQ struct {
	n  int
	pq []Transaction
}

// NewMinPQ ...
func NewMinPQ(n int) *MinPQ {
	pq := make([]Transaction, n+1)
	return &MinPQ{pq: pq}
}

func (mq MinPQ) less(i, j int) bool {
	return mq.pq[i].Amount < mq.pq[j].Amount
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
	for 2*k < mq.n {
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
func (mq *MinPQ) Insert(t Transaction) {
	mq.n++
	mq.pq[mq.n] = t
	mq.swim(mq.n)
}

// DelMin ...
func (mq *MinPQ) DelMin() Transaction {
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
