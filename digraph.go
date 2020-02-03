package algo

import (
	"fmt"

	"github.com/shellfly/algo/stdin"
)

// Digraph ...
type Digraph struct {
	v, e int
	adj  []*Bag
}

// NewDigraph ...
func NewDigraph(in *stdin.In) *Digraph {
	v := in.ReadInt()
	adj := make([]*Bag, v)
	for i := 0; i < v; i++ {
		adj[i] = NewBag()
	}
	g := &Digraph{v: v, e: 0, adj: adj}
	e := in.ReadInt()
	for i := 0; i < e; i++ {
		v, w := in.ReadInt(), in.ReadInt()
		g.AddEdge(v, w)
	}
	return g
}

// NewDigraphV ...
func NewDigraphV(v int) *Digraph {
	adj := make([]*Bag, v)
	for i := 0; i < v; i++ {
		adj[i] = NewBag()
	}
	return &Digraph{v: v, e: 0, adj: adj}
}

// V ...
func (g *Digraph) V() int {
	return g.v
}

// E ...
func (g *Digraph) E() int {
	return g.e
}

// AddEdge ...
func (g *Digraph) AddEdge(v, w int) {
	g.adj[v].Add(w)
	g.e++
}

// Adj ...
func (g *Digraph) Adj(v int) []int {
	var items []int
	for _, item := range g.adj[v].Slice() {
		items = append(items, item.(int))
	}
	return items
}

// Reverse ...
func (g *Digraph) Reverse() *Digraph {
	r := NewDigraphV(g.v)
	for v := 0; v < g.v; v++ {
		for _, w := range g.Adj(v) {
			r.AddEdge(w, v)
		}
	}
	return r
}

func (g *Digraph) String() string {
	s := fmt.Sprintf("%d vertices, %d edges\n", g.v, g.e)
	for i := 0; i < g.v; i++ {
		var adjs string
		for _, w := range g.Adj(i) {
			adjs = adjs + fmt.Sprintf(" %d", w)
		}
		s += fmt.Sprintf("%d: %s\n", i, adjs)
	}
	return s
}
