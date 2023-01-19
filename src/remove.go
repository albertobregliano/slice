package slice

import "sort"

// Remove returns a copy of slice with its elements indexed in n removed.
func Remove(slice []any, n ...int) []any {
	s := make([]any, len(slice))
	copy(s, slice)
	filter := make(map[int]struct{}, len(n))
	for _, e := range n {
		filter[e] = struct{}{}
	}
	var nn []int
	for e := range filter {
		nn = append(nn, e)
	}
	sort.Ints(nn)
	for index, i := range nn {
		if index < len(s) {
			s = append(s[:i-index], s[i+1-index:]...)
		}
	}
	return s
}
