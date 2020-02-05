package algs4

import "fmt"

// Edge is a ADT for EdgeWeightedGraph
type Edge struct {
	v, w   int
	weight float32
}

// NewEdge ...
func NewEdge(v, w int, weight float32) *Edge {
	return &Edge{v, w, weight}
}

// Weight ...
func (e *Edge) Weight() float32 {
	return e.weight
}

// Either ...
func (e *Edge) Either() int {
	return e.v
}

// Other ...
func (e *Edge) Other(v int) int {
	if v == e.v {
		return e.w
	} else if v == e.w {
		return e.v
	} else {
		panic("invalid vertex")
	}
}

// String ...
func (e *Edge) String() string {
	return fmt.Sprintf("%d-%d %.5f", e.v, e.w, e.weight)
}

// CompareTo implements PQItem interface
func (e *Edge)CompareTo(other interface{})int{
	ee := other.(*Edge)
	if e.weight > ee.weight {
		return 1
	} else if e.weight < ee.weight {
		return -1 
	} else {
		return 0
	}
}