package maxProfit

import "fmt"

func Example() {
	fmt.Println("\nMax Profit")
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("nums =", nums)
	result := maxProfit(nums)
	fmt.Println("result =", result)
}

func maxProfit(prices []int) int {
	l := len(prices)
	m := make([]int, 1)
	n := make([]int, 0)
	m[0] = l - 1
	k := 0
	for i := l - 1; i > 0; i-- {
		if prices[m[k]] < prices[i-1] {
			m = append(m, i-1)
			k++
		} else {
			n = append(n, i-1)
		}
	}
	max := 0
	tmax := 0
	k = len(m) - 1
	for i := len(n) - 1; i >= 0; i-- {
		for n[i] > m[k] {
			k--
		}
		if k < 0 {
			break
		}
		tmax = prices[m[k]] - prices[n[i]]

		if max < tmax {
			max = tmax
		}
	}
	return max
}
