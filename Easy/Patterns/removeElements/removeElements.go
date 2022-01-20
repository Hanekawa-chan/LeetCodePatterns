package removeElements

import (
	"LeetCodePatterns/Easy/Patterns/isPalindrome"
	"fmt"
)

func Example() {
	fmt.Println("\nRemove Elements of Linked List")
	arr := []int{7, 1, 2, 7, 1, 7}
	head := isPalindrome.ArrayToList(arr, len(arr))
	fmt.Println("list =", arr)
	result := removeElements(head, 7)
	fmt.Println("result = " + result.Show())
}

func removeElements(head *isPalindrome.ListNode, val int) *isPalindrome.ListNode {
	var node *isPalindrome.ListNode
	if head != nil {
		if head.Val != val {
			node = head
		} else {
			node = Walk(head, val)
			if node == nil {
				return node
			}
		}
		next := Walk(node, val)
		node.Next = next
		if node.Next != nil {
			for i := next; i != nil; {
				i.Next = Walk(i, val)
				i = i.Next
			}
		}
		return node
	} else {
		return nil
	}

}

func Walk(head *isPalindrome.ListNode, val int) *isPalindrome.ListNode {
	for i := head.Next; i != nil; i = i.Next {
		if i.Val == val {
			continue
		}
		return i
	}
	return nil
}
