package algs4

import (
	"math/rand"
)

func partition3way(items SortInterface, lo, hi int) (int, int) {
	lt, i, gt := lo, lo+1, hi
	for i <= gt {
		if items.Less(i, lt) {
			items.Swap(i, lt)
			i++
			lt++
		} else if items.Less(lt, i) {
			items.Swap(i, gt)
			gt--
		} else {
			i++
		}

	}
	return lt, gt
}

func quicksort3way(items SortInterface, lo, hi int) {
	if lo >= hi {
		return
	}

	lt, gt := partition3way(items, lo, hi)
	quicksort3way(items, lo, lt-1)
	quicksort3way(items, gt+1, hi)
}

// QuickSort3way ...
func QuickSort3way(items SortInterface) {
	rand.Shuffle(items.Len(), func(i, j int) {
		items.Swap(i, j)
	})
	quicksort3way(items, 0, items.Len()-1)
}
