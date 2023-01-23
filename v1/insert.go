package slice

// Insert returns the elements of s and e inserted in the i position of s.
func Insert[V any](s []V, i int, e ...V) []V {
	if i < len(s) {
		s = append(s[:i], append(e, s[i:]...)...)
	}
	return s
}