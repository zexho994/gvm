package comparisons

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

type IF_GE struct {
	base.InstructionIndex16
}

func (i IF_GE) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, int(i.Index))
	}
}
