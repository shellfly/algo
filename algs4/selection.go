package algs4

// SelectionSort ...
func SelectionSort(items SortInterface) {
	length := items.Len()
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if items.Less(j, i) {
				items.Swap(i, j)
			}
		}
	}
}
