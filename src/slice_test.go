package slice_test

import (
	"fmt"

	slice "github.com/albertobregliano/slice/src"
)

func ExampleDuplicate_int() {
	intSlice := []int{0, 1, 2, 3, 4}
	s := slice.Duplicate(intSlice)
	fmt.Println(s)
	// Output:
	// [0 1 2 3 4]
}

func ExampleDuplicate_string() {
	strSlice := []string{"zero", "one", "two", "three", "four"}
	s := slice.Duplicate(strSlice)
	fmt.Println(s)
	// Output:
	// [zero one two three four]
}

func ExampleFilter_int() {
	intSlice := []int{0, 1, 2, 3, 4}
	s := slice.Filter(intSlice, func(x int) bool { return x < 3 })
	fmt.Println(s)
	// Output:
	// [0 1 2]
}

func ExampleFilter_string() {
	strSlice := []string{"zero", "one", "two", "three", "four"}
	s := slice.Filter(strSlice,
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

func ExampleRemove_string() {
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

func ExampleReverse_string() {
	strSlice := []string{"zero", "one", "two", "three", "four"}
	s := slice.Reverse(strSlice)
	fmt.Println(s)
	// Output:
	// [four three two one zero]
}

func ExampleInsert_string() {
	strSlice := []string{"zero", "one", "two", "three", "four"}
	s := slice.Insert(strSlice, 2, "two2", "three3")
	fmt.Println(s)
	// Output:
	// [zero one two2 three3 two three four]
}

func ExampleInsert_int() {
	intSlice := []int{0, 1, 2, 3, 4}
	s := slice.Insert(intSlice, 2, 22, 33)
	fmt.Println(s)
	// Output:
	// [0 1 22 33 2 3 4]
}

func ExampleScramble_int() {
	s := slice.Scramble([]int{0, 1, 2, 3, 4}, 1000)
	fmt.Println(s)
	// Output:
	// [1 3 0 4 2]
}

func ExampleScramble_string() {
	strSlice := []string{"zero", "one", "two", "three", "four"}
	s := slice.Scramble(strSlice, 1000)
	fmt.Println(s)
	// Output:
	// [one three zero four two]
}

func ExampleShift_int() {
	intSlice := []int{0, 1, 2, 3, 4}
	x := slice.Shift(&intSlice)
	fmt.Println(x)
	fmt.Println(intSlice)
	// Output:
	// 0
	// [1 2 3 4]
}

func ExampleUnshift_int() {
	intSlice := []int{0, 1, 2, 3, 4}
	slice.Unshift(&intSlice, -1)
	fmt.Println(intSlice)
	// Output:
	// [-1 0 1 2 3 4]
}

func ExamplePushFront_int() {
	intSlice := []int{0, 1, 2, 3, 4}
	slice.PushFront(&intSlice, -1)
	fmt.Println(intSlice)
	// Output:
	// [-1 0 1 2 3 4]
}

func ExamplePop_int() {
	intSlice := []int{0, 1, 2, 3, 4}
	x := slice.Pop(&intSlice)
	fmt.Println(x)
	fmt.Println(intSlice)
	// Output:
	// 4
	// [0 1 2 3]
}

func ExamplePopFront_string() {
	strSlice := []string{"zero", "one", "two", "three", "four"}
	x := slice.PopFront(&strSlice)
	fmt.Println(x)
	fmt.Println(strSlice)
	// Output:
	// zero
	// [one two three four]
}
