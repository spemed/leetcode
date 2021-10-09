package arr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSubArrayLenBase(t *testing.T) {
	arr := []int{2, 3, 1, 2, 4, 3}
	assert.Equal(t, 2, minSubArrayLenBase(7, arr))
	arr1 := []int{1, 4, 4}
	assert.Equal(t, 1, minSubArrayLenBase(4, arr1))
	arr2 := []int{1, 1, 1, 1, 1, 1, 1, 1}
	assert.Equal(t, 0, minSubArrayLenBase(11, arr2))
}

func TestMinSubArrayLenV2(t *testing.T) {
	arr := []int{2, 3, 1, 2, 4, 3}
	assert.Equal(t, 2, minSubArrayLenV2(7, arr))
	arr1 := []int{1, 4, 4}
	assert.Equal(t, 1, minSubArrayLenV2(4, arr1))
	arr2 := []int{1, 1, 1, 1, 1, 1, 1, 1}
	assert.Equal(t, 0, minSubArrayLenV2(11, arr2))
}
