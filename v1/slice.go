package slice

/* Alias functions */

// Pop returns the last element of s, removes it from s and updates s.
func Pop[V any](s *[]V) V {
	var x V
	x, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return x
}

// PopFront is and alias of the Shift function.
// Shift returns the first value of s removes it from s and updates s.
func PopFront[V any](s *[]V) V {
	return Shift(s)
}

// PushFront is an alias of Unshift function.
// Unshift appends x in front of s and updates s.
func PushFront[V any](s *[]V, x V) {
	Unshift(s, x)
}

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
