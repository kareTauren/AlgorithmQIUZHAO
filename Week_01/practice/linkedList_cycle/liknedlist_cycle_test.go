package linkedList_cycle

import (
	"fmt"
	"testing"
)

type ListNode struct {
	val  int
	Next *ListNode
}

func New(val int) *ListNode {
	return &ListNode{val: val}
}

func hasCycle(head *ListNode) bool {
	cache := make(map[*ListNode]int)
	for head != nil {
		if _, ok := cache[head]; ok {
			return true
		}
		cache[head] = head.val
		head = head.Next
	}
	return false
}

func hasCycleSlowFast(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func (n *ListNode) add(val int) {
	last := n
	for last.Next != nil {
		last = last.Next
	}
	last.Next = &ListNode{val: val}
}

func Print(node *ListNode) {
	for node != nil {
		fmt.Println(node.val)
		node = node.Next
	}
}

func TestLinkedListCycle(t *testing.T) {
	node := New(3)
	node.add(2)
	node.add(0)
	node.add(-4)
	Print(node)
	fmt.Println(hasCycle(node))
}
