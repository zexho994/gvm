package comparisons

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type LCMP struct {
	base.InstructionIndex0
}

func (i LCMP) Execute(frame *runtime.Frame) {
	stack := frame
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	if val1 > val2 {
		stack.PushInt(1)
	} else if val1 < val2 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(0)
	}
}
