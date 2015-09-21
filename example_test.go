package permute_test

import (
	"fmt"
	"sort"

	"github.com/lvdlvd/go-sort-permute"
)

func ExampleNext() {
	x := []int{1, 2, 3} // expect 6 permutations
	fmt.Println(0, x)

	for i := 1; permute.Next(sort.IntSlice(x)); i++ {
		fmt.Println(i, x)
	}
	// Output:
	// 0 [1 2 3]
	// 1 [1 3 2]
	// 2 [2 1 3]
	// 3 [2 3 1]
	// 4 [3 1 2]
	// 5 [3 2 1]
}

func ExampleNext_withduplicate() {
	x := []int{0, 1, 1, 3} // expect 12 permutations
	fmt.Println(0, x)

	for i := 1; permute.Next(sort.IntSlice(x)); i++ {
		fmt.Println(i, x)
	}
	// Output:
	// 0 [0 1 1 3]
	// 1 [0 1 3 1]
	// 2 [0 3 1 1]
	// 3 [1 0 1 3]
	// 4 [1 0 3 1]
	// 5 [1 1 0 3]
	// 6 [1 1 3 0]
	// 7 [1 3 0 1]
	// 8 [1 3 1 0]
	// 9 [3 0 1 1]
	// 10 [3 1 0 1]
	// 11 [3 1 1 0]
}
