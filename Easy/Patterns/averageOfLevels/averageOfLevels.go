package averageOfLevels

import (
	"LeetCodePatterns/Easy/utils"
	"fmt"
)

func Example() {
	fmt.Println("\nAverage of Levels in Tree")
	nodeArr := []int{3, 9, 20, -1, -1, 15, 7}
	node := utils.TreeNode{
		Val:   3,
		Left:  &utils.TreeNode{Val: 9},
		Right: &utils.TreeNode{Val: 20, Left: &utils.TreeNode{Val: 15}, Right: &utils.TreeNode{Val: 7}},
	}
	fmt.Println("node =", nodeArr)
	arr := averageOfLevels(&node)
	fmt.Println("arr =", arr)
}

func averageOfLevels(root *utils.TreeNode) []float64 {
	arr := make([]float64, 0)
	s := make([]int, 0)
	i := 0
	ret(root, &arr, &s, i)
	for j := range arr {
		arr[j] /= float64(s[j])
	}
	return arr
}

func ret(root *utils.TreeNode, arr *[]float64, s *[]int, i int) {
	if len(*arr) == i {
		*arr = append(*arr, 0)
		*s = append(*s, 0)

	}
	(*arr)[i] += float64(root.Val)
	(*s)[i] += 1
	if root.Left != nil {
		ret(root.Left, arr, s, i+1)
	}
	if root.Right != nil {
		ret(root.Right, arr, s, i+1)
	}
}
