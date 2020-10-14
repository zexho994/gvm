package math

import (
	"../../rtda"
	"../base"
)

type ISUB struct {
	base.NoOperandsInstruction
}

type LSUB struct {
	base.NoOperandsInstruction
}

type DSUB struct {
	base.NoOperandsInstruction
}

type FSUB struct {
	base.NoOperandsInstruction
}

func (self *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	subResult := v2 - v1
	stack.PushInt(subResult)
}

func (self *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	subResule := v2 - v1
	stack.PushLong(subResule)
}

func (self *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	subResult := v2 - v1
	stack.PushDouble(subResult)
}

func (self *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	subResult := v2 - v1
	stack.PushFloat(subResult)
}
