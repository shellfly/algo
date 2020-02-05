package algs4

// DepthFirstSearch ...
type DepthFirstSearch struct {
	marked []bool
	count  int
}

// NewDepthFirstSearch ...
func NewDepthFirstSearch(g *Graph, s int) *DepthFirstSearch {
	dfs := &DepthFirstSearch{marked: make([]bool, g.V())}
	dfs.Dfs(g, s)
	return dfs
}

// Dfs ...
func (s *DepthFirstSearch) Dfs(g *Graph, v int) {
	s.marked[v] = true
	s.count++
	for _, w := range g.Adj(v) {
		if !s.marked[w] {
			s.Dfs(g, w)
		}
	}
}

// Marked ...
func (s *DepthFirstSearch) Marked(v int) bool {
	return s.marked[v]
}

// Count ...
func (s *DepthFirstSearch) Count() int {
	return s.count
}
