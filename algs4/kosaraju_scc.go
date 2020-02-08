package algs4

// KosarajuSCC ...
type KosarajuSCC struct {
	marked []bool
	id     []int
	count  int
}

// NewKosarajuSCC ...
func NewKosarajuSCC(g *Digraph) *KosarajuSCC {
	marked := make([]bool, g.V())
	id := make([]int, g.V())
	k := &KosarajuSCC{marked: marked, id: id}
	order := NewDepthFirstOrder(g.Reverse())
	for _, v := range order.ReversePost().IntSlice() {
		if !k.marked[v] {
			k.Dfs(g, v)
			k.count++
		}
	}
	return k
}

// Dfs ...
func (k *KosarajuSCC) Dfs(g *Digraph, v int) {
	k.marked[v] = true
	k.id[v] = k.count
	for _, w := range g.Adj(v) {
		if !k.marked[w] {
			k.Dfs(g, w)
		}
	}
}

// StronglyConnected ...
func (k *KosarajuSCC) StronglyConnected(v, w int) bool {
	return k.id[v] == k.id[w]
}

// ID ...
func (k *KosarajuSCC) ID(v int) int {
	return k.id[v]
}

// Count ...
func (k *KosarajuSCC) Count() int {
	return k.count
}
