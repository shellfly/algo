package algs4

// Topological ...
type Topological struct {
	order []int
}

// NewTopological ...
func NewTopological(g *Digraph) *Topological {
	t := new(Topological)
	cycleFinder := NewDirectedCycle(g)
	if !cycleFinder.HasCycle() {
		dfs := NewDepthFirstOrder(g)
		order := []int{}
		for _, v := range dfs.ReversePost().Slice() {
			order = append(order, v.(int))
		}
		t.order = order
	}
	return t
}

// Order ...
func (t *Topological) Order() []int {
	return t.order
}

// IsDAG ...
func (t *Topological) IsDAG() bool {
	return t.order != nil
}
