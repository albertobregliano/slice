package slice_test

import (
	"github.com/albertobregliano/slice"
	"testing"
)

func TestCopy(t *testing.T) {
	a := []string{"a", "b"}
	b := slice.Copy(a)
	if b[0] != "a" || b[1] != "b" {
		t.Fatal("error")
	}
}
