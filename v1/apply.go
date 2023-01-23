package slice

// Apply applies the f function to each element of s.
func Apply[V any](s []V, f func(x V) V) []V {
	var out = make([]V, 0, len(s))
	for _, e := range s {
		out = append(out, f(e))
	}
	return out
}
