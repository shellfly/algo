package algo

// SortInterface is borrow from go lib sort.Interface
type SortInterface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// StringSlice implements SortInterface
type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// IsSorted reports if the given items is sorted
func IsSorted(items SortInterface) bool {
	for i := 1; i < items.Len(); i++ {
		if items.Less(i, i-1) {
			return false
		}
	}
	return true
}
