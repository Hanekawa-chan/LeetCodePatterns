package missingNumber

import "fmt"

func Example() {
	fmt.Println("\nMissing Number")
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println("nums =", nums)
	result := missingNumber(nums)
	fmt.Println("result =", result)
}

func missingNumber(nums []int) int {
	sum := 0
	num := 0
	for i, j := range nums {
		sum += j
		num += i + 1
	}
	return num - sum
}
