package homework

import (
	"fmt"
	"testing"
)

// 未完成 现在的是一个普通的队列实现，不是deque, 也不是循环双端队列
type MyCircularDeque struct {
	size  int
	queue []int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{size: k}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if len(this.queue) == this.size {
		return false
	}

	if this.size != 0 {
		tmp := make([]int, 1)
		tmp[0] = value
		tmp = append(tmp, this.queue...)
		this.queue = tmp
		return true
	}
	this.queue[0] = value
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if len(this.queue) >= this.size {
		return false
	}
	this.queue = append(this.queue, value)
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if len(this.queue) <= 0 {
		return false
	}

	this.queue = this.queue[1 : len(this.queue)-1]
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if len(this.queue) <= 0 {
		return false
	}

	this.queue = this.queue[:len(this.queue)-1]
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.size == 0 {
		return -1
	}
	front := this.queue[:1]
	return front[0]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.size == 0 {
		return -1
	}
	front := this.queue[:len(this.queue)-1]
	return front[0]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if len(this.queue) > 0 {
		return false
	}
	return true
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if len(this.queue) == this.size {
		return true
	}
	return false
}

func TestMyCircularDeque(t *testing.T) {
	obj := Constructor(3)
	fmt.Println(obj.InsertFront(1))
	fmt.Println(obj.InsertLast(2))
	fmt.Println(obj.InsertLast(3))
	fmt.Println(obj.InsertLast(4))
	fmt.Println("rear: ", obj.GetRear())
	fmt.Println(obj.IsFull())
	fmt.Println(obj.DeleteLast())
	fmt.Println(obj.InsertFront(4))
	fmt.Println(obj.GetFront())
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
