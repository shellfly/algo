package algs4

// UF ...
type UF struct {
	id    []int
	n     int
	size  map[int]int
	Count int
}

// NewUF initialize and return a UF
func NewUF(n int) *UF {
	uf := UF{n: n, Count: n}
	uf.id = make([]int, n)
	uf.size = map[int]int{}
	for i := 0; i < n; i++ {
		uf.id[i] = i
		uf.size[i] = 1
	}
	return &uf
}

// Connected returns if p and q is connected
func (uf *UF) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// Quick Find
// Find return the id of point p
// func (uf *UF) Find(p int) int {
// 	return uf.id[p]
// }

// // Union unions p and q
// func (uf *UF) Union(p, q int) {
// 	pID := uf.Find(p)
// 	qID := uf.Find(q)
// 	if pID == qID {
// 		return
// 	}
// 	for index, value := range uf.id {
// 		if value == pID {
// 			uf.id[index] = qID
// 		}
// 	}
// 	uf.Count--
// }

// Quick Union

// Find return the id of point p
// func (uf *UF) Find(p int) int {
// 	for ;p!=uf.id[p];p=uf.id[p]{}
// 	return p
// }

// // Union unions p and q
// func (uf *UF) Union(p, q int) {
// 	pRoot := uf.Find(p)
// 	qRoot := uf.Find(q)
// 	if pRoot == qRoot {
// 		return
// 	}
// 	uf.id[pRoot] = qRoot
// 	uf.Count--
// }

// Weighted Quick Union with path compression

// Find ...
func (uf *UF) Find(p int) int {
	for ; p != uf.id[p]; p = uf.id[p] {
		uf.id[p] = uf.id[uf.id[p]] // path compression by halving
	}
	return p
}

// Union unions p and q
func (uf *UF) Union(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	if uf.size[pRoot] > uf.size[qRoot] {
		uf.id[qRoot] = pRoot
		uf.size[pRoot]++
	} else {
		uf.id[pRoot] = qRoot
		uf.size[qRoot]++
	}
	uf.Count--
}
