package algo

import "fmt"

type FloatPQItem float32

func (f FloatPQItem) CompareTo(other interface{}) int {
	ff := other.(FloatPQItem)
	if f < ff {
		return -1
	} else if f > ff {
		return 1
	} else {
		return 0
	}
}
func (f FloatPQItem) String() string {
	return fmt.Sprintf("%.5f", f)
}

// PrimMST ...
type PrimMST struct {
	edgeTo []*Edge
	distTo []float32
	marked []bool
	pq     *IndexMinPQ
}

// NewPrimMST ...
func NewPrimMST(g *EdgeWeightedGraph) *PrimMST {
	edgeTo := make([]*Edge, g.V())
	distTo := make([]float32, g.V())
	marked := make([]bool, g.V())
	pq := NewIndexMinPQ(g.V())
	l := &PrimMST{edgeTo, distTo, marked, pq}
	for v := 0; v < g.V(); v++ {
		l.distTo[v] = 999.0 // fake infinity
	}
	for v := 0; v < g.V(); v++ {
		if !l.marked[v] {
			l.prim(g, v)
		}
	}

	return l
}

func (l *PrimMST) prim(g *EdgeWeightedGraph, s int) {
	l.distTo[s] = 0.0
	l.pq.Insert(s, FloatPQItem(l.distTo[s]))

	for !l.pq.IsEmpty() {
		v := l.pq.DelMin()
		l.visit(g, v)
	}
}
func (l *PrimMST) visit(g *EdgeWeightedGraph, v int) {
	l.marked[v] = true
	for _, e := range g.Adj(v) {
		w := e.Other(v)
		if l.marked[w] {
			continue
		}
		if e.Weight() < l.distTo[w] {
			l.distTo[w] = e.Weight()
			l.edgeTo[w] = e
			if l.pq.Contains(w) {
				l.pq.DecreaseKey(w, FloatPQItem(l.distTo[w]))
			} else {
				l.pq.Insert(w, FloatPQItem(l.distTo[w]))
			}
		}
	}
}

// Weight ...
func (l *PrimMST) Weight() float32 {
	var weight float32
	for _, e := range l.Edges() {
		weight += e.Weight()
	}
	return weight
}

// Edges ...
func (l *PrimMST) Edges() (edges []*Edge) {
	for v := 0; v < len(l.edgeTo); v++ {
		e := l.edgeTo[v]
		if e != nil {
			edges = append(edges, e)
		}
	}
	return
}
