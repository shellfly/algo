package algs4

import (
	"fmt"

	"github.com/shellfly/algo/stdin"
)

// EdgeWeightedGraph ...
type EdgeWeightedGraph struct {
	v, e int
	adj  []*Bag
}

// NewEdgeWeightedGraph ...
func NewEdgeWeightedGraph(in *stdin.In) *EdgeWeightedGraph {
	v := in.ReadInt()
	adj := make([]*Bag, v)
	for i := 0; i < v; i++ {
		adj[i] = NewBag()
	}
	g := &EdgeWeightedGraph{v: v, e: 0, adj: adj}
	e := in.ReadInt()
	for i := 0; i < e; i++ {
		v, w, weight := in.ReadInt(), in.ReadInt(), in.ReadFloat32()
		g.AddEdge(NewEdge(v, w, weight))
	}
	return g
}

// NewEdgeWeightedGraphV ...
func NewEdgeWeightedGraphV(v int) *EdgeWeightedGraph {
	adj := make([]*Bag, v)
	for i := 0; i < v; i++ {
		adj[i] = NewBag()
	}
	return &EdgeWeightedGraph{v: v, e: 0, adj: adj}
}

// V ...
func (g *EdgeWeightedGraph) V() int {
	return g.v
}

// E ...
func (g *EdgeWeightedGraph) E() int {
	return g.e
}

// AddEdge ...
func (g *EdgeWeightedGraph) AddEdge(e *Edge) {
	v := e.Either()
	w := e.Other(v)
	g.adj[v].Add(e)
	g.adj[w].Add(e)
	g.e++
}

// Adj ...
func (g *EdgeWeightedGraph) Adj(v int) []*Edge {
	var items []*Edge
	for _, item := range g.adj[v].Slice() {
		items = append(items, item.(*Edge))
	}
	return items
}

// Edges ...
func (g *EdgeWeightedGraph) Edges() (edges []*Edge) {
	for v := 0; v < g.v; v++ {
		for _, e := range g.Adj(v) {
			if e.Other(v) > v {
				edges = append(edges, e)
			}
		}
	}
	return
}

func (g *EdgeWeightedGraph) String() string {
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
