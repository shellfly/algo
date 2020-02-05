package algo

// BellmanFordSP ...
type BellmanFordSP struct {
	distTo []float32
	edgeTo []*DirectedEdge
	onQ    []bool
	queue  *Queue
	cost   int
	cycle  []*DirectedEdge
}

// NewBellmanFordSP ...
func NewBellmanFordSP(g *EdgeWeightedDigraph, s int) *BellmanFordSP {
	edgeTo := make([]*DirectedEdge, g.V())
	distTo := make([]float32, g.V())
	onQ := make([]bool, g.V())
	queue := NewQueue()
	for v := 0; v < g.V(); v++ {
		distTo[v] = PositiveInfinity
	}
	distTo[s] = 0.0
	queue.Enqueue(s)
	onQ[s] = true
	b := &BellmanFordSP{distTo, edgeTo, onQ, queue, 0, nil}
	for !queue.IsEmpty() && !b.hasNegativeCycle() {
		v := queue.Dequeue().(int)
		b.onQ[v] = false
		b.relax(g, v)
	}
	return b
}

func (b *BellmanFordSP) relax(g *EdgeWeightedDigraph, v int) {
	for _, e := range g.Adj(v) {
		w := e.To()
		if b.distTo[w] > b.distTo[v]+e.Weight() {
			b.distTo[w] = b.distTo[v] + e.Weight()
			b.edgeTo[w] = e
			if !b.onQ[w] {
				b.queue.Enqueue(w)
				b.onQ[w] = true
			}
		}
		b.cost++
		if b.cost%g.V() == 0 {
			b.findNegativeCycle()
			if b.hasNegativeCycle() {
				return // found a negative cycle
			}
		}
	}
}

func (b *BellmanFordSP) findNegativeCycle() {
	V := len(b.edgeTo)
	spt := NewEdgeWeightedDigraphV(V)
	for v := 0; v < V; v++ {
		if b.edgeTo[v] != nil {
			spt.AddEdge(b.edgeTo[v])
		}
	}

	finder := NewEdgeWeightedDirectedCycle(spt)
	b.cycle = finder.Cycle()
}
func (b *BellmanFordSP) hasNegativeCycle() bool {
	return b.cycle != nil
}

// DistTo ...
func (b *BellmanFordSP) DistTo(v int) float32 {
	return b.distTo[v]
}

// HasPathTo ...
func (b *BellmanFordSP) HasPathTo(v int) bool {
	return b.distTo[v] < PositiveInfinity
}

// PathTo ...
func (b *BellmanFordSP) PathTo(v int) (edges []*DirectedEdge) {
	if !b.HasPathTo(v) {
		return nil
	}

	for e := b.edgeTo[v]; e != nil; e = b.edgeTo[e.From()] {
		// stack
		edges = append([]*DirectedEdge{e}, edges...)
	}
	return
}
