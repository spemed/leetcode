package arr

import "testing"

func TestBinarySearch(t *testing.T) {
	target := 20
	var nums []int
	t.Log(binarySearch(nums,target))
	nums = []int{1,2,3,4,5,6,7,8,9}
	t.Log(binarySearch(nums,target))
	nums = []int{1,2,3,4,20,6,7,8,9}
	t.Log(binarySearch(nums,target))
	nums = []int{20}
	t.Log(binarySearch(nums,target))
}


func TestSearchInsert(t *testing.T) {
	target := 20
	var nums []int
	t.Log(searchInsert(nums,target))
	nums = []int{1,2,3,4,5,6,7,8,9}
	t.Log(searchInsert(nums,target))
	nums = []int{1,2,3,4,20,6,7,8,9}
	t.Log(searchInsert(nums,target))
	nums = []int{20}
	t.Log(searchInsert(nums,target))
}


func TestSearchRange(t *testing.T) {
	var nums []int
	//t.Log(searchRange(nums,9))
	//nums = []int{5,7,7,8,8,10}
	//t.Log(searchRange(nums,8))
	//nums = []int{5,7,7,8,8,10}
	//t.Log(searchRange(nums,6))
	nums = []int{1,1,2}
	t.Log(searchRange(nums,1))
}

func TestMySqrt(t *testing.T) {
	mySqrt(4)
}

func TestIsPerfectSquare(t *testing.T) {
	t.Log(isPerfectSquare(16))
	t.Log(isPerfectSquare(12))
}