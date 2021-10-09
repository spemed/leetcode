package arr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquaresBase(t *testing.T) {
	arr := []int{-4, -1, 0, 3, 10}
	assert.Equal(t, []int{0, 1, 9, 16, 100}, sortedSquaresBase(arr))
	arr1 := []int{-7, -3, 2, 3, 11}
	assert.Equal(t, []int{4, 9, 9, 49, 121}, sortedSquaresBase(arr1))
}

func TestSortedSquaresV2(t *testing.T) {
	arr := []int{-4, -1, 0, 3, 10}
	assert.Equal(t, []int{0, 1, 9, 16, 100}, sortedSquaresV2(arr))
	arr1 := []int{-7, -3, 2, 3, 11}
	assert.Equal(t, []int{4, 9, 9, 49, 121}, sortedSquaresV2(arr1))
}

/**
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
*/

func TestTotalFruits(t *testing.T) {
	assert.Equal(t, 3, totalFruit([]int{1, 2, 1}))
	assert.Equal(t, 3, totalFruit([]int{0, 1, 2, 2}))
	assert.Equal(t, 4, totalFruit([]int{1, 2, 3, 2, 2}))
	assert.Equal(t, 5, totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))
}

func TestMinWindowStandard(t *testing.T) {
	assert.Equal(t, "BANC", minWindowStandard("ADOBECODEBANC", "ABC"))
	assert.Equal(t, "a", minWindowStandard("a", "a"))
	assert.Equal(t, "", minWindowStandard("a", "aa"))
}
