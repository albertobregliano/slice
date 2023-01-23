package slice

// Filter returns only the elements of s that satisfy the keep function.
func Filter[V any](s []V, keep func(V) bool) []V {
	var out []V
	for _, x := range s {
		if keep(x) {
			out = append(out, x)
		}
	}
	return out
}
