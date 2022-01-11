package findDisappearedNumbers

import "fmt"

func Example() {
	fmt.Println("\nFind Disappeared Numbers")
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println("nums =", nums)
	result := findDisappearedNumbers(nums)
	fmt.Println("result =", result)
}

func findDisappearedNumbers(nums []int) []int {
	res := make(map[int]bool)
	r := make([]int, 0)
	for _, j := range nums {
		res[j] = true
	}
	for i := range nums {
		if !res[i+1] {
			r = append(r, i+1)
		}
	}
	return r
}
