package comparisons

import (
	"github.com/zouzhihao-994/gvm/src/vm/instructions/base"
	"github.com/zouzhihao-994/gvm/src/vm/runtime"
)

/*
用于比较long变量
*/
type LCMP struct {
	base.NoOperandsInstruction
}

func (self *LCMP) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	// 1大于，0等于，-1小于
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}

}
