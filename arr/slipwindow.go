package arr

import (
	"math"
	"sort"
)

/**
977. 有序数组的平方
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。



示例 1：

输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]
示例 2：

输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]


提示：

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 已按 非递减顺序 排序


进阶：

请你设计时间复杂度为 O(n) 的算法解决本问题
*/

//基础版，遍历数组每项元素进行平方后，在对数组进行排序，时间复杂度是o(n+logn)
func sortedSquaresBase(nums []int) []int {
	for i, v := range nums {
		nums[i] = v * v
	}
	sort.Ints(nums)
	return nums
}

//进阶版，使用双指针法，需要额外的空间复杂度o(n),但只需要遍历一次数组，因此时间复杂度是o(n)
//核心思路：已知数组是非递减顺序，设数组元素的区间范围为[L,R],L和R为整数。
//当|L| <= |R|时，最大值出现在数组的最右侧
//当|L| >= |R|时，最大值出现在数组的最左侧
//取双指针，从L和R出发，判断平方值的大小，而后填充到临时数组中，直到左指针下标大于右指针下标时算法结束
func sortedSquaresV2(nums []int) []int {
	start, end := 0, len(nums)-1
	temp := make([]int, len(nums))
	i := len(nums) - 1
	for start <= end {
		left := nums[start] * nums[start]
		right := nums[end] * nums[end]
		if left > right {
			temp[i] = left
			start++
		} else {
			temp[i] = right
			end--
		}
		i--
	}
	return temp
}

/**
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。



注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。


示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
示例 2：

输入：s = "a", t = "a"
输出："a"
示例 3:

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。


提示：

1 <= s.length, t.length <= 105
s 和 t 由英文字母组成


进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-window-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minWindowStandard(s string, t string) string {
	ori, cnt := make(map[byte]int), make(map[byte]int)
	for i := range t {
		v := t[i]
		ori[v]++
	}

	lStart, lEnd := -1, -1
	subLen := math.MaxInt32
	sLen := len(s)

	check := func() bool {
		for i := range ori {
			if cnt[i] < ori[i] {
				return false
			}
		}
		return true
	}

	for l, r := 0, 0; r < sLen; r++ {
		_, ok := ori[s[r]]
		if ok {
			cnt[s[r]]++
		}
		for check() && l <= r {
			tempLen := r - l + 1
			if tempLen < subLen {
				subLen = tempLen
				lStart, lEnd = l, l+tempLen
			}
			cnt[s[l]]--
			l++
		}
	}

	if lStart == -1 {
		return ""
	}

	return s[lStart:lEnd]
}
