package slice

// Unshift appends x in front of s and updates s.
func Unshift[V any](s *[]V, x V) {
	*s = append([]V{x}, *s...)
}
