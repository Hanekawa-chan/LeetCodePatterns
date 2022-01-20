package removeElements

import (
	"LeetCodePatterns/Easy/utils"
	"fmt"
)

func Example() {
	fmt.Println("\nRemove Elements of Linked List")
	arr := []int{7, 1, 2, 7, 1, 7}
	head := utils.ArrayToList(arr, len(arr))
	fmt.Println("list =", arr)
	result := removeElements(head, 7)
	fmt.Println("result = " + result.Show())
}

func removeElements(head *utils.ListNode, val int) *utils.ListNode {
	var node *utils.ListNode
	if head != nil {
		if head.Val != val {
			node = head
		} else {
			node = utils.Walk(head, val)
			if node == nil {
				return node
			}
		}
		next := utils.Walk(node, val)
		node.Next = next
		if node.Next != nil {
			for i := next; i != nil; {
				i.Next = utils.Walk(i, val)
				i = i.Next
			}
		}
		return node
	} else {
		return nil
	}

}
