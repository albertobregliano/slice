package slice

import (
	"log"
	"math/rand"
	"time"
)

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
