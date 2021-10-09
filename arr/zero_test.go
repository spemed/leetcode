package arr

import "testing"

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Log(nums)
}

func TestMoveZeroesV2(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroesV2(nums)
	t.Log(nums)
}
