package slice

func Copy[V any](a []V) []V {
	return append(a[:0:0], a...)
}

func Delete[V any](a []V, i int) []V {
	return append(a[:i], a[i+1:]...)
}

// Filter returns a slice with only the elements of a that meet
// the requrirements of the keep function.
func Filter[V any](a []V, keep func(V) bool) []V {
	n := 0
	for _, x := range a {
		if keep(x) {
			a[n] = x
			n++
		}
	}
	return a[:n]
}

// Shift returns the first value of a, removes it from a and updates a.
func Shift[V any](a *[]V) V {
	var x V
	x, *a = (*a)[0], (*a)[1:]
	return x
}

// Unshift appends x in front of a and updates a.
func Unshift[V any](a *[]V, x V) {
	*a = append([]V{x}, *a...)
}

// PushFront is an alias of Unshift function.
// Unshift appends x in front of a and updates a.
func PushFront[V any](a *[]V, x V) {
	Unshift(a, x)
}

// Pop returns the last element of a removing it from aand updates a.
func Pop[V any](a *[]V) V {
	var x V
	x, *a = (*a)[len(*a)-1], (*a)[:len(*a)-1]
	return x
}

// PopFront is and alias of the Shift function.
// Shift returns the first value of a removes it from a and updates a.
func PopFront[V any](a *[]V) V {
	return Shift(a)
}

// Insert inserts x in the i position of a and updates a.
func Insert[V any](a *[]V, i int, x any) {
	switch v := x.(type) {
	case V:
		*a = append((*a)[:i], append([]V{v}, (*a)[i:]...)...)
	case []V:
		*a = append((*a)[:i], append(v, (*a)[i:]...)...)
	}
}
