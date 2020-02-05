package algs4

import (
	"strings"

	"github.com/shellfly/algo/stdin"
)

// SymbolGraph ...
type SymbolGraph struct {
	st   ST
	keys []string
	g    *Graph
}

// NewSymbolGraph ...
func NewSymbolGraph(filename, sp string) *SymbolGraph {
	st := NewST()
	in := stdin.NewInByLine(filename)
	for !in.IsEmpty() {
		a := strings.Split(in.ReadString(), sp)
		for i := 0; i < len(a); i++ {
			if !st.Contains(a[i]) {
				st.Put(a[i], st.Size())
			}
		}
	}

	keys := make([]string, st.Size())
	for _, name := range st.Keys() {
		keys[st.Get(name)] = name
	}

	g := NewGraphV(st.Size())
	in = stdin.NewInByLine(filename)
	for !in.IsEmpty() {
		a := strings.Split(in.ReadString(), sp)
		v := st.Get(a[0])
		for i := 1; i < len(a); i++ {
			g.AddEdge(v, st.Get(a[i]))
		}
	}
	return &SymbolGraph{st, keys, g}
}

// Contains ...
func (s *SymbolGraph) Contains(key string) bool {
	return s.st.Contains(key)
}

// Index ...
func (s *SymbolGraph) Index(key string) int {
	return s.st.Get(key)
}

// Name ...
func (s *SymbolGraph) Name(v int) string {
	return s.keys[v]
}

// G ...
func (s *SymbolGraph) G() *Graph {
	return s.g
}
