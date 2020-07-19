package reverse_linkedList

import (
	"fmt"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	node := New(1)
	node.add(2)
	node.add(3)
	node.add(4)
	node.add(5)
	PrintNode(node)
	// node = reverseList(node)
	node = Reverse2(node)
	PrintNode(node)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(value int) *ListNode {
	return &ListNode{Val: value}
}

func (node *ListNode) add(value int) {
	last := node
	for last.Next != nil {
		last = last.Next
	}
	last.Next = &ListNode{Val: value}
}

func PrintNode(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current := head
	var next, prev *ListNode
	for head != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func Reverse2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	for head != nil {
		current := head
		head = head.Next
		current.Next = prev
		prev = current
	}
	return prev
}

// prev = None
// while head:
// curr = head
// head = head.next
// curr.next = prev
// prev = curr
// return prev
