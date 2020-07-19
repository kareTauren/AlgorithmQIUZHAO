package swap_Nodes_in_Pairs

import (
	"fmt"
	"strings"
	"testing"
)

type ListNode struct {
	val  int
	Next *ListNode
}

func New(val int) *ListNode {
	return &ListNode{val: val}
}

func Print(node *ListNode) {
	for node != nil {
		fmt.Println(node.val)
		node = node.Next
	}
}

func (head *ListNode) add(val int) {
	node := head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = New(val)
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head
	current := dummy
	for current.Next != nil && current.Next.Next != nil {
		first := current.Next       // 1,2,3,4
		second := current.Next.Next // 2,3,4
		first.Next = second.Next    // 1,3,4 作用是：把 2 提出来
		current.Next = second       // 0,2,3,4
		current.Next.Next = first   // 0,2,1,3,4 作用是构成一个颠倒第一个和第二个数据位置的链表
		current = current.Next.Next // 1,3,4 // 去掉头两个数据，因为下一轮要操作的数据是 3，4，需要1占位。
	}
	return dummy.Next
}

func swapPairsUseRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// n := head.Next
	// Print(head)
	// tmp := swapPairsUseRecursion(head.Next.Next)
	// head.Next = tmp
	// n.Next = head

	n := head.Next
	head.Next = swapPairs(head.Next.Next)
	n.Next = head
	return n

}

func TestSwapPairs(t *testing.T) {
	node := New(1)
	node.add(2)
	node.add(3)
	node.add(4)
	node.add(5)
	Print(node)
	fmt.Println("ffff")
	// node = swapPairs(node)
	node = swapPairsUseRecursion(node)
	Print(node)

}

func TestXXX(t *testing.T) {
	s := "hello world!"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	if len(s) < 1 {
		return ""
	}

	sAry := strings.Split(s, " ")

	str := ""
	for i := len(sAry) - 1; i >= 0; i-- {
		if sAry[i] == "" {
			continue
		}
		if i != 0 {
			str += sAry[i] + " "
		} else {
			str += sAry[i]
		}
	}
	if len(str) > 0 {
		idx := strings.LastIndex(str, " ")
		if idx == len(str)-1 {
			str = str[:len(str)-1]
		}
	}
	return str
}

func TestXx(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(check(s))
}

func check(s string) bool {
	rs := strings.ToLower(s)
	isAlum := func(b byte) bool {
		if (b >= 'A' && b <= 'Z') ||
			(b >= 'a' && b <= 'z') ||
			(b >= '0' && b <= '9') {
			return true
		}
		return false
	}

	l := 0
	r := len(rs) - 1

	for l < r {
		for l < r && !isAlum(rs[l]) {
			l++
		}
		for l < r && !isAlum(rs[r]) {
			r--
		}
		if l < r {
			if rs[l] != rs[r] {
				return false
			}
			l++
			r--
		}
	}
	return true
}
