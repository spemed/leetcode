package arr

/**
leetCode 704. 二分查找
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

//因为数组已经是有序的，固我们可以使用二分查找的方式做到o(logn)的时间复杂度（对于已经排序好的，确定数据规模的静态数据非常实用）

func binarySearch(nums []int, target int) int {
	left,right := 0,len(nums)-1
	for right >= left {
		mid := ((right - left) / 2) + left
		midVal := nums[mid]
		if midVal > target {
			right = mid - 1
		} else if midVal < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}


//以下的题目为二分算法相关的推荐题目

/**
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法。

 

示例 1:

输入: nums = [1,3,5,6], target = 5
输出: 2
示例 2:

输入: nums = [1,3,5,6], target = 2
输出: 1
示例 3:

输入: nums = [1,3,5,6], target = 7
输出: 4
示例 4:

输入: nums = [1,3,5,6], target = 0
输出: 0
示例 5:

输入: nums = [1], target = 0
输出: 0
 

提示:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 为无重复元素的升序排列数组
-104 <= target <= 104


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-insert-position
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

使用二分算法寻找元素，如果找到了则返回该元素的下标，如果找不到呢，则返回left的下标。
 */

func searchInsert(nums []int, target int) int {
	left,right := 0,len(nums)-1
	for right >= left {
		mid := ((right - left) / 2) + left
		midVal := nums[mid]
		if midVal > target {
			right = mid - 1
		} else if midVal < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}


/**
34. 在排序数组中查找元素的第一个和最后一个位置
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：
你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？


示例 1：
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：

你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？


输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]


提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

二分查找适用于有序的数组（容器的长度已知，且不会动态变化）。
思路：
	（1）二分查找是将数组进行排序后，通过取中值同目标值做比较，如果等于目标值则搜索完成。如果比目标值小就往右区间继续搜索，比目标值大则往左区间继续搜索，直到左区间的指针大于右区间的指针时判定为元素搜索不到。
    （2）因此使用二分查找一定可以得到一个准确的结果，即是目标值在不在数组中
    （3）本题相比传统的二分法有一个小变化，变化的地方在于搜索的是若干个target而不是一个target
    （4）因此我们需要执行两次二分查找来确定target区间的起始坐标和终止坐标，分别记为startIndex和endIndex
     (5) 我们首先确定startIndex的位置，二分查找时，如果判断到target值等于mid值，我们使用临时变量记录temp记录下当前mid的值，令右区间减一后继续搜索target值，如果再次搜索到target值则继续更新临时变量的下标，直到搜索结束（二分查找一定可以得到在or不在的两个结果），此时tempIndex为startIndex
    （6）然后确定endIndex的位置，二分查找时，如果判断到target值等于mid值，我们使用临时变量记录temp记录下当前mid的值，令左区间加一后继续搜索target值，如果再次搜索到target值则继续更新临时变量的下标，直到搜索结束（二分查找一定可以得到在or不在的两个结果），此时tempIndex为endIndex
     (7) 返回区间
 */

func searchRange(nums []int, target int) []int {
	search := func(nums[]int,target int,lower bool) int {
		l,r := 0,len(nums)-1
		ans := -1
		for l <= r {
			mid := (r + l) >> 1
			midVal := nums[mid]
			if  target < midVal {
				r = mid - 1
			} else if target > midVal {
				l = mid + 1
			} else {
				ans = mid
				if lower {
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
		}
		return ans
	}

	searchLeft := func(nums[]int,target int) int {
		return search(nums,target,true)
	}

	searchRight := func(nums[]int,target int) int {
		return search(nums,target,false)
	}
	return []int{searchLeft(nums,target),searchRight(nums,target)}
}

/**
69. x 的平方根
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2

示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
     由于返回类型是整数，小数部分将被舍去。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sqrtx
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

给定目标值为x的情况下，假设存在整数的集合K，对于K中的每一个元素k，都满足k^2 <= x，取得集合K中整数的最大值则为x的平方根
k的上下限可以进行模糊设置为0和x本身
 */

func mySqrt(x int) int {
	left,right := 0,x
	ans := -1
	for left <= right {
		mid := (right + left) >> 1
		if mid * mid > x {
			right = mid - 1
		} else {
			ans = mid
			left = mid + 1
		}
	}
	return ans
}


/**
给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。

进阶：不要 使用任何内置的库函数，如  sqrt 。

 

示例 1：

输入：num = 16
输出：true
示例 2：

输入：num = 14
输出：false
 

提示：

1 <= num <= 2^31 - 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-perfect-square
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func isPerfectSquare(num int) bool {
	left,right := 0,num
	for left <= right {
		mid := (left + right) >> 1
		midVal := mid * mid
		if midVal < num {
			left = mid + 1
		} else if midVal > num {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}