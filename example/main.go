package main

import (
	"fmt"

	slice "github.com/albertobregliano/slice/src"
)

func main() {
	var g = slice.New([]any{0, 1, 2, 3, 4, 5, 8, 9, 10, 12})
	fmt.Println(g.Elements...)
	s := slice.Remove(g.Elements, 3, 0, 2, 6, 1, 3, 3, 3, 4, 5, 100)
	fmt.Println(s)
	s = slice.Remove(g.Elements, 0)
	fmt.Println(s)
	fmt.Println(g.Elements...)
	g.Remove(1)
	g.Remove(0)
	g.Insert(1, 10)
	fmt.Println(g.Elements...)
	gg := slice.Duplicate(g.Elements)
	fmt.Println(gg...)
	fmt.Println("tee")
	t := []int64{2, 3}
	tee := slice.Shift(&t)
	fmt.Println(tee)
	fmt.Println(t)
	slice.Unshift(&t, 5)
	slice.Unshift(&t, 5)
	fmt.Println(t)
	tt := slice.Duplicate(t)
	for _, ttt := range tt {
		fmt.Println(ttt, len(t))
	}
	scramble := slice.Scramble([]rune("ammaccabanane"))
	fmt.Println(string(scramble), "\n", g.Elements)
	//scramble = slice.Scramble(g.Elements)
	//fmt.Println(scramble...)
	ss := slice.Filter(g.Elements, func(x any) bool { return x.(int)%2 == 0 })
	fmt.Println(ss...)
	fmt.Println("g")
	g.Filter(func(x any) bool { return x.(int)%2 == 0 })
	fmt.Println(g.Elements...)
}
