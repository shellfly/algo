package algo

// KruskalMST ...
type KruskalMST struct {
	mst    *Queue
	weight float32
}

// NewKruskalMST ...
func NewKruskalMST(g *EdgeWeightedGraph) *KruskalMST {
	mst := NewQueue()
	k := &KruskalMST{mst, 0}
	pq := NewMinPQ()
	for _, e := range g.Edges() {
		pq.Insert(e)
	}
	uf := NewUF(g.V())
	for !pq.IsEmpty() && mst.Size() < g.V()-1 {
		e := pq.DelMin().(*Edge)
		v := e.Either()
		w := e.Other(v)
		if uf.Connected(v, w) {
			continue
		}
		uf.Union(v, w)
		k.mst.Enqueue(e)
		k.weight += e.Weight()
	}
	return k
}

// Weight ...
func (k *KruskalMST) Weight() float32 {
	return k.weight
}

// Edges ...
func (k *KruskalMST) Edges() (edges []*Edge) {
	for _, e := range k.mst.Slice() {
		edges = append(edges, e.(*Edge))
	}
	return
}
