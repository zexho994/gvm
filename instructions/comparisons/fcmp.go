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
	val2 := frame.PopFloat()
	val1 := frame.PopFloat()
	if math.IsNaN(float64(val1)) || math.IsNaN(float64(val2)) {
		frame.PushInt(1)
		return
	}
	if val1 > val2 {
		frame.PushInt(1)
	} else if val1 < val2 {
		frame.PushInt(-1)
	} else {
		frame.PushInt(0)
	}
}

func (i FCMPL) Execute(frame *runtime.Frame) {
	val2 := frame.PopFloat()
	val1 := frame.PopFloat()
	if math.IsNaN(float64(val1)) || math.IsNaN(float64(val2)) {
		frame.PushInt(-1)
		return
	}
	if val1 > val2 {
		frame.PushInt(1)
	} else if val1 < val2 {
		frame.PushInt(-1)
	} else {
		frame.PushInt(0)
	}
}
