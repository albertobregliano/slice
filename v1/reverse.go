package slice

// Reverse returns the elemts of s from the last one to the first.
func Reverse[V any](s []V) []V {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return s
}
