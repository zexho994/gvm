package comparisons

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
	"math"
)

type FCMPG struct {
	base.InstructionIndex0
}

type FCMPL struct {
	base.InstructionIndex0
}

func (i FCMPG) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopFloat()
	val1 := stack.PopFloat()
	if math.IsNaN(float64(val1)) || math.IsNaN(float64(val2)) {
		stack.PushInt(1)
		return
	}
	if val1 > val2 {
		stack.PushInt(1)
	} else if val1 < val2 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(0)
	}
}

func (i FCMPL) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopFloat()
	val1 := stack.PopFloat()
	if math.IsNaN(float64(val1)) || math.IsNaN(float64(val2)) {
		stack.PushInt(-1)
		return
	}
	if val1 > val2 {
		stack.PushInt(1)
	} else if val1 < val2 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(0)
	}
}
