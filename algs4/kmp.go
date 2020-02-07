package algs4

// KMP ...
type KMP struct {
	pattern string
	dfa     [][]int
}

// makeTwoD return a two dimensional slice
func makeTwoD(rows, columns int) [][]int {
	twoD := make([][]int, rows)
	for i := 0; i < rows; i++ {
		twoD[i] = make([]int, columns)
	}
	return twoD
}

// NewKMP ...
func NewKMP(pattern string) *KMP {
	R := 256
	M := len(pattern)
	dfa := makeTwoD(R, M)
	k := &KMP{pattern, dfa}
	dfa[pattern[0]][0] = 1
	X := 0
	for j := 1; j < M; j++ {
		for c := 0; c < R; c++ {
			dfa[c][j] = dfa[c][X]
		}
		dfa[pattern[j]][j] = j + 1
		X = dfa[pattern[j]][X]
	}
	return k
}

// Search ...
func (k *KMP) Search(txt string) int {
	N, M := len(txt), len(k.pattern)
	var i, j int
	for i, j = 0, 0; i < N && j < M; i++ {
		j = k.dfa[txt[i]][j]
	}
	// Found (hit end of pattern)
	if j == M {
		return i - M
	}
	// Not Found (hit end of text)
	return N
}
