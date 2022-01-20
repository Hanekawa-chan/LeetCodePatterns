package isPalindrome

import (
	"LeetCodePatterns/Easy/utils"
	"fmt"
)

func Example() {
	fmt.Println("\nPalindrome Linked List")
	arr := []int{5, 1, 1, 5}
	head := utils.ArrayToList(arr, len(arr))
	fmt.Println("list =", arr)
	result := isPalindrome(head)
	fmt.Println("result =", result)
}

func isPalindrome(head *utils.ListNode) bool {
	size := 1
	var temp *utils.ListNode
	if head.Next != nil {
		for i := head; i != nil; i = i.Next {
			size++
			next := utils.ListNode{Val: i.Val, Next: temp}
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
