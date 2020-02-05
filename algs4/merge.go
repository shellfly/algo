// this module only work for string slice

package algs4

func merge(items StringSlice, lo, mid, hi int) {
	aux := make(StringSlice, len(items))
	copy(aux, items)

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			items[k] = aux[j]
			j++
		} else if j > hi {
			items[k] = aux[i]
			i++
		} else if aux.Less(i, j) {
			items[k] = aux[i]
			i++
		} else {
			items[k] = aux[j]
			j++
		}
	}
}

func mergesort(items StringSlice, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := (lo + hi) / 2
	mergesort(items, lo, mid)
	mergesort(items, mid+1, hi)
	merge(items, lo, mid, hi)
}

// MergeSort ...
func MergeSort(items StringSlice) {
	mergesort(items, 0, len(items)-1)
}
