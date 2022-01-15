package middleNode

import (
	"fmt"
	"strconv"
)

func Example() {
	fmt.Println("\nLinked List Cycle")
	head := ListNode{5, nil}
	next := ListNode{1, nil}
	next1 := ListNode{2, &next}
	head.Next = &next1
	fmt.Println("list = [5 2 1]")
	result := middleNode(&head)
	fmt.Println("result = " + result.show())
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) show() string {
	res := ""
	for i := n; i != nil; i = i.Next {
		res += strconv.Itoa(i.Val) + " "
	}
	return res
}

func middleNode(head *ListNode) *ListNode {
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
