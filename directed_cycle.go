package algo

// DirectedCycle ...
type DirectedCycle struct {
	marked  []bool
	edgeTo  []int
	cycle   *Stack
	onStack []bool
}

// NewDirectedCycle ...
func NewDirectedCycle(g *Digraph) *DirectedCycle {
	marked := make([]bool, g.V())
	edgeTo := make([]int, g.V())
	onStack := make([]bool, g.V())
	c := &DirectedCycle{marked: marked, edgeTo: edgeTo, onStack: onStack}
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
func (c *DirectedCycle) hasSelfLoop(g *Digraph) bool {
	for v := 0; v < g.V(); v++ {
		for _, w := range g.Adj(v) {
			if v == w {
				c.cycle = NewStack()
				c.cycle.Push(v)
				c.cycle.Push(w)
				return true
			}
		}
	}
	return false
}

// hasParallelEdges ...
func (c *DirectedCycle) hasParallelEdges(g *Digraph) bool {
	c.marked = make([]bool, g.V())
	for v := 0; v < g.V(); v++ {
		for _, w := range g.Adj(v) {
			if c.marked[w] {
				c.cycle = NewStack()
				c.cycle.Push(v)
				c.cycle.Push(w)
				c.cycle.Push(v)
				return true
			}
			c.marked[w] = true
		}

		// reset so marked[v] = false for all v
		for _, w := range g.Adj(v) {
			c.marked[w] = false
		}
	}
	return false
}

// Dfs ...
func (c *DirectedCycle) Dfs(g *Digraph, u, v int) {
	c.marked[v] = true
	c.onStack[v] = true
	for _, w := range g.Adj(v) {
		// short circuit if cycle already found
		if c.HasCycle() {
			return
		}

		if !c.marked[w] {
			c.edgeTo[w] = v
			c.Dfs(g, v, w)
		} else if c.onStack[w] {
			c.cycle = NewStack()
			for x := v; x != w; x = c.edgeTo[x] {
				c.cycle.Push(x)
			}
			c.cycle.Push(w)
			c.cycle.Push(v)
		}
	}
	c.onStack[v] = false
}

// HasCycle ...
func (c *DirectedCycle) HasCycle() bool {
	return c.cycle != nil
}

// Cycle ...
func (c *DirectedCycle) Cycle() *Stack {
	return c.cycle
}
