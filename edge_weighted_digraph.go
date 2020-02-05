package algo

import (
	"fmt"

	"github.com/shellfly/algo/stdin"
)

// EdgeWeightedDigraph ...
type EdgeWeightedDigraph struct {
	v, e int
	adj  []*Bag
}

// NewEdgeWeightedDigraph ...
func NewEdgeWeightedDigraph(in *stdin.In) *EdgeWeightedDigraph {
	v := in.ReadInt()
	adj := make([]*Bag, v)
	for i := 0; i < v; i++ {
		adj[i] = NewBag()
	}
	g := &EdgeWeightedDigraph{v: v, e: 0, adj: adj}
	e := in.ReadInt()
	for i := 0; i < e; i++ {
		v, w, weight := in.ReadInt(), in.ReadInt(), in.ReadFloat32()
		g.AddEdge(NewDirectedEdge(v, w, weight))
	}
	return g
}

// NewEdgeWeightedDigraphV ...
func NewEdgeWeightedDigraphV(v int) *EdgeWeightedDigraph {
	adj := make([]*Bag, v)
	for i := 0; i < v; i++ {
		adj[i] = NewBag()
	}
	return &EdgeWeightedDigraph{v: v, e: 0, adj: adj}
}

// V ...
func (g *EdgeWeightedDigraph) V() int {
	return g.v
}

// E ...
func (g *EdgeWeightedDigraph) E() int {
	return g.e
}

// AddEdge ...
func (g *EdgeWeightedDigraph) AddEdge(e *DirectedEdge) {
	g.adj[e.From()].Add(e.To())
	g.e++
}

// Adj ...
func (g *EdgeWeightedDigraph) Adj(v int) []*DirectedEdge {
	var items []*DirectedEdge
	for _, item := range g.adj[v].Slice() {
		items = append(items, item.(*DirectedEdge))
	}
	return items
}

// Edges ...
func (g *EdgeWeightedDigraph) Edges() (edges []*DirectedEdge) {
	for v := 0; v < g.v; v++ {
		for _, e := range g.Adj(v) {
			edges = append(edges, e)
		}
	}
	return
}

func (g *EdgeWeightedDigraph) String() string {
	s := fmt.Sprintf("%d vertices, %d edges\n", g.v, g.e)
	for i := 0; i < g.v; i++ {
		var adjs string
		for _, e := range g.Adj(i) {
			adjs = fmt.Sprintf("%s %s ", adjs, e)
		}
		s += fmt.Sprintf("%d: %s\n", i, adjs)
	}
	return s
}
