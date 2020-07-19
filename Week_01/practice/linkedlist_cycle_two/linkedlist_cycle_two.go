package linkedlist_cycle_two

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		if fast.Next != nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		// 第一次相遇了，就证明一定有环
		if fast == slow {
			for head != fast {
				// head 从头开始，然后 fast 依然在环内，然后每一次各走一步，最终肯定会相遇
				head = head.Next
				fast = fast.Next
			}
			return fast
		}
	}
	return nil
}

// https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/linked-list-cycle-ii-kuai-man-zhi-zhen-shuang-zhi-/
