package arr

/**
移除数组中的元素
*/

/**
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。



说明:

为什么返回数值是整数，但输出的答案是数组呢?

请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下:

// nums 是以“引用”方式传递的。也就是说，不对实参作任何拷贝
int len = removeElement(nums, val);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}


示例 1：

输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2]
解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。你不需要考虑数组中超出新长度后面的元素。例如，函数返回的新长度为 2 ，而 nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。
示例 2：

输入：nums = [0,1,2,2,3,0,4,2], val = 2
输出：5, nums = [0,1,4,0,3]
解释：函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。注意这五个元素可为任意顺序。你不需要考虑数组中超出新长度后面的元素。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//解法1，两层for循环暴力求解
//nums = [0,1,2,2,3,0,4,2] val = 2
//nums = [0,1,2,3,0,4,2, 2]
//nums = [0,1,3,0,4,2, 2,2]
func removeElement(nums []int, val int) int {
	end := len(nums)
	for i := 0; i < end; i++ {
		num := nums[i]
		if num == val {
			for j := i + 1; j < end; j++ {
				nums[j-1] = nums[j]
			}
			end--
			//因为无法确定向前迁移的第一个元素的值是否等同目标值，因此需要将指针自减指向原目标值的下标
			i--
		}
	}
	return end
}

//解法2，快慢指针法，常应用于字符串，数组，链表等题的
//思路，准备快慢指针
//（1）快慢指针的起点都在0
//（2）我们仅使用快指针遍历数组
// (3) 如果快指针的下标元素不等于目标值，则做一个赋值操作后，arr[慢指针]=arr[快指针]，并对慢指针做自增。
// (4) 如果快指针的下标元素等于目标值，则慢指针原地不动
// (5) 这就意味着，当发现一个目标值时，快指针前进了，慢指针原地不动，快指针-慢指针的差值恰好就是目标值的个数
// (6) 因为每次快指针的移动（快指针不指向目标值的情况下）都伴随着，arr[慢指针]=arr[快指针]，因此慢指针指向的目标值元素总会被快指针指向的非目标值覆盖，相当于该元素被删除了。
// 总结一下就是：慢指针停滞的时候，指向的是第一个目标值，当快指针找到非目标值时，将会第一时间覆盖慢指针的下标元素，使得第一个目标值被删除，第一个非目标值位置前移。
// 只要快指针扫描不到非目标值，因为arr[慢指针]=arr[快指针]的操作，则从第一个非目标值开始，后续的所有元素都会向前移动，偏移量为快慢指针的差值，相当于渐近式地完成了数组元素的迁移。
// 当快指针遍历完成时，目标元素的覆盖也就完成了（数组元素的删除只能通过覆盖，因此一般使用len来标识动态数组的可用长度,每次删除元素就是len-1+数组元素向前移动），此时慢指针的值为剩余数组的长度
func removeElementDoublePointer(nums []int, val int) int {
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		num := nums[fastIndex]
		if num != val {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	return slowIndex
}

/**
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。



说明:

为什么返回数值是整数，但输出的答案是数组呢?

请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下:

// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}

示例 1：

输入：nums = [1,1,2]
输出：2, nums = [1,2]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
示例 2：

输入：nums = [0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。


提示：

0 <= nums.length <= 3 * 104
-104 <= nums[i] <= 104
nums 已按升序排列

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func removeDuplicates(nums []int) int {
	end := len(nums)
	for i := 1; i < end; i++ {
		lastValue := nums[i-1]
		currentValue := nums[i]
		if lastValue == currentValue {
			for j := i; j < end; j++ {
				nums[j-1] = nums[j]
			}
			end--
			i--
		}
	}
	return end
}

/**
使用双指针法，慢指针指向可以被覆盖的元素，快指针用于遍历数组，只要arr[fast] != arr[fast]-1，则触发快慢指针的数据移动
*/
func removeDuplicatesV2(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	slow, fast := 1, 1
	for fast < len(nums) {
		curValue := nums[fast]
		prevValue := nums[fast-1]
		if curValue != prevValue {
			nums[slow] = curValue
			slow++
		}
		fast++
	}

	return slow
}
