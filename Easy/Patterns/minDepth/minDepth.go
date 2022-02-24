package minDepth

import (
	"LeetCodePatterns/Easy/utils"
	"fmt"
)

func Example() {
	fmt.Println("\nMinimum Depth of Binary Tree")
	nodeArr := []int{3, 9, 20, -1, -1, 15, 7}
	node := utils.TreeNode{
		Val:   3,
		Left:  &utils.TreeNode{Val: 9},
		Right: &utils.TreeNode{Val: 20, Left: &utils.TreeNode{Val: 15}, Right: &utils.TreeNode{Val: 7}},
	}
	fmt.Println("node =", nodeArr)
	min := minDepth(&node)
	fmt.Println("min =", min)
}

func minDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	min := 100000
	ret(root, &min, 0)
	return min + 1
}

func ret(root *utils.TreeNode, min *int, i int) {
	if root.Left != nil {
		ret(root.Left, min, i+1)
	}
	if root.Right != nil {
		ret(root.Right, min, i+1)
	}
	if root.Right == nil && root.Left == nil {
		if *min > i {
			*min = i
		}
	}
}
