package slice

// Even returns the elements with even index.
// Counting starts from 1 not 0.
func Even[V any](s []V) []V {
	var out = make([]V, 0, len(s)/2)
	for i, e := range s {
		if (i+1)%2 == 0 {
			out = append(out, e)
		}
	}
	return out
}
