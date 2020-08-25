package math

import (
	"../../instructions/base"
	"../../rtda"
)

type LAND struct {
	base.NoOperandsInstruction
}
type IAND struct {
	base.NoOperandsInstruction
}

func (self LAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}

func (self IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}
