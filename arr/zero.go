package arr

/**
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/move-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//暴力解法，只要轮询到当前元素是0，就从倒数第n位开始向前移动，最终再把倒数第n位置为0（n为0已经出现过的次数）
func moveZeroes(nums []int) {
	end := len(nums) - 1
	for i := 0; i <= end; i++ {
		v := nums[i]
		if v == 0 {
			j := i + 1
			for ; j <= end; j++ {
				nums[j-1] = nums[j]
			}
			nums[end] = 0
			end--
			i--
		}
	}
}

func moveZeroesV2(nums []int) {
	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
