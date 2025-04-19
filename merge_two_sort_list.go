package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 1}
	l1Node2 := &ListNode{Val: 3}
	l1Node3 := &ListNode{Val: 6}
	l1Node4 := &ListNode{Val: 8}

	l1Node2.Next = l1Node3
	l1Node3.Next = l1Node4
	l1Node4.Next = nil
	l1.Next = l1Node2

	l2 := &ListNode{Val: 2}
	l2Node2 := &ListNode{Val: 2}
	l2Node3 := &ListNode{Val: 4}
	l2Node4 := &ListNode{Val: 7}
	l2Node5 := &ListNode{Val: 9}

	l2Node2.Next = l2Node3
	l2Node3.Next = l2Node4
	l2Node4.Next = l2Node5
	l2Node5.Next = nil

	l2.Next = l2Node2

	result := mergeTwoLists(l1, l2)

	current := result
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode
	var nodeValue int

	for list1 != nil || list2 != nil {

		if list1 != nil && list2 != nil {
			if list1.Val >= list2.Val {
				nodeValue = list2.Val
				list2 = list2.Next
			} else {
				nodeValue = list1.Val
				list1 = list1.Next
			}

		} else if list1 != nil {
			nodeValue = list1.Val
			list1 = list1.Next
		} else {
			nodeValue = list2.Val
			list2 = list2.Next
		}

		newNode := &ListNode{Val: nodeValue, Next: nil}

		if result == nil {
			result = newNode
		} else {
			current := result
			for current.Next != nil {
				current = current.Next
			}

			current.Next = newNode
		}
	}

	return result
}
