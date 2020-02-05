package algs4

// AcyclicSP ...
type AcyclicSP struct {
	distTo []float32
	edgeTo []*DirectedEdge
}

// NewAcyclicSP ...
func NewAcyclicSP(g *EdgeWeightedDigraph, s int) *AcyclicSP {
	edgeTo := make([]*DirectedEdge, g.V())
	distTo := make([]float32, g.V())
	for v := 0; v < g.V(); v++ {
		distTo[v] = PositiveInfinity
	}
	distTo[s] = 0.0
	a := &AcyclicSP{distTo, edgeTo}
	top := NewEWDTopological(g)
	for _, e := range top.Order() {
		a.relax(e)
	}
	return a

}

func (a *AcyclicSP) relax(e *DirectedEdge) {
	v := e.From()
	w := e.To()
	if a.distTo[w] > a.distTo[v]+e.Weight() {
		a.distTo[w] = a.distTo[v] + e.Weight()
		a.edgeTo[w] = e
	}
}

// DistTo ...
func (a *AcyclicSP) DistTo(v int) float32 {
	return a.distTo[v]
}

// HasPathTo ...
func (a *AcyclicSP) HasPathTo(v int) bool {
	return a.distTo[v] < PositiveInfinity
}

// PathTo ...
func (a *AcyclicSP) PathTo(v int) (edges []*DirectedEdge) {
	if !a.HasPathTo(v) {
		return nil
	}

	for e := a.edgeTo[v]; e != nil; e = a.edgeTo[e.From()] {
		// stack
		edges = append([]*DirectedEdge{e}, edges...)
	}
	return
}
