package binarySearch

import "fmt"

func Example() {
	fmt.Println("\nMax Profit")
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println("nums =", nums)
	result := search(nums, 0)
	fmt.Println("result =", result)
}

func search(nums []int, target int) int {
	if nums[0] == target {
		return 0
	} else if len(nums) == 1 {
		return -1
	}
	mid := len(nums) / 2
	num := mid
	if nums[mid] == target {
		return mid
	}
	arr := nums
	for arr[mid] != target {
		if len(arr) == 1 {
			return -1
		}
		if arr[mid] > target {
			arr = arr[0:mid]
			mid = len(arr) / 2
			num -= mid + 1
		} else {
			arr = arr[mid+1:]
			mid = len(arr) / 2
			num += mid + 1
			if len(arr) == 0 {
				return -1
			}
		}

	}
	if nums[num] == target {
		return num
	}
	if nums[num] > target {
		for nums[num] != target {
			num--
		}
	} else {
		for nums[num] != target {
			num++
		}
	}
	return num
}
