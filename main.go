package main

import (
	"fmt"
	"iter"
	"slices"

	"github.com/danielmesquitta/iterators/xiter"
)

func SquareFirstTwoEvens() {

	nums := []int{1, 2, 3, 4, 5, 6}
	seq := slices.Values(nums)
	seq = xiter.Filter( // Filter out odd numbers
		func(v int) bool { return v%2 == 0 },
		seq,
	)
	seq = xiter.Limit(seq, 2) // Limit to 2 values
	seq = xiter.Map(          // Map the values to their squares
		func(v int) int { return v * v },
		seq,
	)
	nums = slices.Collect(seq)
	fmt.Println(nums) // [4 16]

}

// Countdown takes a value and returns a function
// that takes `func(int) bool` as a parameter,
// named yield by convention
func Countdown(v int) iter.Seq[int] {
	return func(yield func(int) bool) {
		// Loop from the value down to 0
		for i := v; i >= 0; i-- {
			// Stop if yield returns false
			// this is used to break the loop
			if !yield(i) {
				return
			}
		}
	}
}

func UseCountdown() {
	for v := range Countdown(5) {
		fmt.Printf("%v ", v) // 5 4 3 2 1 0
	}
	fmt.Println()
}

func UseBackward() {
	s := []int{10, 20, 30}
	for i, v := range slices.Backward(s) {
		fmt.Printf("i:%v,v:%v; ", i, v)
	}
	fmt.Println() // i:2,v:30; i:1,v:20; i:0,v:10;
}

func main() {
	UseCountdown()
	SquareFirstTwoEvens()
}
