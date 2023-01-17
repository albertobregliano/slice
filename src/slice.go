package slice

import (
	"sort"
	"sync"
)

type Slice struct {
	Elements []any
	sync.RWMutex
}

// New returns a Slice instance.
func New(elements []any) *Slice {
	return &Slice{
		Elements: elements,
	}
}

// Remove returns a copy of slice with its elements indexed in n removed.
func Remove(slice []any, n ...int) []any {
	s := make([]any, len(slice))
	copy(s, slice)
	filter := make(map[int]struct{}, len(n))
	for _, e := range n {
		filter[e] = struct{}{}
	}
	var nn []int
	for e := range filter {
		nn = append(nn, e)
	}
	sort.Ints(nn)
	for index, i := range nn {
		if len(slice) > index {
			s = append(s[:i-index], s[i+1-index:]...)
		}
	}
	return s
}

// Remove removes the elements in n from s and updates s.
func (s *Slice) Remove(n ...int) {
	sort.Ints(n)
	s.Lock()
	defer s.Unlock()
	for _, i := range n {
		if i <= len(s.Elements)-1 {
			s.Elements = append(s.Elements[:i], s.Elements[i+1:]...)
		}
	}
}

func Copy[V any](a []V) []V {
	return append(a[:0:0], a...)
}

// Filter returns a slice with only the elements of a that meet
// the requrirements of the keep function.
func Filter(slice *Slice, keep func(v any) bool) []any {
	n := 0
	for _, x := range slice.Elements {
		if keep(x) {
			slice.Elements[n] = x
			n++
		}
	}
	return slice.Elements[:n]
}

// Filter returns a slice with only the elements of a that meet
// the requrirements of the keep function.
// func Filter[V any](a []V, keep func(V) bool) []V {
// 	n := 0
// 	for _, x := range a {
// 		if keep(x) {
// 			a[n] = x
// 			n++
// 		}
// 	}
// 	return a[:n]
// }

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
func Insert(g *Slice, i int, e ...any) {
	g.Lock()
	defer g.Unlock()
	if i <= len(g.Elements)-1 {
		g.Elements = append(g.Elements[:i], append(e, g.Elements[i:]...)...)
	}
}

// Insert inserts x in the i position of a and updates a.
// func Insert[V any](a *[]V, i int, x any) {
// 	switch v := x.(type) {
// 	case V:
// 		*a = append((*a)[:i], append([]V{v}, (*a)[i:]...)...)
// 	case []V:
// 		*a = append((*a)[:i], append(v, (*a)[i:]...)...)
// 	}
// }
