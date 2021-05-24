package math

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type IADD struct {
	base.InstructionIndex0
}

type LADD struct {
	base.InstructionIndex0
}

type FADD struct {
	base.InstructionIndex0
}

type DADD struct {
	base.InstructionIndex0
}

func (this *IADD) Execute(frame *runtime.Frame) {
	v1 := frame.PopInt()
	v2 := frame.PopInt()
	addResult := v1 + v2
	frame.PushInt(addResult)
}

func (this *LADD) Execute(frame *runtime.Frame) {
	v1 := frame.PopLong()
	v2 := frame.PopLong()
	addResult := v1 + v2
	frame.PushLong(addResult)
}

func (this *FADD) Execute(frame *runtime.Frame) {
	v1 := frame.PopFloat()
	v2 := frame.PopFloat()
	addResult := v1 + v2
	frame.PushFloat(addResult)
}

func (this *DADD) Execute(frame *runtime.Frame) {
	v1 := frame.PopDouble()
	v2 := frame.PopDouble()
	addResult := v1 + v2
	frame.PushDouble(addResult)
}
