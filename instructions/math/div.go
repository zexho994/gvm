package math

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type IDIV struct {
	base.InstructionIndex0
}

type LDIV struct {
	base.InstructionIndex0
}

type DDIV struct {
	base.InstructionIndex0
}

type FDIV struct {
	base.InstructionIndex0
}

func (self *IDIV) Execute(frame *runtime.Frame) {
	v1 := frame.PopInt()
	v2 := frame.PopInt()
	if v1 == 0 {
		panic("[gvm][div] Arithmetic error ，divisor cannot be 0")
	}
	result := v2 / v1
	frame.PushInt(result)
}

func (self *DDIV) Execute(frame *runtime.Frame) {
	v1 := frame.PopDouble()
	v2 := frame.PopDouble()
	if v1 == 0 {
		panic("[gvm][div] Arithmetic error ，divisor cannot be 0")
	}
	result := v2 / v1
	frame.PushDouble(result)
}

func (self *LDIV) Execute(frame *runtime.Frame) {
	v1 := frame.PopLong()
	v2 := frame.PopLong()
	if v1 == 0 {
		panic("[gvm][div] Arithmetic error ，divisor cannot be 0")
	}
	result := v2 / v1
	frame.PushLong(result)
}

func (self *FDIV) Execute(frame *runtime.Frame) {
	v1 := frame.PopFloat()
	v2 := frame.PopFloat()
	if v1 == 0 {
		panic("[gvm][div] Arithmetic error ，divisor cannot be 0")
	}
	result := v2 / v1
	frame.PushFloat(result)
}
