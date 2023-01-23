package slice

// Shift returns the first element of s, removes it from s and updates s.
func Shift[V any](s *[]V) V {
	var x V
	x, *s = (*s)[0], (*s)[1:]
	return x
}

// PopFront is and alias of the Shift function.
// Shift returns the first value of s removes it from s and updates s.
func PopFront[V any](s *[]V) V {
	return Shift(s)
}
