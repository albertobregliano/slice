package slice

// Odd returns the elements with odd index.
// Counting starts from 1 not 0.
func Odd[V any](s []V) []V {
	var out = make([]V, 0, 1+len(s)/2)
	for i, e := range s {
		if (i+1)%2 != 0 {
			out = append(out, e)
		}
	}
	return out
}
