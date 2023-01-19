package slice

import (
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

// Remove removes all indexed elements in n from s.
func (s *Slice) Remove(n ...int) {
	s.Lock()
	defer s.Unlock()
	s.Elements = Remove(s.Elements, n...)
}

// Duplicate returns a slice with the same elements of s.
func Duplicate[V any](s []V) []V {
	return append(s[:0:0], s...)
}

// Shift returns the first element of a, removes it from a and updates a.
func Shift[V any](a *[]V) V {
	var x V
	x, *a = (*a)[0], (*a)[1:]
	return x
}

// Filter keeps in s.Elements only the elements that satisfy the keep function.
func (s *Slice) Filter(keep func(v any) bool) {
	s.Lock()
	defer s.Unlock()
	s.Elements = Filter(s.Elements, keep)
}

// Scramble modify the order of s.Elements randomly.
func (s *Slice) Scramble() {
	s.Lock()
	defer s.Unlock()
	s.Elements = Scramble(s.Elements)
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
// func PopFront[V any](a *[]V) V {
// 	return Shift(a)
// }

// Insert adds e elements to s starting from i position.
func (s *Slice) Insert(i int, e ...any) {
	s.Lock()
	defer s.Unlock()
	if i < len(s.Elements) {
		s.Elements = append(s.Elements[:i], append(e, s.Elements[i:]...)...)
	}
}

// Insert returns the elements of s and e inserted in the i position of s.
func Insert[V any](s []V, i int, e ...V) []V {
	if i < len(s) {
		s = append(s[:i], append(e, s[i:]...)...)
	}
	return s
}
