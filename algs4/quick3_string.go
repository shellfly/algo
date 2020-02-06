package algs4

// Quick3String ...
type Quick3String struct {
	CUTOFF int // cutoff to insertion sort
	R      int // extended ASCII alphabet size
	aux    []string
}

// NewQuick3String ...
func NewQuick3String(a []string) *Quick3String {
	m := &Quick3String{CUTOFF: 15, R: 256, aux: make([]string, len(a))}
	m.sort(a, 0, len(a)-1, 0)
	return m
}

func (m *Quick3String) insertion(a []string, lo, hi, d int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && a[j][d] < a[j-1][d]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}
func (m *Quick3String) sort(a []string, lo, hi, d int) {
	if hi <= lo+m.CUTOFF {
		m.insertion(a, lo, hi, d)
		return
	}

	lt, gt := lo, hi
	v := a[lo][d]
	for i := lo + 1; i <= gt; {
		t := a[i][d]
		if t < v {
			m.exch(a, lt, i)
			lt++
			i++
		} else if t > v {
			m.exch(a, i, gt)
			gt--
		} else {
			i++
		}
	}
	m.sort(a, lo, lt-1, d)
	if v >= 0 {
		m.sort(a, lt, gt, d+1)
	}
	m.sort(a, gt+1, hi, d)

}

func (m *Quick3String) exch(a []string, i, j int) {
	a[i], a[j] = a[j], a[i]
}
