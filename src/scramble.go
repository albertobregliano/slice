package slice

import (
	"math/rand"
	"time"
)

// Scramble returns the elements of slice in random order.
func Scramble[V any](slice []V) []V {
	s := make([]V, len(slice))
	copy(s, slice)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10*len(s); i++ {
		n := r.Intn(len(s))
		m := r.Intn(len(s))
		s[m], s[n] = s[n], s[m]
	}
	return s
}
