package algs4

// InsertionSort ...
func InsertionSort(items SortInterface) {
	length := items.Len()
	for i := 1; i < length; i++ {
		for j := i; j > 0 && items.Less(j, j-1); j-- {
			items.Swap(j, j-1)
		}
	}
}
