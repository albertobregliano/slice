package slice

// Unshift appends x in front of s and updates s.
func Unshift[V any](s *[]V, x V) {
	*s = append([]V{x}, *s...)
}

// PushFront is an alias of Unshift function.
// Unshift appends x in front of s and updates s.
func PushFront[V any](s *[]V, x V) {
	Unshift(s, x)
}
