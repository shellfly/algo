package algs4

// MSD ...
type MSD struct {
	CUTOFF int // cutoff to insertion sort
	R      int // extended ASCII alphabet size
	aux    []string
}

// NewMSD ...
func NewMSD(a []string) *MSD {
	m := &MSD{CUTOFF: 15, R: 256, aux: make([]string, len(a))}
	m.sort(a, 0, len(a)-1, 0)
	return m
}

func (m *MSD) insertion(a []string, lo, hi, d int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && a[j][d] < a[j-1][d]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}
func (m *MSD) sort(a []string, lo, hi, d int) {
	if hi <= lo+m.CUTOFF {
		m.insertion(a, lo, hi, d)
		return
	}

	count := make([]int, m.R+2)
	for i := lo; i <= hi; i++ {
		count[a[i][d]+2]++
	}
	for r := 0; r < m.R+1; r++ {
		count[r+1] += count[r]
	}
	for i := lo; i <= hi; i++ {
		m.aux[count[a[i][d]+1]] = a[i]
		count[a[i][d]+1]++
	}
	for i := lo; i <= hi; i++ {
		a[i] = m.aux[i-lo]
	}

	for r := 0; r < m.R; r++ {
		m.sort(a, lo+count[r], lo+count[r+1]-1, d+1)
	}
}
