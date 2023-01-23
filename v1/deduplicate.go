package slice

// Deduplicate returns the elements of s without duplicates.
func Deduplicate[V comparable](s []V) []V {
	var out = make([]V, 0, len(s))
	filter := make(map[V]bool, len(s)/2)
	for _, e := range s {
		if filter[e] {
			continue
		}
		filter[e] = true
		out = append(out, e)
	}
	return out
}
