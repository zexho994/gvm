package math

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type IMUL struct {
	base.InstructionIndex0
}

type FMUL struct {
	base.InstructionIndex0
}

type DMUL struct {
	base.InstructionIndex0
}

type LMUL struct {
	base.InstructionIndex0
}

func (self *IMUL) Execute(frame *runtime.Frame) {
	v1 := frame.PopInt()
	v2 := frame.PopInt()
	result := v1 * v2
	frame.PushInt(result)
}

func (self *LMUL) Execute(frame *runtime.Frame) {
	v1 := frame.PopLong()
	v2 := frame.PopLong()
	result := v1 * v2
	frame.PushLong(result)
}

func (self *DMUL) Execute(frame *runtime.Frame) {
	v1 := frame.PopDouble()
	v2 := frame.PopDouble()
	result := v1 * v2
	frame.PushDouble(result)
}

func (self *FMUL) Execute(frame *runtime.Frame) {
	v1 := frame.PopFloat()
	v2 := frame.PopFloat()
	result := v1 * v2
	frame.PushFloat(result)
}
