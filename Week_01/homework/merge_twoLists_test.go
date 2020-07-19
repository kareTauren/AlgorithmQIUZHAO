package homework

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode {
	return &ListNode{Val: val}
}

func (head *ListNode) add(val int) {
	for head.Next != nil {
		head = head.Next
	}
	head.Next = &ListNode{Val: val}
}

func (head *ListNode) Print() {
	current := head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

func TestMergeTwoLists(t *testing.T) {
	first := New(1)
	first.add(2)
	first.add(4)
	// first.Print()

	second := New(1)
	second.add(3)
	second.add(4)
	// second.Print()

	// 1->1->2->3->4->4
	new := mergeTwoLists(first, second)
	new.Print()

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l2 != nil {
		prev.Next = l2
	} else if l1 != nil {
		prev.Next = l1
	}

	return dummy.Next
}
