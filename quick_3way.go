package algo

import "math/rand"

func partition3way(items SortInterface, lo, hi int) (int, int) {
	i, lt, gt := lo+1, lo+1, hi
	for i <= gt {
		if items.Less(i, lo) {
			items.Swap(i, lt)
			i++
			lt++
		} else if items.Less(lo, i) {
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
	quicksort(items, lo, lt-1)
	quicksort(items, gt+1, hi)
}

// QuickSort3way ...
func QuickSort3way(items SortInterface) {
	rand.Shuffle(items.Len(), func(i, j int) {
		items.Swap(i, j)
	})
	quicksort3way(items, 0, items.Len()-1)
}
