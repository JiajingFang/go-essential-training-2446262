// Calculate maximal value in a slice
package main

import (
	"fmt"
)

func main() {
	// nums := []int{16, 8, 42, 4, 23, 15}
	// // fmt.Println(nums)

	// max := nums[0] // Initialize max with first value
	// // [1:] skips the first value
	// for _, value := range nums {
	// 	if value > max {
	// 		max = value
	// 	}
	// }

	// fmt.Println(max)
	nums := []int{2, 5, 3, 7, 9, 19}

	fmt.Println(nums)

	max := nums[0]

	for _, value := range nums {
		if value > max {
			max = value
		}
	}
	fmt.Println("the max number is", max)
}
