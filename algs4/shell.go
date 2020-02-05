package algs4

// ShellSort ...
func ShellSort(items SortInterface) {
	N := items.Len()
	h := 1
	for ; h < N/3; h = 3*h + 1 {
	}
	for ; h >= 1; h = h / 3 {
		for i := h; i < N; i++ {
			for j := i; j >= h && items.Less(j, j-h); j -= h {
				items.Swap(j, j-h)
			}
		}
	}
}
