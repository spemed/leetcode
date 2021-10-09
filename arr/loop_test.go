package arr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMatrix(t *testing.T) {
	actual := generateMatrix(3)
	assert.Equal(t, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}, actual)
}
