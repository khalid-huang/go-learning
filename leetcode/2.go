// 就是一个进位加法，要学会如何组织代码
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {
	sum, carry := 0, 0
	head := &ListNode{}
	cur := head
	for {
		sum, carry = add(l1, l2, carry)
		cur.Val = sum
		l1 = next(l1)
		l2 = next(l2)
		if l1 == nil && l2 == nil {
			break
		}
		cur.Next = &ListNode{}
		cur = cur.Next
	}
	if carry == 1 {
		cur.Next = &ListNode{Val: carry}
	}
	return head
}

func next(l *ListNode) *ListNode {
	if l != nil {
		return l.Next
	}
	return nil
}

func add(l1 *ListNode, l2 *ListNode, preCarry int) (sum, carry int) {
	if l1 != nil {
		sum += l1.Val
	}

	if l2 != nil {
		sum += l2.Val
	}

	sum += preCarry

	if sum > 9 {
		sum -= 10
		carry = 1
	}

	return
}

func test_2() {
	l11 := &ListNode{Val: 2}
	l12 := &ListNode{Val: 4}
	l13 := &ListNode{Val: 3}
	l11.Next = l12
	l12.Next = l13

	l21 := &ListNode{Val: 5}
	l22 := &ListNode{Val: 6}
	l23 := &ListNode{Val: 4}
	l21.Next = l22
	l22.Next = l23

	result := addTwoNumber(l11, l21)
	cur := result
	for {
		if cur != nil {
			fmt.Print(cur.Val)
			cur = cur.Next
		}
	}

}
