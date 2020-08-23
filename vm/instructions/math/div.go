package math

import (
	"../../instructions/base"
	"../../rtda"
)

type IDIV struct {
	base.NoOperandsInstruction
}

type LDIV struct {
	base.NoOperandsInstruction
}

type DDIV struct {
	base.NoOperandsInstruction
}

type FDIV struct {
	base.NoOperandsInstruction
}

func (self IDIV) Execute(frame rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	if v1 == 0 {
		panic("[gvm][div] Arithmetic error ，divisor cannot be 0")
	}
	result := v2 / v1
	stack.PushInt(result)
}

func (self DDIV) Execute(frame rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	if v1 == 0 {
		panic("[gvm][div] Arithmetic error ，divisor cannot be 0")
	}
	result := v2 / v1
	stack.PushDouble(result)
}

func (self LDIV) Execute(frame rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	if v1 == 0 {
		panic("[gvm][div] Arithmetic error ，divisor cannot be 0")
	}
	result := v2 / v1
	stack.PushLong(result)
}

func (self FDIV) Execute(frame rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	if v1 == 0 {
		panic("[gvm][div] Arithmetic error ，divisor cannot be 0")
	}
	result := v2 / v1
	stack.PushFloat(result)
}
