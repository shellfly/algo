package algs4

// CC ...
type CC struct {
	marked []bool
	id     []int
	count  int
}

// NewCC ...
func NewCC(g *Graph) *CC {
	cc := &CC{marked: make([]bool, g.V()), id: make([]int, g.V())}
	for s := 0; s < g.V(); s++ {
		if !cc.marked[s] {
			cc.Dfs(g, s)
			cc.count++
		}
	}
	return cc
}

// Dfs ...
func (cc *CC) Dfs(g *Graph, v int) {
	cc.marked[v] = true
	cc.id[v] = cc.count
	for _, w := range g.Adj(v) {
		if !cc.marked[w] {
			cc.Dfs(g, w)
		}
	}
}

// Connected ...
func (cc *CC) Connected(v, w int) bool {
	return cc.id[v] == cc.id[w]
}

// ID ...
func (cc *CC) ID(v int) int {
	return cc.id[v]
}

// Count ...
func (cc *CC) Count() int {
	return cc.count
}
