package arr

import "math"

/**
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。



示例 1：

输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
示例 2：

输入：target = 4, nums = [1,4,4]
输出：1
示例 3：

输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0


提示：

1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 105

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-size-subarray-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
暴力解法，直接两层for循环，时间复杂度是o(N^2)
*/

func minSubArrayLenBase(target int, nums []int) int {
	minLen := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		result := 0
		for j := i; j < len(nums); j++ {
			result += nums[j]
			if result >= target {
				subLen := j - i + 1
				if subLen < minLen {
					minLen = subLen
				}
				break
			}
		}
	}
	//未曾进过第二层循环
	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

/**
滑动窗口解法
通过动态的控制双指针指向的子序列的起止下标，不断计算出子序列和，并根据和是否大于target
如果小于target，则右下标继续前进增大窗口范围
如果大于等于target，则左下边继续前进缩小窗口范围
因为本题要求的是，连续子序列的最小长度，点题是"连续"
因此当第一个连续子串之和小于target的条件被破坏后
就应该先缩小滑动窗口后的左边界，重新推演当前子串之和
*/
func minSubArrayLenV2(target int, nums []int) int {
	i := 0
	l := len(nums)
	result := l + 1
	sum := 0
	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum >= target {
			subLen := j - i + 1
			if subLen < result {
				result = subLen
			}
			sum -= nums[i]
			i++
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}
}

/**
904. 水果成篮
在一排树中，第 i 棵树产生 tree[i] 型的水果。
你可以从你选择的任何树开始，然后重复执行以下步骤：

1.把这棵树上的水果放进你的篮子里。如果你做不到，就停下来。
2.移动到当前树右侧的下一棵树。如果右边没有树，就停下来。

请注意，在选择一颗树后，你没有任何选择：你必须执行步骤 1，然后执行步骤 2，然后返回步骤 1，然后执行步骤 2，依此类推，直至停止。

你有两个篮子，每个篮子可以携带任何数量的水果，但你希望每个篮子只携带一种类型的水果。
用这个程序你能收集的水果树的最大总量是多少？



示例 1：

输入：[1,2,1]
输出：3
解释：我们可以收集 [1,2,1]。
示例 2：

输入：[0,1,2,2]
输出：3
解释：我们可以收集 [1,2,2]
如果我们从第一棵树开始，我们将只能收集到 [0, 1]。
示例 3：

输入：[1,2,3,2,2]
输出：4
解释：我们可以收集 [2,3,2,2]
如果我们从第一棵树开始，我们将只能收集到 [1, 2]。
示例 4：

输入：[3,3,3,1,2,1,1,2,3,3,4]
输出：5
解释：我们可以收集 [1,2,1,1,2]
如果我们从第一棵树或第八棵树开始，我们将只能收集到 4 棵水果树。


提示：

1 <= tree.length <= 40000
0 <= tree[i] < tree.length

英文题解（本题英文比中文更好理解。。。）
You are visiting a farm that has a single row of fruit trees arranged from left to right. The trees are represented by an integer array fruits where fruits[i] is the type of fruit the ith tree produces.

You want to collect as much fruit as possible. However, the owner has some strict rules that you must follow:

You only have two baskets, and each basket can only hold a single type of fruit. There is no limit on the amount of fruit each basket can hold.
Starting from any tree of your choice, you must pick exactly one fruit from every tree (including the start tree) while moving to the right. The picked fruits must fit in one of your baskets.
Once you reach a tree with fruit that cannot fit in your baskets, you must stop.
Given the integer array fruits, return the maximum number of fruits you can pick.



Example 1:

Input: fruits = [1,2,1]
Output: 3
Explanation: We can pick from all 3 trees.
Example 2:

Input: fruits = [0,1,2,2]
Output: 3
Explanation: We can pick from trees [1,2,2].
If we had started at the first tree, we would only pick from trees [0,1].
Example 3:

Input: fruits = [1,2,3,2,2]
Output: 4
Explanation: We can pick from trees [2,3,2,2].
If we had started at the first tree, we would only pick from trees [1,2].
Example 4:

Input: fruits = [3,3,3,1,2,1,1,2,3,3,4]
Output: 5
Explanation: We can pick from trees [1,2,1,1,2].


Constraints:

1 <= fruits.length <= 105
0 <= fruits[i] < fruits.length

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fruit-into-baskets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

题目链接(https://leetcode-cn.com/problems/fruit-into-baskets/)
*/

func totalFruit(fruits []int) int {
	if len(fruits) == 0 {
		return 0
	}
	l := len(fruits)
	first, second, tempIndex := 0, 0, 0
	minLen := math.MinInt32
	for i := 0; i < l; i++ {
		fruitType := fruits[i]
		if fruitType != fruits[first] && fruitType != fruits[second] {
			first = tempIndex
			second = i
		}
		subLen := i - first + 1
		if subLen > minLen {
			minLen = subLen
		}
		if fruitType != fruits[tempIndex] {
			tempIndex = i
		}
	}
	return minLen
}
