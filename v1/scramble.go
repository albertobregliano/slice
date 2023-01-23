package slice

import (
	"math/rand"
	"time"
)

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
