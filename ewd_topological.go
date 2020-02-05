package algo

// EWDTopological implements topological order for a edge weighted digraph
type EWDTopological struct {
	order []*DirectedEdge
}

// NewEWDTopological ...
func NewEWDTopological(g *EdgeWeightedDigraph) *EWDTopological {
	// See topological.go for detail algorithm
	return &EWDTopological{}
}

// Order ...
func (t *EWDTopological) Order() []*DirectedEdge {
	return t.order
}
