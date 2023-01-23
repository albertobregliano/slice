package slice

import (
	"log"
	"math/rand"
	"time"
)

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

// Contains returns whether x is present in s elements.
func Contains[V comparable](s []V, x V) bool {
	for _, e := range s {
		if e == x {
			return true
		}
	}
	return false
}

// Deduplicate returns the elements of s without duplicates.
func Deduplicate[V comparable](s []V) []V {
	var out = make([]V, 0, len(s))
	filter := make(map[V]bool, len(s)/2)
	for _, e := range s {
		if filter[e] {
			continue
		}
		filter[e] = true
		out = append(out, e)
	}
	return out
}

// Apply applies the f function to each element of s.
func Apply[V any](s []V, f func(x V) V) []V {
	var out = make([]V, 0, len(s))
	for _, e := range s {
		out = append(out, f(e))
	}
	return out
}

// Odd returns the elements with odd index.
// Counting starts from 1 not 0.
func Odd[V any](s []V) []V {
	var out = make([]V, 0, len(s)/2)
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

// Percentage returns a percent number of s elements chosen randomly.
// A seed can be used for deterministic outputs.
func Percentage[V any](s []V, percent int, seed ...int64) []V {
	var randSeed = time.Now().UnixNano()
	if len(seed) > 0 {
		randSeed = seed[0]
	}
	var randomizer = rand.New(rand.NewSource(randSeed))
	x := len(s) * percent / 100
	log.Println(x)
	var out = make([]V, 0, x)
	filter := make(map[int]bool, x)
	for len(out) != x {
		r := randomizer.Intn(len(s))
		if !filter[r] {
			out = append(out, s[r])
		}
		filter[r] = true
	}
	return out
}
