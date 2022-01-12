package containsDuplicate

import "fmt"

func Example() {
	fmt.Println("\nContains Duplicate")
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println("nums =", nums)
	result := containsDuplicate(nums)
	fmt.Println("result =", result)
}

func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for _, j := range nums {
		if _, ok := seen[j]; ok {
			return true
		}
		seen[j] = struct{}{}
	}
	return false
}
