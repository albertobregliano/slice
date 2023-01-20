package slice_test

import (
	"fmt"

	slice "github.com/albertobregliano/slice/src"
)

func ExampleFilter_int() {
	s := slice.Filter([]int{0, 1, 2, 3, 4}, func(x int) bool { return x < 3 })
	fmt.Println(s)
	// Output:
	// [0 1 2]
}

func ExampleFilter_str() {
	s := slice.Filter([]string{"zero", "one", "two", "three", "four"},
		func(x string) bool { return []rune(x)[0] == 't' })
	fmt.Println(s)
	// Output:
	// [two three]
}

func ExampleRemove_int() {
	s := slice.Remove([]int{0, 1, 2, 3}, 0, 2)
	fmt.Println(s)
	// Output:
	// [1 3]
}

func ExampleRemove_str() {
	s := slice.Remove([]string{"0", "1", "2", "3"}, 1, 3)
	fmt.Println(s)
	// Output:
	// [0 2]
}

func ExampleReverse_int() {
	s := slice.Reverse([]int{0, 1, 2, 3, 4})
	fmt.Println(s)
	// Output:
	// [4 3 2 1 0]
}

func ExampleReverse_str() {
	s := slice.Reverse([]string{"zero", "one", "two", "three", "four"})
	fmt.Println(s)
	// Output:
	// [four three two one zero]
}

func ExampleInsert_str() {
	s := slice.Insert([]string{"zero", "one", "two", "three", "four"}, 2, "two2", "three3")
	fmt.Println(s)
	// Output:
	// [zero one two2 three3 two three four]
}

func ExampleInsert_int() {
	s := slice.Insert([]int{0, 1, 2, 3, 4}, 2, 22, 33)
	fmt.Println(s)
	// Output:
	// [0 1 22 33 2 3 4]
}
