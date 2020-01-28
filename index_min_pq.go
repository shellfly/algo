package algo

import "fmt"

// IndexMinPQ is a Minimum-oriented indexed PQ implementation using a binary heap.
type IndexMinPQ struct {
	maxN int      // maximum number of elements on PQ
	n    int      // number of elements on PQ
	pq   []int    // binary heap using 1-based indexing
	qp   []int    // inverse of pq - qp[pq[i]] = pq[qp[i]] = i
	keys []string //  keys[i] = priority of i
}

// NewIndexMinPQ ...
func NewIndexMinPQ(maxN int) *IndexMinPQ {
	pq := make([]int, maxN+1)
	qp := make([]int, maxN)
	for i := 0; i < len(qp); i++ {
		qp[i] = -1
	}
	keys := make([]string, maxN)
	return &IndexMinPQ{maxN, 0, pq, qp, keys}
}

func (mq IndexMinPQ) less(i, j int) bool {
	return mq.keys[mq.pq[i]] < mq.keys[mq.pq[j]]
}

func (mq *IndexMinPQ) exch(i, j int) {
	mq.pq[i], mq.pq[j] = mq.pq[j], mq.pq[i]
	mq.qp[mq.pq[i]], mq.qp[mq.pq[j]] = i, j
}

func (mq *IndexMinPQ) swim(k int) {
	for k > 1 && mq.less(k, k/2) {
		mq.exch(k, k/2)
		k = k / 2
	}
}

func (mq *IndexMinPQ) sink(k int) {
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

func (mq IndexMinPQ) validateIndex(index int) {
	if index < 0 || index >= mq.maxN {
		panic("Invalid Index")
	}
}

// Insert ...
func (mq *IndexMinPQ) Insert(i int, key string) {
	mq.validateIndex(i)
	if mq.Contains(i) {
		panic("Index is alreay in the priority queue")
	}

	mq.n++
	mq.qp[i] = mq.n
	mq.pq[mq.n] = i
	mq.keys[i] = key
	mq.swim(mq.n)
}

// ChangeKey ...
func (mq *IndexMinPQ) ChangeKey(i int, key string) {
	mq.validateIndex(i)
	if mq.Contains(i) {
		panic("Index is alreay in the priority queue")
	}

	mq.keys[i] = key
	n := mq.qp[i]
	mq.swim(n)
	mq.sink(n)
}

// Contains ...
func (mq *IndexMinPQ) Contains(i int) bool {
	mq.validateIndex(i)
	return mq.qp[i] != -1
}

// Delete ...
func (mq *IndexMinPQ) Delete(i int) {
	mq.validateIndex(i)
	if !mq.Contains(i) {
		panic("Index is not in the priority queue")
	}
	index := mq.qp[i]
	mq.exch(index, mq.n)
	mq.n--
	mq.swim(index)
	mq.sink(index)
	mq.keys[i] = ""
	mq.qp[i] = -1
}

// Min ...
func (mq *IndexMinPQ) Min() string {
	return mq.keys[mq.MinIndex()]
}

// MinIndex ...
func (mq *IndexMinPQ) MinIndex() int {
	return mq.pq[1]
}

// DelMin ...
func (mq *IndexMinPQ) DelMin() int {
	mIndex := mq.MinIndex()
	mq.exch(1, mq.n)
	mq.n--
	mq.sink(1)
	mq.qp[mIndex] = -1
	return mIndex
}

// IsEmpty ...
func (mq IndexMinPQ) IsEmpty() bool {
	return mq.n == 0
}

// Size ...
func (mq IndexMinPQ) Size() int {
	return mq.n
}

// Show ...
func (mq IndexMinPQ) Show() {
	fmt.Println("pq", mq.pq)
	fmt.Println("qp", mq.qp)
	fmt.Println("keys", mq.keys)
}
