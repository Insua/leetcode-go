package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	tail = head
	carry := 0
	for true {
		p1 := 0
		p2 := 0
		if l1 != nil {
			p1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			p2 = l2.Val
			l2 = l2.Next
		}
		sum := p1 + p2 + carry
		onePlace := sum % 10
		carry = sum / 10
		l := &ListNode{Val: onePlace}
		if tail == nil {
			head = l
			tail = head
		} else {
			tail.Next = l
			tail = l
		}

		if l1 == nil && l2 == nil {
			if carry > 0 {
				cl := &ListNode{Val: carry}
				tail.Next = cl
			}
			break
		}
	}

	return head
}

func fillList(nums []int) *ListNode {
	var head = new(ListNode)
	head.Val = nums[0]
	var tail *ListNode
	tail = head
	for i := 1; i < len(nums); i++ {
		var node = ListNode{Val: nums[i]}
		(*tail).Next = &node
		tail = &node
	}
	return head
}

func dump(l *ListNode) {
	for l != nil {
		fmt.Println(*l)
		l = l.Next
	}
}

func main() {
	l1 := fillList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	l2 := fillList([]int{5, 6, 4})

	l := addTwoNumbers(l1, l2)

	dump(l)
}
