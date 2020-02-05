package algo

import "fmt"

// DirectedEdge is a ADT for DirectedEdgeWeightedGraph
type DirectedEdge struct {
	v, w   int
	weight float32
}

// NewDirectedEdge ...
func NewDirectedEdge(v, w int, weight float32) *DirectedEdge {
	return &DirectedEdge{v, w, weight}
}

// Weight ...
func (e *DirectedEdge) Weight() float32 {
	return e.weight
}

// From ...
func (e *DirectedEdge) From() int {
	return e.v
}

// To ...
func (e *DirectedEdge) To() int {
	return e.w
}

// String ...
func (e *DirectedEdge) String() string {
	return fmt.Sprintf("%d->%d %.5f", e.v, e.w, e.weight)
}

// CompareTo implements PQItem interface
func (e *DirectedEdge) CompareTo(other interface{}) int {
	ee := other.(*DirectedEdge)
	if e.weight > ee.weight {
		return 1
	} else if e.weight < ee.weight {
		return -1
	} else {
		return 0
	}
}
