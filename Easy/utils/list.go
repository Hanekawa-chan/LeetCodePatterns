package utils

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insert(root *ListNode, item int) *ListNode {
	temp := ListNode{item, root}
	root = &temp
	return root
}

func ArrayToList(arr []int, n int) *ListNode {
	var root *ListNode
	for i := n - 1; i >= 0; i-- {
		root = insert(root, arr[i])
	}
	return root
}

func (n *ListNode) Show() string {
	res := ""
	for i := n; i != nil; i = i.Next {
		res += strconv.Itoa(i.Val) + " "
	}
	return res
}

func Walk(head *ListNode, val int) *ListNode {
	for i := head.Next; i != nil; i = i.Next {
		if i.Val == val {
			continue
		}
		return i
	}
	return nil
}
