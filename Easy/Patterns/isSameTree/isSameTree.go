package isSameTree

import (
	"LeetCodePatterns/Easy/utils"
	"fmt"
)

func Example() {
	fmt.Println("\nSame Tree")
	nodeArr1 := []int{3, 9, 20, -1, -1, 15, 7}
	nodeArr2 := []int{3, 9, 20, -1, -1, 12, 7}
	node1 := utils.TreeNode{
		Val:   3,
		Left:  &utils.TreeNode{Val: 9},
		Right: &utils.TreeNode{Val: 20, Left: &utils.TreeNode{Val: 15}, Right: &utils.TreeNode{Val: 7}},
	}
	node2 := utils.TreeNode{
		Val:   3,
		Left:  &utils.TreeNode{Val: 9},
		Right: &utils.TreeNode{Val: 20, Left: &utils.TreeNode{Val: 12}, Right: &utils.TreeNode{Val: 7}},
	}
	fmt.Println("first node =", nodeArr1)
	fmt.Println("second node =", nodeArr2)
	isSame := isSameTree(&node1, &node2)
	fmt.Println("result =", isSame)
}

func isSameTree(p *utils.TreeNode, q *utils.TreeNode) bool {
	if p == nil {
		if q == nil {
			return true
		} else {
			return false
		}
	}
	if q == nil {
		if p == nil {
			return true
		} else {
			return false
		}
	}
	equal := true
	ret(p, q, &equal)
	return equal
}

func ret(p *utils.TreeNode, q *utils.TreeNode, i *bool) {
	if !*i {
		return
	}
	if p.Val != q.Val {
		*i = false
		return
	}
	if (p.Left != nil && q.Left != nil) || (p.Left == nil && q.Left == nil) {
		if p.Left != nil {
			ret(p.Left, q.Left, i)
		}
	} else {
		*i = false
		return
	}
	if (p.Right != nil && q.Right != nil) || (p.Right == nil && q.Right == nil) {
		if p.Right != nil {
			ret(p.Right, q.Right, i)
		}
	} else {
		*i = false
		return
	}
}
