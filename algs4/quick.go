package algs4

import (
	"math/rand"
)

func partition(items SortInterface, lo, hi int) int {
	i, j := lo, hi+1
	for {
		for {
			i++
			if !(items.Less(i, lo) && i < hi) {
				break
			}
		}
		for {
			j--
			if !(items.Less(lo, j) && j > lo) {
				break
			}
		}
		if i >= j {
			break
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
	quicksort(items, lo, j-1)
	quicksort(items, j+1, hi)
}

// QuickSort ...
func QuickSort(items SortInterface) {
	rand.Shuffle(items.Len(), func(i, j int) {
		items.Swap(i, j)
	})
	quicksort(items, 0, items.Len()-1)
}
