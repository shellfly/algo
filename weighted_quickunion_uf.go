package algo

// WeightedQuickUnionUF ...
type WeightedQuickUnionUF struct {
	id    []int
	n     int
	size map[int]int
	Count int
}

// NewWeightedQuickUnionUF initialize and return a UF
func NewWeightedQuickUnionUF(n int) *WeightedQuickUnionUF {
	uf := WeightedQuickUnionUF{n: n, Count: n}
	uf.id = make([]int, n)
	uf.size = map[int]int{}
	for i := 0; i < n; i++ {
		uf.id[i] = i
		uf.size[i] = 1
	}
	return &uf
}
// Connected returns if p and q is connected
func (uf *WeightedQuickUnionUF) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// Find return the id of point p
func (uf *WeightedQuickUnionUF) Find(p int) int {
	for ;p!=uf.id[p];p=uf.id[p]{}
	return p
}

// Union unions p and q
func (uf *WeightedQuickUnionUF) Union(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	if uf.size[pRoot] >= uf.size[qRoot] {
		uf.id[qRoot] = pRoot
		uf.size[pRoot]++
	} else {
		uf.id[pRoot] = qRoot
		uf.size[qRoot] += uf.size[pRoot]
	}
	
	uf.Count--
}