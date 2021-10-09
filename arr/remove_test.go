package arr

import "testing"

func TestRemoveElement(t *testing.T) {
	t.Log(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	t.Log(removeDuplicates(nums))
	t.Log(nums)
}

func TestRemoveDuplicatesV2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	t.Log(removeDuplicatesV2(nums))
	t.Log(nums)
}
