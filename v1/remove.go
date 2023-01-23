package slice

import "sort"

// Remove returns all slice elements but those whose index is in indexes.
func Remove[V any](slice []V, indexes ...int) []V {
	s := make([]V, len(slice))
	copy(s, slice)
	for i, index := range toRemove(indexes) {
		if i < len(s) {
			s = append(s[:index-i], s[index+1-i:]...)
		}
	}
	return s
}

func toRemove(n []int) []int {
	filter := make(map[int]struct{}, len(n))
	for _, e := range n {
		filter[e] = struct{}{}
	}
	var nn = make([]int, 0, len(filter))
	for e := range filter {
		nn = append(nn, e)
	}
	sort.Ints(nn)
	return nn
}
