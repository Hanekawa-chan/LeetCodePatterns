package maximumSubarray

import "fmt"

func Example() {
	fmt.Println("\nMaximum Subarray")
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("nums =", nums)
	result := maxSubArray(nums)
	fmt.Println("result =", result)
}

func maxSubArray(nums []int) int {
	t, val := hasPositive(nums)
	if !t {
		if val != -1 {
			return nums[val]
		}
		return handleNegative(nums)
	} else {
		if val == len(nums) {
			return sumArray(nums)
		}
		return handlePositive(nums)
	}
}

func sumArray(nums []int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

func hasPositive(nums []int) (bool, int) {
	sum := 0
	k := -1
	for i, j := range nums {
		if j >= 0 {
			k = i
			sum += 1
			if sum > 1 {
				return true, sum
			}
		}
	}
	if sum == 1 {
		return false, k
	}
	return false, -1
}

func handlePositive(nums []int) int {
	v := 0
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] >= 0 {
			break
		}
	}
	pos := make([]int, 0)
	neg := make([]int, 0)
	con := false
	j := -1
	for ; i < len(nums); i++ {
		if nums[i] >= 0 {
			if con {
				pos[j] += nums[i]
			} else {
				j += 1
				pos = append(pos, nums[i])
				con = true
			}
		} else {
			if con {
				con = false
				neg = append(neg, nums[i]+pos[j])
			} else {
				neg[j] += nums[i]
			}
		}
	}
	if len(neg) == len(pos) {
		l := len(neg) - 1
		neg[l] = pos[l]
	} else {
		neg = append(neg, pos[len(neg)])
	}
	max := 0
	posmax := 0
	nmax := 0
	i = 0
	for ; i < len(neg); i++ {
		if pos[i] > posmax {
			posmax = pos[i]
		}
		if pos[i]+nmax > max {
			max = nmax + pos[i]
		}
		if neg[i]+nmax >= 0 {
			nmax += neg[i]
		} else {
			if nmax >= max {
				nmax += pos[i]
				max = nmax
			}
			nmax = 0
		}
	}
	if nmax > max {
		max = nmax
	}
	if posmax > max {
		max = posmax
	}
	v = max
	return v
}

func handleNegative(nums []int) int {
	max := nums[0]
	for _, j := range nums {
		if j > max {
			max = j
		}
	}
	return max
}
