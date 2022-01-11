package climbingStairs

import "fmt"

func Example() {
	fmt.Println("\nClimbing stairs")
	num := 3
	fmt.Println("num =", num)
	result := climbStairs(num)
	fmt.Println("result =", result)
}

func climbStairs(n int) int {
	k := 1
	j := 2
	t := 0
	if n == 1 {
		return k
	} else if n == 2 {
		return j
	} else {
		for i := 3; i <= n; i++ {
			t = j
			j = k + j
			k = t
		}
	}
	return j
}
