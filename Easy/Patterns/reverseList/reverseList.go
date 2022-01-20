package reverseList

import (
	"LeetCodePatterns/Easy/utils"
	"fmt"
)

func Example() {
	fmt.Println("\nReverse Linked List")
	arr := []int{1, 7, 5, 2}
	head := utils.ArrayToList(arr, len(arr))
	fmt.Println("list =", arr)
	result := reverseList(head)
	fmt.Println("result = " + result.Show())
}

func reverseList(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}
	var temp *utils.ListNode
	for i := head; i != nil; i = i.Next {
		next := utils.ListNode{Val: i.Val, Next: temp}
		temp = &next
	}
	return temp
}
