package arr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackTrim(t *testing.T) {
	str := backTrim("ab#c")
	assert.Equal(t, "ac", str)

	str = backTrim("ad#c")
	assert.Equal(t, "ac", str)

	str = backTrim("ab##")
	assert.Equal(t, "", str)

	str = backTrim("a##c")
	assert.Equal(t, "c", str)

	str = backTrim("a#c")
	assert.Equal(t, "c", str)
}

func TestA(t *testing.T) {
	assert.Equal(t, false, backspaceCompareV2("bbbextm", "bbb#extm"))
	assert.Equal(t, true, backspaceCompareV2("ac", "ad#c"))
	assert.Equal(t, true, backspaceCompareV2("", "ab##"))
	assert.Equal(t, true, backspaceCompareV2("c", "a##c"))
	assert.Equal(t, true, backspaceCompareV2("c", "a#c"))
}
