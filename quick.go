package algo

import "math/rand"

func partition(items SortInterface, lo, hi int) int {
	i, j := lo, hi
	for i < j {
		for ; i <= hi && items.Less(i, lo); i++ {
		}
		for ; j > lo && items.Less(lo, j); j-- {
		}
		items.Swap(i, j)
	}
	items.Swap(lo, j)
	return j
}

func quicksort(items SortInterface, lo, hi int) {
	if lo >= hi {
		return
	}

	j := partition(items, lo, hi)
	quicksort(items, lo, j)
	quicksort(items, j+1, hi)
}

// QuickSort ...
func QuickSort(items SortInterface) {
	rand.Shuffle(items.Len(), func(i, j int) {
		items.Swap(i, j)
	})
	quicksort(items, 0, items.Len()-1)
}
