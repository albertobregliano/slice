package slice

// Duplicate returns a slice with the same elements of s.
func Duplicate[V any](s []V) []V {
	return append(s[:0:0], s...)
}
