package slice

// Pop returns the last element of s, removes it from s and updates s.
func Pop[V any](s *[]V) V {
	var x V
	x, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return x
}
