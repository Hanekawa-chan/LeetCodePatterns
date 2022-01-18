package isPalindrome

import (
	"fmt"
	"strconv"
)

func Example() {
	fmt.Println("\nPalindrome Linked List")
	arr := []int{5, 1, 1, 5}
	head := ArrayToList(arr, len(arr))
	fmt.Println("list =", arr)
	result := isPalindrome(head)
	fmt.Println("result =", result)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) Show() string {
	res := ""
	for i := n; i != nil; i = i.Next {
		res += strconv.Itoa(i.Val) + " "
	}
	return res
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

func isPalindrome(head *ListNode) bool {
	size := 1
	var temp *ListNode
	if head.Next != nil {
		for i := head; i != nil; i = i.Next {
			size++
			next := ListNode{i.Val, temp}
			temp = &next
		}
		t := head
		for i := 0; i < size/2; i++ {
			if temp.Val != t.Val {
				return false
			}
			temp = temp.Next
			t = t.Next
		}
	}
	return true
}
