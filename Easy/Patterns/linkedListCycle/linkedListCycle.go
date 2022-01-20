package linkedListCycle

import "fmt"

func Example() {
	fmt.Println("\nLinked List Cycle")
	head := ListNode{5, nil}
	next := ListNode{1, nil}
	head.Next = &next
	fmt.Println("list = [5 1]; pos = -1")
	result := hasCycle(&head)
	fmt.Println("result =", result)
	head.Next.Next = &head
	fmt.Println("\nlist = [5 1]; pos = 0")
	result = hasCycle(&head)
	fmt.Println("result =", result)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	temp := ListNode{}
	for head != nil {
		if head.Next == nil {
			return false
		}
		if head.Next == &temp {
			return true
		}
		next := head.Next
		head.Next = &temp
		head = next
	}
	return false
}
