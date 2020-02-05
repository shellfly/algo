package algo

// EdgeWeightedDirectedCycle for EdgeWeightedDigrapy
// Note: this is file is just placeholder, see detail algorithm in directed_cycle.go
type EdgeWeightedDirectedCycle struct {
	marked  []bool
	edgeTo  []int
	cycle   []*DirectedEdge
	onStack []bool
}

// NewEdgeWeightedDirectedCycle ...
func NewEdgeWeightedDirectedCycle(g *EdgeWeightedDigraph) *EdgeWeightedDirectedCycle {
	marked := make([]bool, g.V())
	edgeTo := make([]int, g.V())
	onStack := make([]bool, g.V())
	c := &EdgeWeightedDirectedCycle{marked: marked, edgeTo: edgeTo, onStack: onStack}
	if c.hasSelfLoop(g) {
		return c
	}
	if c.hasParallelEdges(g) {
		return c
	}

	for v := 0; v < g.V(); v++ {
		if !c.marked[v] {
			c.Dfs(g, -1, v)
		}
	}
	return c
}

// hasSelfLoop ...
func (c *EdgeWeightedDirectedCycle) hasSelfLoop(g *EdgeWeightedDigraph) bool {

	return false
}

// hasParallelEdges ...
func (c *EdgeWeightedDirectedCycle) hasParallelEdges(g *EdgeWeightedDigraph) bool {

	return false
}

// Dfs ...
func (c *EdgeWeightedDirectedCycle) Dfs(g *EdgeWeightedDigraph, u, v int) {

}

// HasCycle ...
func (c *EdgeWeightedDirectedCycle) HasCycle() bool {
	return c.cycle != nil
}

// Cycle ...
func (c *EdgeWeightedDirectedCycle) Cycle() []*DirectedEdge {
	return c.cycle
}
