// Package slice uses generics to apply basic operations on slices.
package slice

import (
	"math/rand"
	"sort"
	"time"
)

// Duplicate returns a slice with the same elements of s.
func Duplicate[V any](s []V) []V {
	return append(s[:0:0], s...)
}

// Filter returns only the elements of s that satisfy the keep function.
func Filter[V any](s []V, keep func(V) bool) []V {
	var out []V
	for _, x := range s {
		if keep(x) {
			out = append(out, x)
		}
	}
	return out
}

// Insert returns the elements of s and e inserted in the i position of s.
func Insert[V any](s []V, i int, e ...V) []V {
	if i < len(s) {
		s = append(s[:i], append(e, s[i:]...)...)
	}
	return s
}

// Remove returns all slice elements but those whose index is in indexes.
func Remove[V any](slice []V, indexes ...int) []V {
	s := make([]V, len(slice))
	copy(s, slice)
	for i, index := range toRemove(indexes) {
		if i < len(s) {
			s = append(s[:index-i], s[index+1-i:]...)
		}
	}
	return s
}

func toRemove(n []int) []int {
	filter := make(map[int]struct{}, len(n))
	for _, e := range n {
		filter[e] = struct{}{}
	}
	var nn []int
	for e := range filter {
		nn = append(nn, e)
	}
	sort.Ints(nn)
	return nn
}

// Reverse returns the elemts of s from the last one to the first.
func Reverse[V any](s []V) []V {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return s
}

// Scramble returns the elements of slice in random order.
// A seed can be specified for deterministic outputs.
func Scramble[V any](slice []V, seed ...int64) []V {
	s := make([]V, len(slice))
	copy(s, slice)
	var randSeed = time.Now().UnixNano()
	if len(seed) > 0 {
		randSeed = seed[0]
	}
	var randomizer = rand.New(rand.NewSource(randSeed))
	for i := 0; i < 10*len(s); i++ {
		n := randomizer.Intn(len(s))
		m := randomizer.Intn(len(s))
		s[m], s[n] = s[n], s[m]
	}
	return s
}

// Shift returns the first element of s, removes it from s and updates s.
func Shift[V any](s *[]V) V {
	var x V
	x, *s = (*s)[0], (*s)[1:]
	return x
}

// Unshift appends x in front of s and updates s.
func Unshift[V any](s *[]V, x V) {
	*s = append([]V{x}, *s...)
}

// PushFront is an alias of Unshift function.
// Unshift appends x in front of s and updates s.
func PushFront[V any](s *[]V, x V) {
	Unshift(s, x)
}

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
