package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	neg := make([]int, 0)
	pos := make([]int, 0)
	mor := make([]int, 0)
	les := make([]int, 0)
	equ := make([]int, 0)
	even := (target % 2) == 0
	sig := target >= 0
	hal := target / 2
	zer := -1

	for i, j := range nums {
		if j >= 0 {
			pos = append(pos, i)
		} else {
			neg = append(neg, i)
		}
	}

	for i, j := range pos {
		if nums[j] == 0 {
			zer = j
			pos[i] = pos[len(pos)-1]
			pos[len(pos)-1] = 0
			pos = pos[:len(pos)-1]
			break
		}
	}

	if zer != -1 {
		if sig {
			for _, j := range pos {
				if nums[j] == target {
					return []int{zer, j}
				}
			}
		} else {
			for _, j := range neg {
				if nums[j] == target {
					return []int{zer, j}
				}
			}
		}
	}

	if sig {
		for i := 0; i < len(pos); i++ {
			for nums[pos[i]] > target {
				mor = append(mor, pos[i])
				pos[i] = pos[len(pos)-1]
				pos[len(pos)-1] = 0
				pos = pos[:len(pos)-1]
				if i == len(pos) {
					break
				}
			}
		}
		if len(neg) > 0 && len(mor) > 0 {
			for _, j := range neg {
				for _, n := range mor {
					if (nums[n] + nums[j]) == target {
						return []int{n, j}
					}
				}
			}
		}
		if !even {
			for i := 0; i < len(pos); i++ {
				for nums[pos[i]] <= hal {
					les = append(les, pos[i])
					pos[i] = pos[len(pos)-1]
					pos[len(pos)-1] = 0
					pos = pos[:len(pos)-1]
					if i == len(pos) {
						break
					}
				}
			}
		} else {
			for i := 0; i < len(pos); i++ {
				for nums[pos[i]] == hal {
					equ = append(equ, pos[i])
					if len(equ) == 2 {
						return []int{equ[0], equ[1]}
					}
					pos[i] = pos[len(pos)-1]
					pos[len(pos)-1] = 0
					pos = pos[:len(pos)-1]
					if i == len(pos) {
						break
					}
				}
			}
			for i := 0; i < len(pos); i++ {
				for nums[pos[i]] < hal {
					les = append(les, pos[i])
					pos[i] = pos[len(pos)-1]
					pos[len(pos)-1] = 0
					pos = pos[:len(pos)-1]
					if i == len(pos) {
						break
					}
				}
			}
		}

		for j := len(pos) - 1; j >= 0; j-- {
			for _, n := range les {
				if (nums[n] + nums[pos[j]]) == target {
					return []int{n, pos[j]}
				}
			}
		}
	} else {
		for i := 0; i < len(neg); i++ {
			for -nums[neg[i]] > -target {
				mor = append(mor, neg[i])
				neg[i] = neg[len(neg)-1]
				neg[len(neg)-1] = 0
				neg = neg[:len(neg)-1]
				if i == len(neg) {
					break
				}
			}
		}
		if len(pos) > 0 && len(mor) > 0 {
			for _, j := range pos {
				for _, n := range mor {
					if (nums[n] + nums[j]) == target {
						return []int{n, j}
					}
				}
			}
		}

		if !even {
			for i := 0; i < len(neg); i++ {
				for nums[neg[i]] > hal {
					les = append(les, neg[i])
					neg[i] = neg[len(neg)-1]
					neg[len(neg)-1] = 0
					neg = neg[:len(neg)-1]
					if i == len(neg) {
						break
					}
				}
			}
		} else {
			for i := 0; i < len(neg); i++ {
				for nums[neg[i]] == hal {
					equ = append(equ, neg[i])
					if len(equ) == 2 {
						return []int{equ[0], equ[1]}
					}
					neg[i] = neg[len(neg)-1]
					neg[len(neg)-1] = 0
					neg = neg[:len(neg)-1]
					if i == len(neg) {
						break
					}
				}
			}
			for i := 0; i < len(neg); i++ {
				for nums[neg[i]] < hal {
					les = append(les, neg[i])
					neg[i] = neg[len(neg)-1]
					neg[len(neg)-1] = 0
					neg = neg[:len(neg)-1]
					if i == len(neg) {
						break
					}
				}
			}
		}

		for j := len(neg) - 1; j >= 0; j-- {
			for _, n := range les {
				if (nums[n] + nums[neg[j]]) == target {
					return []int{n, neg[j]}
				}
			}
		}
	}

	return []int{0, 0}
}
