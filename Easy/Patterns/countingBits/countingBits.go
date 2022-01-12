package countingBits

import (
	"fmt"
)

func Example() {
	fmt.Println("\nCounting Bits")
	num := 5
	fmt.Println("num =", num)
	result := countBits(num)
	fmt.Println("result =", result)
}

//Given an integer n, return an array ans of length n + 1
//such that for each i (0 <= i <= n),
//ans[i] is the number of 1's in the binary representation of i.
func countBits(n int) []int {
	arr := make([]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = countBit(i)
	}
	return arr
}

//medium dumb way to do this
func countBit(n int) int {
	t := n
	res := 0
	for i := 0; t >= 1; i++ {
		res += t & 1 //* int(math.Pow(10, float64(i)))
		t >>= 1
	}
	return res
}
