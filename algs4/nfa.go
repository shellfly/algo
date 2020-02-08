package algs4

// NFA ...
type NFA struct {
	re string
	g  *Digraph
	m  int
}

// NewNFA ...
func NewNFA(re string) *NFA {
	m := len(re)
	g := NewDigraphV(m + 1)
	nfa := &NFA{re, g, m}
	ops := NewStack()
	for i := 0; i < m; i++ {
		lp := i
		c := string(re[i])
		if c == "(" || c == "|" {
			ops.Push(i)
		} else if c == ")" {
			op := ops.Pop().(int)
			if string(re[op]) == "|" {
				lp = ops.Pop().(int)
				g.AddEdge(lp, op+1)
				g.AddEdge(op, i)
			} else {
				lp = op
			}
		}
		if i < m-1 && string(re[i+1]) == "*" {
			g.AddEdge(lp, i+1)
			g.AddEdge(i+1, lp)
		}

		if c == "(" || c == "*" || c == ")" {
			g.AddEdge(i, i+1)
		}
	}
	return nfa
}

// Recognizes ...
func (n *NFA) Recognizes(txt string) bool {
	pc := NewBag()
	dfs := NewDirectedDFS(n.g, 0)
	for v := 0; v < n.g.V(); v++ {
		if dfs.Marked(v) {
			pc.Add(v)
		}
	}

	for i := 0; i < len(txt); i++ {
		match := NewBag()
		for _, v := range pc.IntSlice() {
			if v < n.m {
				if n.re[v] == txt[i] || string(n.re[v]) == "." {
					match.Add(v + 1)
				}
			}
		}

		pc = NewBag()
		dfs = NewDirectedDFSSources(n.g, match.IntSlice())
		for v := 0; v < n.g.V(); v++ {
			if dfs.Marked(v) {
				pc.Add(v)
			}
		}
	}

	for _, v := range pc.IntSlice() {
		if v == n.m {
			return true
		}
	}
	return false
}
