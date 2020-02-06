package algs4

// LSDSort ...
func LSDSort(a []string, w int) {
	N := len(a)
	R := 256
	aux := make([]string, N)
	for d := w - 1; d >= 0; d-- {
		count := make([]int, R+1)
		// Compute frequency counts.
		for i := 0; i < N; i++ {
			count[a[i][d]+1]++
		}
		// Transform counts to indices.
		for r := 0; r < R; r++ {
			count[r+1] += count[r]
		}
		// Distribute.
		for i := 0; i < N; i++ {
			aux[count[a[i][d]]] = a[i]
			count[a[i][d]]++
		}
		// Copy Back
		for i := 0; i < N; i++ {
			a[i] = aux[i]
		}

	}
}
