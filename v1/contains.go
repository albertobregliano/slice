package slice

// Contains returns whether x is present in s elements.
func Contains[V comparable](s []V, x V) bool {
	for _, e := range s {
		if e == x {
			return true
		}
	}
	return false
}
