package main

import (
	"fmt"

	slice "github.com/albertobregliano/slice/src"
)

func main() {
	var g = slice.New([]any{0, 1, 2, 3, 4, 5, 6})
	fmt.Println(g.Elements...)
	s := slice.Remove(g.Elements, 3, 0, 2, 6, 1, 3, 3, 3, 4, 5, 100)
	fmt.Println(s)
	s = slice.Remove(g.Elements, 0)
	fmt.Println(s)
	fmt.Println(g.Elements...)
	g.Remove(1)
	g.Remove(0)
	slice.Insert(g, 1, 10)
	fmt.Println(g.Elements...)
}
