package home_work

import (
	"fmt"
	"testing"
)

func TestLeastInterval(t *testing.T) {
	tasks := []byte{
		'A', 'A', 'A',
		'B', 'B', 'B', 'C', 'C', 'D'}
	fmt.Println(leastInterval(tasks, 3))
}

// 这是解决此问题的绝佳方法。
// 我的理解：
//
// * 首先计算每个元素的出现次数。
// * 假设元素E的最大频率为M
// * 我们可以安排E的前M-1次出现，在E的连续计划之间，每个E之后至少要有N个CPU周期
// * 调度E-1的M-1次出现后的总CPU周期=（M-1）*（N +1）// 1代表E自身的CPU周期
// * 现在安排最后一轮任务。最后一次出现E至少需要1个CPU周期。如果有多个频率为M的任务，则它们都都需要再增加1个周期。
// * 遍历频率字典，并为每个具有频率== M的元素添加1个周期。
// * 如果我们的任务数量超过了上面计算的所需最大插槽数，则将返回task数组的长度，因为我们至少需要那么多个CPU周期。

func leastInterval(tasks []byte, n int) int {
	// counter := make(map[byte]int, 0)
	// max := 0
	// maxCount := 0
	// for i := 0; i < len(tasks); i++ {
	// 	task := tasks[i]
	// 	counter[task-'A']++
	// 	if max == counter[task-'A'] {
	// 		maxCount++
	// 	} else if max < counter[task-'A'] {
	// 		max = counter[task-'A']
	// 		maxCount = 1
	// 	}
	// }
	//
	// fmt.Println(max)
	// fmt.Println(maxCount)
	//
	// return 1
	c := make(map[byte]int, 0)
	max := 0
	for i := 0; i < len(tasks); i++ {
		c[tasks[i]]++
		if c[tasks[i]] > max {
			max = c[tasks[i]]
		}
	}

	maxFunc := func(n, m int) int {
		if n > m {
			return n
		}
		return m
	}

	ans := (max - 1) * (n + 1)
	for _, v := range c {
		if v == max {
			ans++
		}
	}
	return maxFunc(len(tasks), ans)
}
