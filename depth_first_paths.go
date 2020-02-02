package algo

// DepthFirstPaths ...
type DepthFirstPaths struct {
	marked []bool
	edgeTo []int
	s      int
}

// NewDepthFirstPaths ...
func NewDepthFirstPaths(g *Graph, s int) *DepthFirstPaths {
	dfs := &DepthFirstPaths{marked: make([]bool, g.V()), edgeTo: make([]int, g.V()), s: s}
	dfs.Dfs(g, s)
	return dfs
}

// Dfs ...
func (s *DepthFirstPaths) Dfs(g *Graph, v int) {
	s.marked[v] = true
	for _, w := range g.Adj(v) {
		if !s.marked[w] {
			s.edgeTo[w] = v
			s.Dfs(g, w)
		}
	}
}

// HasPathTo ...
func (s *DepthFirstPaths) HasPathTo(v int) bool {
	return s.marked[v]
}

// PathTo ...
func (s *DepthFirstPaths) PathTo(v int) []int {
	if !s.HasPathTo(v) {
		return nil
	}
	path := NewStack()
	for x := v; x != s.s; x = s.edgeTo[x] {
		path.Push(x)
	}
	path.Push(s.s)
	vertexes := []int{}
	for _, x := range path.Slice() {
		vertexes = append(vertexes, x.(int))
	}
	return vertexes
}
