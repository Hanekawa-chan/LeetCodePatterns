package middleNode

import (
	"LeetCodePatterns/Easy/utils"
	"fmt"
)

func Example() {
	fmt.Println("\nMiddle of the Linked List")
	arr := []int{5, 1, 1, 5}
	head := utils.ArrayToList(arr, len(arr))
	fmt.Println("list = [5 2 1]")
	result := middleNode(head)
	fmt.Println("result = " + result.Show())
}

func middleNode(head *utils.ListNode) *utils.ListNode {
	c := 0
	for i := head; i != nil; i = i.Next {
		c++
	}
	t := head
	for i := 0; i < c/2+c%2%1; i++ {
		t = t.Next
	}
	return t
}
