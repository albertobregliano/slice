package slice

// Filter returns only the elements that satisfy the keep function.
func Filter[V any](a []V, keep func(V) bool) []V {
	var out []V
	for _, x := range a {
		if keep(x) {
			out = append(out, x)
		}
	}
	return out
}
