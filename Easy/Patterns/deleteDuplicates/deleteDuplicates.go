package deleteDuplicates

import (
	"LeetCodePatterns/Easy/utils"
	"fmt"
)

func Example() {
	fmt.Println("\nRemove Duplicates from Linked List")
	arr := []int{1, 1, 2, 2, 3, 3, 3, 4}
	head := utils.ArrayToList(arr, len(arr))
	fmt.Println("list =", arr)
	result := deleteDuplicates(head)
	fmt.Println("result = " + result.Show())
}

func deleteDuplicates(head *utils.ListNode) *utils.ListNode {
	node := head
	if node == nil {
		return node
	}
	next := utils.Walk(node, node.Val)
	node.Next = next
	if node.Next != nil {
		for i := next; i != nil; {
			i.Next = utils.Walk(i, i.Val)
			i = i.Next
		}
	}
	return node
}
