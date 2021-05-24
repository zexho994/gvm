package math

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type ISUB struct {
	base.InstructionIndex0
}

type LSUB struct {
	base.InstructionIndex0
}

type DSUB struct {
	base.InstructionIndex0
}

type FSUB struct {
	base.InstructionIndex0
}

func (sub *ISUB) Execute(frame *runtime.Frame) {
	v1 := frame.PopInt()
	v2 := frame.PopInt()
	subResult := v2 - v1
	frame.PushInt(subResult)
}

func (sub *LSUB) Execute(frame *runtime.Frame) {
	v1 := frame.PopLong()
	v2 := frame.PopLong()
	subResule := v2 - v1
	frame.PushLong(subResule)
}

func (sub *DSUB) Execute(frame *runtime.Frame) {
	v1 := frame.PopDouble()
	v2 := frame.PopDouble()
	subResult := v2 - v1
	frame.PushDouble(subResult)
}

func (sub *FSUB) Execute(frame *runtime.Frame) {
	v1 := frame.PopFloat()
	v2 := frame.PopFloat()
	subResult := v2 - v1
	frame.PushFloat(subResult)
}
