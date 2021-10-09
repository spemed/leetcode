package optional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var emptyTest = new(Test)

func isValid(test interface{}) bool {
	return test != emptyTest
}

type Test struct {
	a int
	b int
}

func TestOf(t *testing.T) {
	valid := Of(&Test{})
	do1 := valid.MayBe(emptyTest)
	dd := do1(func(i interface{}) interface{} {
		dd := i.(*Test)
		dd.a = 1
		dd.b = 2
		return dd
	})
	assert.Equal(t, isValid(dd), true)

	valid1 := Of(nil)
	do2 := valid1.MayBe(emptyTest)
	dd1 := do2(func(i interface{}) interface{} {
		dd := i.(*Test)
		dd.a = 1
		dd.b = 2
		return dd
	})
	assert.Equal(t, isValid(dd1), false)
}

func TestAdd(t *testing.T) {
	addOne := Add(1)
	assert.Equal(t, 11, addOne(10))

	addTen := Add(10)
	assert.Equal(t, 60, addTen(50))
}
