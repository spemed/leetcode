package optional

type Optional struct {
	value interface{}
}

type Unpack func(interface{}) interface{}

func Of(value interface{}) Optional {
	return Optional{value: value}
}

func (opt *Optional) isNothing() bool {
	return opt.value == nil
}

func (opt *Optional) MayBe(failedVal interface{}) func(Unpack) interface{} {
	return func(unpack Unpack) interface{} {
		if opt.isNothing() {
			return failedVal
		} else {
			return unpack(opt.value)
		}
	}
}

func Add(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
