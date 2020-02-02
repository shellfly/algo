package algo

// DirectedDFS ...
type DirectedDFS struct {
	marked []bool
}

// NewDirectedDFS ...
func NewDirectedDFS(g *Digraph, s int) *DirectedDFS {
	dfs := &DirectedDFS{marked: make([]bool, g.V())}
	dfs.Dfs(g, s)
	return dfs
}

// NewDirectedDFSSources ...
func NewDirectedDFSSources(g *Digraph, sources []int) *DirectedDFS {
	dfs := &DirectedDFS{marked: make([]bool, g.V())}
	for _, s := range sources {
		if !dfs.marked[s] {
			dfs.Dfs(g, s)
		}
	}
	return dfs
}

// Dfs ...
func (s *DirectedDFS) Dfs(g *Digraph, v int) {
	s.marked[v] = true
	for _, w := range g.Adj(v) {
		if !s.marked[w] {
			s.Dfs(g, w)
		}
	}
}

// Marked ...
func (s *DirectedDFS) Marked(v int) bool {
	return s.marked[v]
}
