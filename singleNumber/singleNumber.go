package singleNumber

import "fmt"

func Example() {
	fmt.Println("\nSingle Number")
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println("nums =", nums)
	result := singleNumber(nums)
	fmt.Println("result =", result)
}

func singleNumber(nums []int) int {
	mask := 0
	for _, j := range nums {
		mask ^= j
	}
	return mask
}
