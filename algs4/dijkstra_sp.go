package algs4

// PositiveInfinity is a fake positive infinity value
const PositiveInfinity = 999999.0

// DijkstraSP ...
type DijkstraSP struct {
	distTo []float32
	edgeTo []*DirectedEdge
	pq     *IndexMinPQ
}

// NewDijkstraSP ...
func NewDijkstraSP(g *EdgeWeightedDigraph, s int) *DijkstraSP {
	edgeTo := make([]*DirectedEdge, g.V())
	distTo := make([]float32, g.V())
	for v := 0; v < g.V(); v++ {
		distTo[v] = PositiveInfinity
	}
	distTo[s] = 0.0
	pq := NewIndexMinPQ(g.V())
	d := &DijkstraSP{distTo, edgeTo, pq}

	pq.Insert(s, FloatPQItem(distTo[s]))
	for !pq.IsEmpty() {
		d.relax(g, pq.DelMin())
	}

	return d

}

func (d *DijkstraSP) relax(g *EdgeWeightedDigraph, v int) {
	for _, e := range g.Adj(v) {
		w := e.To()
		if d.distTo[w] > d.distTo[v]+e.Weight() {
			d.distTo[w] = d.distTo[v] + e.Weight()
			d.edgeTo[w] = e
			if d.pq.Contains(w) {
				d.pq.ChangeKey(w, FloatPQItem(d.distTo[w]))
			} else {
				d.pq.Insert(w, FloatPQItem(d.distTo[w]))
			}
		}
	}
}

// DistTo ...
func (d *DijkstraSP) DistTo(v int) float32 {
	return d.distTo[v]
}

// HasPathTo ...
func (d *DijkstraSP) HasPathTo(v int) bool {
	return d.distTo[v] < PositiveInfinity
}

// PathTo ...
func (d *DijkstraSP) PathTo(v int) (edges []*DirectedEdge) {
	if !d.HasPathTo(v) {
		return nil
	}

	for e := d.edgeTo[v]; e != nil; e = d.edgeTo[e.From()] {
		// stack
		edges = append([]*DirectedEdge{e}, edges...)
	}
	return
}
