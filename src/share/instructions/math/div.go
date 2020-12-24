package math

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
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
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	if v1 == 0 {
		panic("[gvm][div] Arithmetic error ，divisor cannot be 0")
	}
	result := v2 / v1
	stack.PushInt(result)
}

func (self *DDIV) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	if v1 == 0 {
		panic("[gvm][div] Arithmetic error ，divisor cannot be 0")
	}
	result := v2 / v1
	stack.PushDouble(result)
}

func (self *LDIV) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	if v1 == 0 {
		panic("[gvm][div] Arithmetic error ，divisor cannot be 0")
	}
	result := v2 / v1
	stack.PushLong(result)
}

func (self *FDIV) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	if v1 == 0 {
		panic("[gvm][div] Arithmetic error ，divisor cannot be 0")
	}
	result := v2 / v1
	stack.PushFloat(result)
}
