package comparisons

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type LCMP struct {
	base.InstructionIndex0
}

func (i LCMP) Execute(frame *runtime.Frame) {
	val2 := frame.PopLong()
	val1 := frame.PopLong()
	if val1 > val2 {
		frame.PushInt(1)
	} else if val1 < val2 {
		frame.PushInt(-1)
	} else {
		frame.PushInt(0)
	}
}
