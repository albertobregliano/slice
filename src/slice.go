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

// Remove removes from s all elements whose indexes are in toBeRemoved.
func (s *Slice) Remove(toBeRemoved ...int) {
	s.Lock()
	defer s.Unlock()
	s.Elements = Remove(s.Elements, toBeRemoved...)
}

// Duplicate returns a slice with the same elements of s.
func Duplicate[V any](s []V) []V {
	return append(s[:0:0], s...)
}

// Shift returns the first element of s, removes it from s and updates s.
func Shift[V any](s *[]V) V {
	var x V
	x, *s = (*s)[0], (*s)[1:]
	return x
}

// Filter keeps in s.Elements only the elements that satisfy the keep function.
func (s *Slice) Filter(keep func(v any) bool) {
	s.Lock()
	defer s.Unlock()
	s.Elements = Filter(s.Elements, keep)
}

// Scramble modify the order of s.Elements randomly.
// A seed can be used for deterministic outputs.
func (s *Slice) Scramble(seed ...int64) {
	s.Lock()
	defer s.Unlock()
	s.Elements = Scramble(s.Elements, seed...)
}

// Unshift appends x in front of a and updates a.
func Unshift[V any](s *[]V, x V) {
	*s = append([]V{x}, *s...)
}

// PushFront is an alias of Unshift function.
// Unshift appends x in front of a and updates a.
func PushFront[V any](s *[]V, x V) {
	Unshift(s, x)
}

// Pop returns the last element of a removing it from aand updates a.
func Pop[V any](s *[]V) V {
	var x V
	x, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return x
}

// PopFront is and alias of the Shift function.
// Shift returns the first value of a removes it from a and updates a.
func PopFront[V any](s *[]V) V {
	return Shift(s)
}

// Insert adds e elements to s starting from i position.
func (s *Slice) Insert(i int, e ...any) {
	s.Lock()
	defer s.Unlock()
	if i < len(s.Elements) {
		s.Elements = append(s.Elements[:i], append(e, s.Elements[i:]...)...)
	}
}
