package rangeSumQuery

import "fmt"

func Example() {
	fmt.Println("\nRange Sum Query")
	nums := []int{-2, 0, 3, -5, 2, -1}
	fmt.Println("nums =", nums)
	numArray := Constructor(nums)
	result := numArray.SumRange(0, 2)
	fmt.Println("result =", result)
}

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += this.nums[i]
	}
	return sum
}
