package mergeTwoLists

import (
	"LeetCodePatterns/Easy/utils"
	"fmt"
)

func Example() {
	fmt.Println("\nMerge Two Sorted Lists")
	arr := []int{1, 2, 4, 5}
	head1 := utils.ArrayToList(arr, len(arr))
	fmt.Println("list1 =", arr)
	arr = []int{1, 2, 3, 6}
	head2 := utils.ArrayToList(arr, len(arr))
	fmt.Println("list2 =", arr)
	result := mergeTwoLists(head1, head2)
	fmt.Println("result = " + result.Show())
}

func mergeTwoLists(list1 *utils.ListNode, list2 *utils.ListNode) *utils.ListNode {
	var node *utils.ListNode
	var cur1 *utils.ListNode
	var cur2 *utils.ListNode
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		cur1 = list1.Next
		cur2 = list2
		node = list1
	} else {
		cur1 = list2.Next
		cur2 = list1
		node = list2
	}
	next, nol, check := compare(cur1, cur2)
	node.Next = next
	if !nol {
		if check {
			cur1 = cur1.Next
		} else {
			cur2 = cur2.Next
		}
		if node.Next != nil {
			for i := next; i != nil || nol; {
				next, nol, check = compare(cur1, cur2)
				i.Next = next
				i = i.Next
				if nol {
					break
				}
				if check {
					cur1 = cur1.Next
				} else {
					cur2 = cur2.Next
				}
			}
		}
	}
	return node
}

func compare(cur1 *utils.ListNode, cur2 *utils.ListNode) (*utils.ListNode, bool, bool) {
	if cur1 == nil {
		return cur2, true, true
	}
	if cur2 == nil {
		return cur1, true, true
	}
	if cur1.Val <= cur2.Val {
		return cur1, false, true
	} else {
		return cur2, false, false
	}
}
