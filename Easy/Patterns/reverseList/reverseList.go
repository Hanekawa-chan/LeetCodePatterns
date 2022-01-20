package reverseList

import (
	"LeetCodePatterns/Easy/utils"
	"fmt"
)

func Example() {
	fmt.Println("\nRemove Duplicates from Linked List")
	arr := []int{1, 1, 2, 2, 3, 3, 3, 4}
	head := utils.ArrayToList(arr, len(arr))
	fmt.Println("list =", arr)
	result := reverseList(head)
	fmt.Println("result = " + result.Show())
}

func reverseList(head *utils.ListNode) *utils.ListNode {
	node := head
	return node
}
