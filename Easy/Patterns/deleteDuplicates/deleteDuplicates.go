package deleteDuplicates

import (
	"LeetCodePatterns/Easy/Patterns/isPalindrome"
	"LeetCodePatterns/Easy/Patterns/removeElements"
	"fmt"
)

func Example() {
	fmt.Println("\nRemove Duplicates from Linked List")
	arr := []int{1, 1, 2, 2, 3, 3, 3, 4}
	head := isPalindrome.ArrayToList(arr, len(arr))
	fmt.Println("list =", arr)
	result := deleteDuplicates(head)
	fmt.Println("result = " + result.Show())
}

func deleteDuplicates(head *isPalindrome.ListNode) *isPalindrome.ListNode {
	node := head
	if node == nil {
		return node
	}
	next := removeElements.Walk(node, node.Val)
	node.Next = next
	if node.Next != nil {
		for i := next; i != nil; {
			i.Next = removeElements.Walk(i, i.Val)
			i = i.Next
		}
	}
	return node
}
