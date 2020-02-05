package algs4

func sink(items SortInterface, k, N int) {
	for 2*k <= N {
		j := 2 * k
		if j < N && less(items, j, j+1) {
			j++
		}
		if !less(items, k, j) {
			break
		}
		exch(items, k, j)
		k = j
	}
}

func less(items SortInterface, i, j int) bool {
	return items.Less(i-1, j-1)
}

func exch(items SortInterface, i, j int) {
	items.Swap(i-1, j-1)
}

// HeapSort ...
func HeapSort(items SortInterface) {
	N := items.Len()
	for i := N / 2; i >= 1; i-- {
		sink(items, i, N)
	}
	for N > 1 {
		exch(items, 1, N)
		N--
		sink(items, 1, N)
	}
}
