#### 分治
分而治之，就是把一个可以使用重复方法解决的问题，分成多个部分去执行，然后把得到的结果合并成一个总的结果集
典型例题：求 Pow(x,y)
```golang
func pow(x, n) {
	if n < 0 {
		return 1 / divide(x, -n)
	}
	return divide(x, n)
}
func divide(x, n) int {
	if n == 0 {
		return 1
	}
	ans := divide(x, n/2)
	if n&1 == 0 {
		return ans * ans
	}
	return ans * ans * x
}
```
#### 回溯
其实是一种暴力求解的算法思想，构建递归的组成形式，比如求解[2,3,4]的两两组合情况，他会和数组中的每个元素进行组合，然后求出结果，回溯的时间复杂度通常会很高。
#### 广度优先
比如要写一个抓取全国省市区的爬虫，如果你是按照先抓取所有省，再抓取所有市，再抓取所有县，就像你扔了一个石头在河里，它的水波纹是一圈圈往外扩散，会越来越广。
#### 深度优先
还是和爬虫一样，比如要写一个抓取全国省市区的爬虫，比如你要抓取陕西的信息，那么就是陕西省，西安市，兰湖区，就是先抓完一个地点的省，市，区，再去抓四川，或者湖南的信息。
#### 贪心
利益最大化，当前最优解，典型代表，股票买卖问题。
#### 二分查找
一般使用二分需要满足一下条件
1. 单调递增，或递减
2. 有界限
3. 可以使用索引访问
代码模板
```golang
nums = make([]int, 100)
l := 0
r := len(nums) - 1
for l <= r {
	mid := (r + l)/2
	if nums[mid] == target {
		return xxxx
	}
	if nums[mid] < target {
		l = mid + 1
	} else {
		r = mid - 1
	}
}	
```
以上是在一个有序数组中查找一个元素是否存在的最基本模板。

HomeWork 完成情况

◉	[多数元素](https://leetcode-cn.com/problems/majority-element/description/)

◉	[柠檬水找零](https://leetcode-cn.com/problems/lemonade-change/description/)

◉	[买卖股票的最佳时机 2](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/description/)

◉	[分发饼干](https://leetcode-cn.com/problems/assign-cookies/description/)

◉	[模拟行走机器人](https://leetcode-cn.com/problems/walking-robot-simulation/description/)

◉	[Pow(x, y)](https://leetcode-cn.com/problems/powx-n/)

◉	[子集](https://leetcode-cn.com/problems/subsets/)

◉	[电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)

◎	[单词接龙](https://leetcode-cn.com/problems/word-ladder/description/)

◉	[岛屿个数](https://leetcode-cn.com/problems/number-of-islands/)

◎	[扫雷游戏](https://leetcode-cn.com/problems/minesweeper/description/)

◎	[跳跃游戏](https://leetcode-cn.com/problems/jump-game/)

◎	[搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)

◉	[搜索二维矩阵](https://leetcode-cn.com/problems/search-a-2d-matrix/)

◉	[寻找旋转排序数组中的最小值](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)

◉	[N queen](https://leetcode-cn.com/problems/n-queens/)

◎	[单词接龙](https://leetcode-cn.com/problems/word-ladder-ii/description/)2

◎	[跳跃游戏2](https://leetcode-cn.com/problems/jump-game-ii/)

