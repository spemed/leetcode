package arr

import "testing"

func TestRemoveElement(t *testing.T) {
	t.Log(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
