package math

import (
	"../../instructions/base"
	"../../rtda"
)

type INEG struct {
	base.NoOperandsInstruction
}

type LNEG struct {
	base.NoOperandsInstruction
}

type FNEG struct {
	base.NoOperandsInstruction
}

type DNEG struct {
	base.NoOperandsInstruction
}

func (self FNEG) Execute(frame rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	n := -v1
	stack.PushFloat(n)
}

func (self INEG) Execute(frame rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	n := -v1
	stack.PushInt(n)
}

func (self DNEG) Execute(frame rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	n := -v1
	stack.PushDouble(n)
}

func (self LNEG) Execute(frame rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	n := -v1
	stack.PushLong(n)
}
