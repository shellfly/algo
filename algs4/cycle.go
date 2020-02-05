package algs4

// Cycle ...
type Cycle struct {
	marked []bool
	edgeTo []int
	cycle  *Stack
}

// NewCycle ...
func NewCycle(g *Graph) *Cycle {
	marked := make([]bool, g.V())
	edgeTo := make([]int, g.V())
	c := &Cycle{marked: marked, edgeTo: edgeTo}
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
func (c *Cycle) hasSelfLoop(g *Graph) bool {
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
func (c *Cycle) hasParallelEdges(g *Graph) bool {
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
func (c *Cycle) Dfs(g *Graph, u, v int) {
	c.marked[v] = true
	for _, w := range g.Adj(v) {
		// short circuit if cycle already found
		if c.cycle != nil {
			return
		}

		if !c.marked[w] {
			c.edgeTo[w] = v
			c.Dfs(g, v, w)
		} else if w != u {
			c.cycle = NewStack()
			for x := v; x != w; x = c.edgeTo[x] {
				c.cycle.Push(x)
			}
			c.cycle.Push(w)
			c.cycle.Push(v)
		}
	}
}

// HasCycle ...
func (c *Cycle) HasCycle() bool {
	return c.cycle != nil
}

// Cycle ...
func (c *Cycle) Cycle() *Stack {
	return c.cycle
}
