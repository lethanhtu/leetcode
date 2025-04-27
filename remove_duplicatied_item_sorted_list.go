package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := buildList([]int{})
	browseList(node)
	node = deleteDuplicates(node)
	browseList(node)

	// fmt.Println("=====================")

	// node = buildList([]int{2, 4, 6, 8, 8})
	// browseList(node)
	// node = deleteDuplicates(node)
	// browseList(node)

	// fmt.Println("=====================")

	// node = buildList([]int{2, 4, 6, 6, 6, 6, 6, 6, 10})
	// browseList(node)
	// node = deleteDuplicates(node)
	// browseList(node)

	// fmt.Println("=====================")

	// node = buildList([]int{2, 4, 6, 6, 6, 6, 6, 6, 6, 6})
	// browseList(node)
	// node = deleteDuplicates(node)
	// browseList(node)

	// fmt.Println("=====================")

	// node = buildList([]int{2, 4})
	// browseList(node)
	// node = deleteDuplicates(node)
	// browseList(node)

	// fmt.Println("=====================")

	// node = buildList([]int{2})
	// browseList(node)
	// node = deleteDuplicates(node)
	// browseList(node)

	// fmt.Println("=====================")

	// node = buildList([]int{2, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6})
	// browseList(node)
	// node = deleteDuplicates(node)
	// browseList(node)

	// fmt.Println("=====================")

	// node = buildList([]int{2, 2})
	// browseList(node)
	// node = deleteDuplicates(node)
	// browseList(node)

	// fmt.Println("=====================")

	// node = buildList([]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 4})
	// browseList(node)
	// node = deleteDuplicates(node)
	// browseList(node)

	// fmt.Println("=====================")

}

func buildList(items []int) *ListNode {
	var result *ListNode
	var current *ListNode

	for _, value := range items {
		if result == nil {
			result = &ListNode{Val: value, Next: nil}
			continue
		}

		current = result
		for current.Next != nil {
			current = current.Next
		}

		current.Next = &ListNode{Val: value, Next: nil}

	}
	fmt.Println(result)

	return result
}

func browseList(node *ListNode) {
	result := ""
	chain := ""
	for node != nil {
		if result != "" {
			chain = "->"
		} else {
			chain = ""
		}

		result = result + chain + strconv.Itoa(node.Val)
		node = node.Next
	}

	fmt.Println(result)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var nodeAfterNext *ListNode
	current := head
	for current.Next != nil {
		if current.Val != current.Next.Val {
			current = current.Next
			continue
		}

		if current.Next.Next == nil {
			current.Next = nil
			return head
		} else {
			nodeAfterNext = current.Next.Next
			current.Next = nodeAfterNext
		}
	}

	return head
}
