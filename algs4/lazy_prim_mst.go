package algs4

// LazyPrimMST ...
type LazyPrimMST struct {
	marked []bool
	mst    *Queue
	pq     *MinPQ
	weight float32
}

// NewLazyPrimMST ...
func NewLazyPrimMST(g *EdgeWeightedGraph) *LazyPrimMST {
	marked := make([]bool, g.V())
	pq := NewMinPQ()
	mst := NewQueue()
	l := &LazyPrimMST{marked, mst, pq, 0}
	for v := 0; v < g.V(); v++ {
		if !l.marked[v] {
			l.prim(g, v)
		}
	}

	return l
}

func (l *LazyPrimMST) prim(g *EdgeWeightedGraph, v int) {
	l.visit(g, v)

	for !l.pq.IsEmpty() {
		e := l.pq.DelMin().(*Edge)
		v := e.Either()
		w := e.Other(v)
		if l.marked[v] && l.marked[w] {
			continue
		}
		l.mst.Enqueue(e)
		l.weight += e.Weight()
		if !l.marked[v] {
			l.visit(g, v)
		}
		if !l.marked[w] {
			l.visit(g, w)
		}
	}

}
func (l *LazyPrimMST) visit(g *EdgeWeightedGraph, v int) {
	l.marked[v] = true
	for _, e := range g.Adj(v) {
		if !l.marked[e.Other(v)] {
			l.pq.Insert(e)
		}
	}
}

// Weight ...
func (l *LazyPrimMST) Weight() float32 {
	return l.weight
}

// Edges ...
func (l *LazyPrimMST) Edges() (edges []*Edge) {
	for _, e := range l.mst.Slice() {
		edges = append(edges, e.(*Edge))
	}
	return
}
