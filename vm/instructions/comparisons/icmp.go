package comparisons

import (
	"../../instructions/base"
	"../../rtda"
)

/*
用于比较long变量
*/
type ICMP struct {
	base.NoOperandsInstruction
}

func (self *ICMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	// 1大于，0等于，-1小于
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}